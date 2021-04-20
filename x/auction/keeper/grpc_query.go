package keeper

import (
	"github.com/jack139/artchain/x/auction/types"
)

var _ types.QueryServer = Keeper{}
