package keeper

import (
	"github.com/jack139/artchain/x/trans/types"
)

var _ types.QueryServer = Keeper{}
