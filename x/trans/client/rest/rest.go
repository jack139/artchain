package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers trans-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/trans/transactions/{id}", getTransactionHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/trans/transactions", listTransactionHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	//r.HandleFunc("/trans/transactions", createTransactionHandler(clientCtx)).Methods("POST")
	//r.HandleFunc("/trans/transactions/{id}", updateTransactionHandler(clientCtx)).Methods("POST")
	//r.HandleFunc("/trans/transactions/{id}", deleteTransactionHandler(clientCtx)).Methods("POST")

}
