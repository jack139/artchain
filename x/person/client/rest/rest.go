package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	// this line is used by starport scaffolding # 1
)

const (
	MethodGet = "GET"
)

// RegisterRoutes registers person-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)

}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 3
	r.HandleFunc("/person/users/{id}", getUserHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/person/users", listUserHandler(clientCtx)).Methods("GET")

}

func registerTxHandlers(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 4
	//r.HandleFunc("/person/users", createUserHandler(clientCtx)).Methods("POST")
	//r.HandleFunc("/person/users/{id}", updateUserHandler(clientCtx)).Methods("POST")
	//r.HandleFunc("/person/users/{id}", deleteUserHandler(clientCtx)).Methods("POST")

}
