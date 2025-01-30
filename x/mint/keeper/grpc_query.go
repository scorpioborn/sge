package keeper

import (
	"github.com/sge-network/sge/x/mint/types"
)

var _ types.QueryServer = queryServer{}


func NewQueryServerImpl(k Keeper) types.QueryServer {
	return queryServer{k}
}

type queryServer struct {
	k Keeper
}