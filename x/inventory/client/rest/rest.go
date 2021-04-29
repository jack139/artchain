package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers inventory-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/inventory/reviews/{id}", getReviewHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/inventory/reviews", listReviewHandler(clientCtx)).Methods("GET")

	r.HandleFunc("/inventory/items/{id}", getItemHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/inventory/items", listItemHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	//r.HandleFunc("/inventory/reviews", createReviewHandler(clientCtx)).Methods("POST")
	//r.HandleFunc("/inventory/reviews/{id}", updateReviewHandler(clientCtx)).Methods("POST")
	//r.HandleFunc("/inventory/reviews/{id}", deleteReviewHandler(clientCtx)).Methods("POST")

	//r.HandleFunc("/inventory/items", createItemHandler(clientCtx)).Methods("POST")
	//r.HandleFunc("/inventory/items/{id}", updateItemHandler(clientCtx)).Methods("POST")
	//r.HandleFunc("/inventory/items/{id}", deleteItemHandler(clientCtx)).Methods("POST")

}
