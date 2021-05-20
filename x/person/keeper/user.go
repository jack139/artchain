package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/person/types"
	"strconv"
	"strings"
)

// GetUserCount get the total number of user
func (k Keeper) GetUserCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCountKey))
	byteKey := types.KeyPrefix(types.UserCountKey)
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

// SetUserCount set the total number of user
func (k Keeper) SetUserCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserCountKey))
	byteKey := types.KeyPrefix(types.UserCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendUser appends a user in the store with a new id and update the count
func (k Keeper) AppendUser(
	ctx sdk.Context,
	creator string,
	recType string,
	name string,
	userType string,
	userInfo string,
	status string,
	regDate string,
	chainAddr string,
	lastDate string,
) uint64 {
	// Create the user
	count := k.GetUserCount(ctx)
	var user = types.User{
		Creator:   creator,
		Id:        count,
		RecType:   recType,
		Name:      name,
		UserType:  userType,
		UserInfo:  userInfo,
		Status:    status,
		RegDate:   regDate,
		ChainAddr: chainAddr,
		LastDate:  lastDate,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	value := k.cdc.MustMarshalBinaryBare(&user)
	store.Set(GetUserIDBytes(user.Id), value)

	// 添加chainAddr到id的索引
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddrIndexKey))
	store2.Set([]byte(chainAddr), GetUserIDBytes(user.Id))

	// Update user count
	k.SetUserCount(ctx, count+1)

	return count
}

// SetUser set a specific user in the store
func (k Keeper) SetUser(ctx sdk.Context, user types.User) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	b := k.cdc.MustMarshalBinaryBare(&user)
	store.Set(GetUserIDBytes(user.Id), b)

	// 添加chainAddr到id的索引
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddrIndexKey))
	store2.Set([]byte(user.ChainAddr), GetUserIDBytes(user.Id))
}

// GetUser returns a user from its id
func (k Keeper) GetUser(ctx sdk.Context, id uint64) types.User {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	var user types.User
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetUserIDBytes(id)), &user)
	return user
}

// HasUser checks if the user exists in the store
func (k Keeper) HasUser(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	return store.Has(GetUserIDBytes(id))
}

// GetUserOwner returns the creator of the user
func (k Keeper) GetUserOwner(ctx sdk.Context, id uint64) string {
	return k.GetUser(ctx, id).Creator
}

// RemoveUser removes a user from the store
func (k Keeper) RemoveUser(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))

	// 删除chainAddr到id的索引
	var user types.User
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetUserIDBytes(id)), &user)
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddrIndexKey))
	store2.Delete([]byte(user.ChainAddr))

	// 删除 user
	store.Delete(GetUserIDBytes(id))
}

// GetAllUser returns all user
func (k Keeper) GetAllUser(ctx sdk.Context) (list []types.User) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.User
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetUserIDBytes returns the byte representation of the ID
func GetUserIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetUserIDFromBytes returns ID in uint64 format from a byte array
func GetUserIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}


// 使用chainAddr取得用户，使用索引 ChainAddr --> Id
func (k Keeper) GetUserByChainAddr(ctx sdk.Context, chainAddr string) types.User {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AddrIndexKey))
	var user types.User
	idBytes := store2.Get([]byte(chainAddr))
	k.cdc.MustUnmarshalBinaryBare(store.Get(idBytes), &user)
	return user
}


// 使用useType取得用户，全表遍历
func (k Keeper) GetUserByUserType(ctx sdk.Context, userType string) (list []types.User) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.User
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		if strings.HasPrefix(val.UserType, userType) {
			list = append(list, val)
		} 
	}
	return
}

// 使用status取得用户，全表遍历
func (k Keeper) GetUserByStatus(ctx sdk.Context, status string) (list []types.User) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.UserKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.User
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		if val.Status==status {
			list = append(list, val)
		} 
	}
	return
}
