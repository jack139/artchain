package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/inventory/types"
	"strconv"
)

// GetReviewCount get the total number of review
func (k Keeper) GetReviewCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewCountKey))
	byteKey := types.KeyPrefix(types.ReviewCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseUint(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to iint64
		panic("cannot decode count")
	}

	return count
}

// SetReviewCount set the total number of review
func (k Keeper) SetReviewCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewCountKey))
	byteKey := types.KeyPrefix(types.ReviewCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendReview appends a review in the store with a new id and update the count
func (k Keeper) AppendReview(
	ctx sdk.Context,
	creator string,
	recType string,
	itemId string,
	reviewerId string,
	reviewDetail string,
	reviewDate string,
	upCount string,
	downCount string,
	status string,
	lastDate string,
) uint64 {
	// Create the review
	count := k.GetReviewCount(ctx)
	var review = types.Review{
		Creator:      creator,
		Id:           count,
		RecType:      recType,
		ItemId:       itemId,
		ReviewerId:   reviewerId,
		ReviewDetail: reviewDetail,
		ReviewDate:   reviewDate,
		UpCount:      upCount,
		DownCount:    downCount,
		Status:       status,
		LastDate:     lastDate,
	}

	// review 的 key 按 itemId 区分，方便按 itemId 查询
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey+itemId))
	value := k.cdc.MustMarshalBinaryBare(&review)
	store.Set(GetReviewIDBytes(review.Id), value)

	// Update review count
	k.SetReviewCount(ctx, count+1)

	return count
}

// SetReview set a specific review in the store
func (k Keeper) SetReview(ctx sdk.Context, review types.Review) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey+review.ItemId))
	b := k.cdc.MustMarshalBinaryBare(&review)
	store.Set(GetReviewIDBytes(review.Id), b)
}

// GetReview returns a review from its id
func (k Keeper) GetReview(ctx sdk.Context, id uint64, itemId string) types.Review {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey+itemId))
	var review types.Review
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetReviewIDBytes(id)), &review)
	return review
}

// HasReview checks if the review exists in the store
func (k Keeper) HasReview(ctx sdk.Context, id uint64, itemId string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey+itemId))
	return store.Has(GetReviewIDBytes(id))
}

// GetReviewOwner returns the creator of the review
func (k Keeper) GetReviewOwner(ctx sdk.Context, id uint64, itemId string) string {
	return k.GetReview(ctx, id, itemId).Creator
}

// RemoveReview removes a review from the store
func (k Keeper) RemoveReview(ctx sdk.Context, id uint64, itemId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey+itemId))
	store.Delete(GetReviewIDBytes(id))
}

// GetAllReview returns all review
func (k Keeper) GetAllReview(ctx sdk.Context, itemId string) (list []types.Review) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ReviewKey+itemId))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Review
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetReviewIDBytes returns the byte representation of the ID
func GetReviewIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetReviewIDFromBytes returns ID in uint64 format from a byte array
func GetReviewIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
