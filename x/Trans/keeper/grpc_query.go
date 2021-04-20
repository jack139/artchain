package keeper

import (
	"github.com/jack139/artchain/x/Trans/types"
)

var _ types.QueryServer = Keeper{}
