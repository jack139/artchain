package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/inventory/types"
	"strconv"
)

// GetItemCount get the total number of item
func (k Keeper) GetItemCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemCountKey))
	byteKey := types.KeyPrefix(types.ItemCountKey)
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

// SetItemCount set the total number of item
func (k Keeper) SetItemCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemCountKey))
	byteKey := types.KeyPrefix(types.ItemCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendItem appends a item in the store with a new id and update the count
func (k Keeper) AppendItem(
	ctx sdk.Context,
	creator string,
	recType string,
	itemDesc string,
	itemDetail string,
	itemDate string,
	itemType string,
	itemSubject string,
	itemMedia string,
	itemSize string,
	itemImage string,
	AESKey string,
	itemBasePrice string,
	currentOwnerId string,
	status string,
	lastDate string,
) uint64 {
	// Create the item
	count := k.GetItemCount(ctx)
	var item = types.Item{
		Creator:        creator,
		Id:             count,
		RecType:        recType,
		ItemDesc:       itemDesc,
		ItemDetail:     itemDetail,
		ItemDate:       itemDate,
		ItemType:       itemType,
		ItemSubject:    itemSubject,
		ItemMedia:      itemMedia,
		ItemSize:       itemSize,
		ItemImage:      itemImage,
		AESKey:         AESKey,
		ItemBasePrice:  itemBasePrice,
		CurrentOwnerId: currentOwnerId,
		LastDate:       lastDate,
		Status:         status,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKey))
	value := k.cdc.MustMarshalBinaryBare(&item)
	store.Set(GetItemIDBytes(item.Id), value)

	// 添加owner到id的索引
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemOwnerKey+currentOwnerId))
	store2.Set(GetItemIDBytes(item.Id), GetItemIDBytes(item.Id))

	// Update item count
	k.SetItemCount(ctx, count+1)

	return count
}

// SetItem set a specific item in the store
func (k Keeper) SetItem(ctx sdk.Context, item types.Item) {
	oldItem := k.GetItem(ctx, item.Id)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKey))
	b := k.cdc.MustMarshalBinaryBare(&item)
	store.Set(GetItemIDBytes(item.Id), b)

	if oldItem.CurrentOwnerId != item.CurrentOwnerId { // 变更所有人，需要变更索引
		store1 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemOwnerKey+oldItem.CurrentOwnerId))
		store1.Delete(GetItemIDBytes(item.Id))
		store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemOwnerKey+item.CurrentOwnerId))
		store2.Set(GetItemIDBytes(item.Id), GetItemIDBytes(item.Id))
	}
}

// GetItem returns a item from its id
func (k Keeper) GetItem(ctx sdk.Context, id uint64) types.Item {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKey))
	var item types.Item
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetItemIDBytes(id)), &item)
	return item
}

// HasItem checks if the item exists in the store
func (k Keeper) HasItem(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKey))
	return store.Has(GetItemIDBytes(id))
}

// GetItemOwner returns the creator of the item
func (k Keeper) GetItemOwner(ctx sdk.Context, id uint64) string {
	return k.GetItem(ctx, id).Creator
}

// RemoveItem removes a item from the store
func (k Keeper) RemoveItem(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKey))

	// 删除owner到id的索引
	var item types.Item
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetItemIDBytes(id)), &item)
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemOwnerKey+item.CurrentOwnerId))
	store2.Delete(GetItemIDBytes(id))

	// 删除 item
	store.Delete(GetItemIDBytes(id))
}

// GetAllItem returns all item
func (k Keeper) GetAllItem(ctx sdk.Context) (list []types.Item) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Item
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetItemIDBytes returns the byte representation of the ID
func GetItemIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetItemIDFromBytes returns ID in uint64 format from a byte array
func GetItemIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

/*
// 使用status取得，全表遍历
func (k Keeper) GetItemByStatus(ctx sdk.Context, status string) (list []types.Item) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ItemKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Item
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		if val.Status==status {
			list = append(list, val)
		} 
	}
	return
}
*/