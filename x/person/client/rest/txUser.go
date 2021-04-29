package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/jack139/artchain/x/person/types"
)

type createUserRequest struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Creator   string       `json:"creator"`
	RecType   string       `json:"recType"`
	Name      string       `json:"name"`
	UserType  string       `json:"userType"`
	UserInfo  string       `json:"userInfo"`
	Status    string       `json:"status"`
	RegDate   string       `json:"regDate"`
	ChainAddr string       `json:"chainAddr"`
}

func createUserHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createUserRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedRecType := req.RecType

		parsedName := req.Name

		parsedUserType := req.UserType

		parsedUserInfo := req.UserInfo

		parsedStatus := req.Status

		parsedRegDate := req.RegDate

		parsedChainAddr := req.ChainAddr

		msg := types.NewMsgCreateUser(
			req.Creator,
			parsedRecType,
			parsedName,
			parsedUserType,
			parsedUserInfo,
			parsedStatus,
			parsedRegDate,
			parsedChainAddr,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateUserRequest struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Creator   string       `json:"creator"`
	RecType   string       `json:"recType"`
	Name      string       `json:"name"`
	UserType  string       `json:"userType"`
	UserInfo  string       `json:"userInfo"`
	Status    string       `json:"status"`
	RegDate   string       `json:"regDate"`
	ChainAddr string       `json:"chainAddr"`
}

func updateUserHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updateUserRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err = sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		parsedRecType := req.RecType

		parsedName := req.Name

		parsedUserType := req.UserType

		parsedUserInfo := req.UserInfo

		parsedStatus := req.Status

		parsedRegDate := req.RegDate

		parsedChainAddr := req.ChainAddr

		msg := types.NewMsgUpdateUser(
			req.Creator,
			id,
			parsedRecType,
			parsedName,
			parsedUserType,
			parsedUserInfo,
			parsedStatus,
			parsedRegDate,
			parsedChainAddr,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteUserRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteUserHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deleteUserRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err = sdk.AccAddressFromBech32(req.Creator)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgDeleteUser(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
