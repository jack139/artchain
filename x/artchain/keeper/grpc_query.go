package keeper

import (
	"github.com/jack139/artchain/x/artchain/types"
)

var _ types.QueryServer = Keeper{}
