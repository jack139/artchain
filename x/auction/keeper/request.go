package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/auction/types"
	"strconv"
)

// GetRequestCount get the total number of request
func (k Keeper) GetRequestCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKey))
	byteKey := types.KeyPrefix(types.RequestCountKey)
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

// SetRequestCount set the total number of request
func (k Keeper) SetRequestCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestCountKey))
	byteKey := types.KeyPrefix(types.RequestCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendRequest appends a request in the store with a new id and update the count
func (k Keeper) AppendRequest(
	ctx sdk.Context,
	creator string,
	recType string,
	itemId string,
	auctionHouseId string,
	SellerId string,
	requestDate string,
	reservePrice string,
	status string,
	openDate string,
	closeDate string,
	lastDate string,
) uint64 {
	// Create the request
	count := k.GetRequestCount(ctx)
	var request = types.Request{
		Creator:        creator,
		Id:             count,
		RecType:        recType,
		ItemId:         itemId,
		AuctionHouseId: auctionHouseId,
		SellerId:       SellerId,
		RequestDate:    requestDate,
		ReservePrice:   reservePrice,
		Status:         status,
		OpenDate:       openDate,
		CloseDate:      closeDate,
		LastDate:       lastDate,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	value := k.cdc.MustMarshalBinaryBare(&request)
	store.Set(GetRequestIDBytes(request.Id), value)

	// Update request count
	k.SetRequestCount(ctx, count+1)

	return count
}

// SetRequest set a specific request in the store
func (k Keeper) SetRequest(ctx sdk.Context, request types.Request) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	b := k.cdc.MustMarshalBinaryBare(&request)
	store.Set(GetRequestIDBytes(request.Id), b)
}

// GetRequest returns a request from its id
func (k Keeper) GetRequest(ctx sdk.Context, id uint64) types.Request {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	var request types.Request
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetRequestIDBytes(id)), &request)
	return request
}

// HasRequest checks if the request exists in the store
func (k Keeper) HasRequest(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	return store.Has(GetRequestIDBytes(id))
}

// GetRequestOwner returns the creator of the request
func (k Keeper) GetRequestOwner(ctx sdk.Context, id uint64) string {
	return k.GetRequest(ctx, id).Creator
}

// RemoveRequest removes a request from the store
func (k Keeper) RemoveRequest(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	store.Delete(GetRequestIDBytes(id))
}

// GetAllRequest returns all request
func (k Keeper) GetAllRequest(ctx sdk.Context) (list []types.Request) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Request
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetRequestIDBytes returns the byte representation of the ID
func GetRequestIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetRequestIDFromBytes returns ID in uint64 format from a byte array
func GetRequestIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}


// 使用seller_addr取得拍卖请求，全表遍历
func (k Keeper) GetRequestByChainAddr(ctx sdk.Context, chainAddr string) (list []types.Request) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.RequestKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Request
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		if val.SellerId==chainAddr {
			list = append(list, val)
		} 
	}
	return
}

