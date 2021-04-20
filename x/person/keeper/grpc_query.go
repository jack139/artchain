package keeper

import (
	"github.com/jack139/artchain/x/person/types"
)

var _ types.QueryServer = Keeper{}
