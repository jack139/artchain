package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx client.Context, r *mux.Router) {
	//r.HandleFunc("/faucet/mint", mintHandlerFn(cliCtx)).Methods("POST")
}
