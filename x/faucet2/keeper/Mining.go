package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/faucet2/types"
	"strconv"
	"time"
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
	LastTime int64,
	Total string,
) uint64 {
	id, _ := k.AppendMining2(ctx, creator, Minter, LastTime, Total)
	return id
}

// 为了返回 mining
func (k Keeper) AppendMining2(
	ctx sdk.Context,
	creator string,
	Minter string,
	LastTime int64,
	Total string,
) (uint64, types.Mining) {
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

	// 保存minter对应id索引
	store2 := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Mining2Key))
	store2.Set([]byte(Minter), GetMiningIDBytes(Mining.Id))

	// Update Mining count
	k.SetMiningCount(ctx, count+1)

	return count, Mining
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
	// 未删除 id 与 Minter 对照关系
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


// MintAndSend mint coins and send to minter.
func (k Keeper) MintAndSend(ctx sdk.Context, sender string, minter string, mintTime int64) error {

	k.Logger(ctx).Error("MintAndSend")

	minterAcc, err := sdk.AccAddressFromBech32(minter)
	if err != nil {
		return err
	}

	var mining types.Mining

	if !k.isPresent2(ctx, minter) {
		// 不存在，新增
		//denom := k.StakingKeeper.BondDenom(ctx)
		_, mining = k.AppendMining2(ctx, sender, minter, mintTime, "0")
	} else {
		mining = k.getMining2(ctx, minter)	
	}

	// refuse mint in 24 hours
	if k.isPresent2(ctx, minter) &&
		time.Unix(mining.LastTime, 0).Add(k.Limit).UTC().After(time.Unix(mintTime, 0)) {
		return types.ErrWithdrawTooOften
	}

	denom := k.StakingKeeper.BondDenom(ctx)
	newCoin := sdk.NewCoin(denom, sdk.NewInt(k.amount))
	//mining.Total = mining.Total.Add(newCoin)
	mining.LastTime = mintTime
	k.setMining2(ctx, minter, mining)

	k.Logger(ctx).Info("Mint coin: %s", newCoin)

	err = k.SupplyKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(newCoin))
	if err != nil {
		return err
	}
	err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, minterAcc, sdk.NewCoins(newCoin))
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) getMining2(ctx sdk.Context, minter string) types.Mining {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Mining2Key))
	//if !k.isPresent2(ctx, minter) {
	//	denom := k.StakingKeeper.BondDenom(ctx)
	//	return types.NewMining2(minter, sdk.NewCoin(denom, sdk.NewInt(0)))
	//}
	bz := store.Get([]byte(minter)) // 取得 Id
	//var mining types.Mining2
	//k.cdc.MustUnmarshalBinaryBare(bz, &mining)
	return k.GetMining(ctx, GetMiningIDFromBytes(bz))
}

func (k Keeper) setMining2(ctx sdk.Context, minter string, mining types.Mining) {
	if len(mining.Minter)==0 {
		return
	}
	//if !mining.Total.IsPositive() {
	//	return
	//}

	// 不用修改 Id 与 Minter 对照
	//store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Mining2Key))
	//store.Set([]byte(ming.Minter), GetMiningIDBytes(ming.Id))

	k.SetMining(ctx, mining)
}

// IsPresent check if the name is present in the store or not
func (k Keeper) isPresent2(ctx sdk.Context, minter string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.Mining2Key))
	return store.Has([]byte(minter))
}
