package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/jack139/artchain/x/auction/types"
)

type createRequestRequest struct {
	BaseReq        rest.BaseReq `json:"base_req"`
	Creator        string       `json:"creator"`
	RecType        string       `json:"recType"`
	ItemId         string       `json:"itemId"`
	AuctionHouseId string       `json:"auctionHouseId"`
	SellerId       string       `json:"SellerId"`
	RequestDate    string       `json:"requestDate"`
	ReservePrice   string       `json:"reservePrice"`
	Status         string       `json:"status"`
	OpenDate       string       `json:"openDate"`
	CloseDate      string       `json:"closeDate"`
}

func createRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createRequestRequest
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

		parsedItemId := req.ItemId

		parsedAuctionHouseId := req.AuctionHouseId

		parsedSellerId := req.SellerId

		parsedRequestDate := req.RequestDate

		parsedReservePrice := req.ReservePrice

		parsedStatus := req.Status

		parsedOpenDate := req.OpenDate

		parsedCloseDate := req.CloseDate

		msg := types.NewMsgCreateRequest(
			req.Creator,
			parsedRecType,
			parsedItemId,
			parsedAuctionHouseId,
			parsedSellerId,
			parsedRequestDate,
			parsedReservePrice,
			parsedStatus,
			parsedOpenDate,
			parsedCloseDate,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateRequestRequest struct {
	BaseReq        rest.BaseReq `json:"base_req"`
	Creator        string       `json:"creator"`
	RecType        string       `json:"recType"`
	ItemId         string       `json:"itemId"`
	AuctionHouseId string       `json:"auctionHouseId"`
	SellerId       string       `json:"SellerId"`
	RequestDate    string       `json:"requestDate"`
	ReservePrice   string       `json:"reservePrice"`
	Status         string       `json:"status"`
	OpenDate       string       `json:"openDate"`
	CloseDate      string       `json:"closeDate"`
}

func updateRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updateRequestRequest
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

		parsedItemId := req.ItemId

		parsedAuctionHouseId := req.AuctionHouseId

		parsedSellerId := req.SellerId

		parsedRequestDate := req.RequestDate

		parsedReservePrice := req.ReservePrice

		parsedStatus := req.Status

		parsedOpenDate := req.OpenDate

		parsedCloseDate := req.CloseDate

		msg := types.NewMsgUpdateRequest(
			req.Creator,
			id,
			parsedRecType,
			parsedItemId,
			parsedAuctionHouseId,
			parsedSellerId,
			parsedRequestDate,
			parsedReservePrice,
			parsedStatus,
			parsedOpenDate,
			parsedCloseDate,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteRequestRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deleteRequestRequest
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

		msg := types.NewMsgDeleteRequest(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
