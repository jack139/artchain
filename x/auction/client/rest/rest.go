package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers auction-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/auction/bids/{id}", getBidHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/auction/bids", listBidHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/auction/requests/{id}", getRequestHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/auction/requests", listRequestHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	r.HandleFunc("/auction/bids", createBidHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/auction/bids/{id}", updateBidHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/auction/bids/{id}", deleteBidHandler(clientCtx)).Methods("POST")

	r.HandleFunc("/auction/requests", createRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/auction/requests/{id}", updateRequestHandler(clientCtx)).Methods("POST")
	r.HandleFunc("/auction/requests/{id}", deleteRequestHandler(clientCtx)).Methods("POST")

}
