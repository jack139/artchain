package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/jack139/artchain/x/trans/types"
)

type createTransactionRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Creator     string       `json:"creator"`
	RecType     string       `json:"recType"`
	AuctionId   string       `json:"auctionId"`
	ItemId      string       `json:"itemId"`
	TransType   string       `json:"transType"`
	UserId      string       `json:"userId"`
	TransDate   string       `json:"transDate"`
	HammerTime  string       `json:"hammerTime"`
	HammerPrice string       `json:"hammerPrice"`
	Details     string       `json:"details"`
	Status      string       `json:"status"`
}

func createTransactionHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createTransactionRequest
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

		parsedAuctionId := req.AuctionId

		parsedItemId := req.ItemId

		parsedTransType := req.TransType

		parsedUserId := req.UserId

		parsedTransDate := req.TransDate

		parsedHammerTime := req.HammerTime

		parsedHammerPrice := req.HammerPrice

		parsedDetails := req.Details

		parsedStatus := req.Status

		msg := types.NewMsgCreateTransaction(
			req.Creator,
			parsedRecType,
			parsedAuctionId,
			parsedItemId,
			parsedTransType,
			parsedUserId,
			parsedTransDate,
			parsedHammerTime,
			parsedHammerPrice,
			parsedDetails,
			parsedStatus,
			"",
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateTransactionRequest struct {
	BaseReq     rest.BaseReq `json:"base_req"`
	Creator     string       `json:"creator"`
	RecType     string       `json:"recType"`
	AuctionId   string       `json:"auctionId"`
	ItemId      string       `json:"itemId"`
	TransType   string       `json:"transType"`
	UserId      string       `json:"userId"`
	TransDate   string       `json:"transDate"`
	HammerTime  string       `json:"hammerTime"`
	HammerPrice string       `json:"hammerPrice"`
	Details     string       `json:"details"`
	Status      string       `json:"status"`
}

func updateTransactionHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updateTransactionRequest
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

		parsedAuctionId := req.AuctionId

		parsedItemId := req.ItemId

		parsedTransType := req.TransType

		parsedUserId := req.UserId

		parsedTransDate := req.TransDate

		parsedHammerTime := req.HammerTime

		parsedHammerPrice := req.HammerPrice

		parsedDetails := req.Details

		parsedStatus := req.Status

		msg := types.NewMsgUpdateTransaction(
			req.Creator,
			id,
			parsedRecType,
			parsedAuctionId,
			parsedItemId,
			parsedTransType,
			parsedUserId,
			parsedTransDate,
			parsedHammerTime,
			parsedHammerPrice,
			parsedDetails,
			parsedStatus,
			"",
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteTransactionRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteTransactionHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deleteTransactionRequest
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

		msg := types.NewMsgDeleteTransaction(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
