package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/faucet2/types"
	"strconv"
)

// GetMiningCount get the total number of Mining
func (k Keeper) GetMiningCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MiningCountKey))
	byteKey := types.KeyPrefix(types.MiningCountKey)
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

// SetMiningCount set the total number of Mining
func (k Keeper) SetMiningCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MiningCountKey))
	byteKey := types.KeyPrefix(types.MiningCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendMining appends a Mining in the store with a new id and update the count
func (k Keeper) AppendMining(
	ctx sdk.Context,
	creator string,
	Minter string,
	LastTime string,
	Total string,
) uint64 {
	// Create the Mining
	count := k.GetMiningCount(ctx)
	var Mining = types.Mining{
		Creator:  creator,
		Id:       count,
		Minter:   Minter,
		LastTime: LastTime,
		Total:    Total,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MiningKey))
	value := k.cdc.MustMarshalBinaryBare(&Mining)
	store.Set(GetMiningIDBytes(Mining.Id), value)

	// Update Mining count
	k.SetMiningCount(ctx, count+1)

	return count
}

// SetMining set a specific Mining in the store
func (k Keeper) SetMining(ctx sdk.Context, Mining types.Mining) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MiningKey))
	b := k.cdc.MustMarshalBinaryBare(&Mining)
	store.Set(GetMiningIDBytes(Mining.Id), b)
}

// GetMining returns a Mining from its id
func (k Keeper) GetMining(ctx sdk.Context, id uint64) types.Mining {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MiningKey))
	var Mining types.Mining
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetMiningIDBytes(id)), &Mining)
	return Mining
}

// HasMining checks if the Mining exists in the store
func (k Keeper) HasMining(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MiningKey))
	return store.Has(GetMiningIDBytes(id))
}

// GetMiningOwner returns the creator of the Mining
func (k Keeper) GetMiningOwner(ctx sdk.Context, id uint64) string {
	return k.GetMining(ctx, id).Creator
}

// RemoveMining removes a Mining from the store
func (k Keeper) RemoveMining(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MiningKey))
	store.Delete(GetMiningIDBytes(id))
}

// GetAllMining returns all Mining
func (k Keeper) GetAllMining(ctx sdk.Context) (list []types.Mining) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.MiningKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Mining
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetMiningIDBytes returns the byte representation of the ID
func GetMiningIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetMiningIDFromBytes returns ID in uint64 format from a byte array
func GetMiningIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
