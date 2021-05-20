package keeper

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jack139/artchain/x/trans/types"
	"strconv"
)

// GetTransactionCount get the total number of transaction
func (k Keeper) GetTransactionCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionCountKey))
	byteKey := types.KeyPrefix(types.TransactionCountKey)
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

// SetTransactionCount set the total number of transaction
func (k Keeper) SetTransactionCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionCountKey))
	byteKey := types.KeyPrefix(types.TransactionCountKey)
	bz := []byte(strconv.FormatUint(count, 10))
	store.Set(byteKey, bz)
}

// AppendTransaction appends a transaction in the store with a new id and update the count
func (k Keeper) AppendTransaction(
	ctx sdk.Context,
	creator string,
	recType string,
	auctionId string,
	itemId string,
	transType string,
	userId string,
	transDate string,
	hammerTime string,
	hammerPrice string,
	details string,
	status string,
	lastDate string,
) uint64 {
	// Create the transaction
	count := k.GetTransactionCount(ctx)
	var transaction = types.Transaction{
		Creator:     creator,
		Id:          count,
		RecType:     recType,
		AuctionId:   auctionId,
		ItemId:      itemId,
		TransType:   transType,
		UserId:      userId,
		TransDate:   transDate,
		HammerTime:  hammerTime,
		HammerPrice: hammerPrice,
		Details:     details,
		Status:      status,
		LastDate:    lastDate,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	value := k.cdc.MustMarshalBinaryBare(&transaction)
	store.Set(GetTransactionIDBytes(transaction.Id), value)

	// Update transaction count
	k.SetTransactionCount(ctx, count+1)

	return count
}

// SetTransaction set a specific transaction in the store
func (k Keeper) SetTransaction(ctx sdk.Context, transaction types.Transaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	b := k.cdc.MustMarshalBinaryBare(&transaction)
	store.Set(GetTransactionIDBytes(transaction.Id), b)
}

// GetTransaction returns a transaction from its id
func (k Keeper) GetTransaction(ctx sdk.Context, id uint64) types.Transaction {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	var transaction types.Transaction
	k.cdc.MustUnmarshalBinaryBare(store.Get(GetTransactionIDBytes(id)), &transaction)
	return transaction
}

// HasTransaction checks if the transaction exists in the store
func (k Keeper) HasTransaction(ctx sdk.Context, id uint64) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	return store.Has(GetTransactionIDBytes(id))
}

// GetTransactionOwner returns the creator of the transaction
func (k Keeper) GetTransactionOwner(ctx sdk.Context, id uint64) string {
	return k.GetTransaction(ctx, id).Creator
}

// RemoveTransaction removes a transaction from the store
func (k Keeper) RemoveTransaction(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	store.Delete(GetTransactionIDBytes(id))
}

// GetAllTransaction returns all transaction
func (k Keeper) GetAllTransaction(ctx sdk.Context) (list []types.Transaction) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TransactionKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Transaction
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTransactionIDBytes returns the byte representation of the ID
func GetTransactionIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTransactionIDFromBytes returns ID in uint64 format from a byte array
func GetTransactionIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
