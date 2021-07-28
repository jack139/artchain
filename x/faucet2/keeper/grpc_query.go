package keeper

import (
	"github.com/jack139/artchain/x/faucet2/types"
)

var _ types.QueryServer = Keeper{}
