package keeper

import (
	"github.com/jack139/artchain/x/inventory/types"
)

var _ types.QueryServer = Keeper{}
