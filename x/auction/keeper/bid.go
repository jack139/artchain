package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/auction/types"
	"strconv"
)

// GetBidCount get the total number of bid
func (k Keeper) GetBidCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidCountKey))
	byteKey := types.KeyPrefix(types.BidCountKey)
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

// SetBidCount set the total number of bid
func (k Keeper) SetBidCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidCountKey))
	byteKey := types.KeyPrefix(types.BidCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendBid appends a bid in the store with a new id and update the count
func (k Keeper) AppendBid(
	ctx sdk.Context,
	creator string,
	recType string,
	auctionId string,
	bidNo string,
	status string,
	buyerId string,
	bidPrice string,
	bidTime string,
	lastDate string,
) uint64 {
	// Create the bid
	count := k.GetBidCount(ctx)
	var bid = types.Bid{
		Creator:   creator,
		Id:        count,
		RecType:   recType,
		AuctionId: auctionId,
		BidNo:     bidNo,
		Status:    status,
		BuyerId:   buyerId,
		BidPrice:  bidPrice,
		BidTime:   bidTime,
		LastDate:  lastDate,
	}

	// bid 的 key 按 auctionId 区分，方便按 auctionId 查询
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidKey+auctionId))
	value := k.cdc.MustMarshalBinaryBare(&bid)
	store.Set(GetBidIDBytes(bid.Id), value)

	// Update bid count
	k.SetBidCount(ctx, count+1)

	return count
}

// SetBid set a specific bid in the store
func (k Keeper) SetBid(ctx sdk.Context, bid types.Bid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidKey+bid.AuctionId))
	b := k.cdc.MustMarshalBinaryBare(&bid)
	store.Set(GetBidIDBytes(bid.Id), b)
}

// GetBid returns a bid from its id
func (k Keeper) GetBid(ctx sdk.Context, id uint64, auctionId string) types.Bid {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidKey+auctionId))
	var bid types.Bid
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetBidIDBytes(id)), &bid)
	return bid
}

// HasBid checks if the bid exists in the store
func (k Keeper) HasBid(ctx sdk.Context, id uint64, auctionId string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidKey+auctionId))
	return store.Has(GetBidIDBytes(id))
}

// GetBidOwner returns the creator of the bid
func (k Keeper) GetBidOwner(ctx sdk.Context, id uint64, auctionId string) string {
	return k.GetBid(ctx, id, auctionId).Creator
}

// RemoveBid removes a bid from the store
func (k Keeper) RemoveBid(ctx sdk.Context, id uint64, auctionId string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidKey+auctionId))
	store.Delete(GetBidIDBytes(id))
}

// GetAllBid returns all bid
func (k Keeper) GetAllBid(ctx sdk.Context, auctionId string) (list []types.Bid) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.BidKey+auctionId))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Bid
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetBidIDBytes returns the byte representation of the ID
func GetBidIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetBidIDFromBytes returns ID in uint64 format from a byte array
func GetBidIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
