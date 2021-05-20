package rest

import (
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/jack139/artchain/x/inventory/types"
)

type createReviewRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Creator      string       `json:"creator"`
	RecType      string       `json:"recType"`
	ItemId       string       `json:"itemId"`
	ReviewerId   string       `json:"reviewerId"`
	ReviewDetail string       `json:"reviewDetail"`
	ReviewDate   string       `json:"reviewDate"`
	UpCount      string       `json:"upCount"`
	DownCount    string       `json:"downCount"`
	Status       string       `json:"status"`
}

func createReviewHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createReviewRequest
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

		parsedReviewerId := req.ReviewerId

		parsedReviewDetail := req.ReviewDetail

		parsedReviewDate := req.ReviewDate

		parsedUpCount := req.UpCount

		parsedDownCount := req.DownCount

		parsedStatus := req.Status

		msg := types.NewMsgCreateReview(
			req.Creator,
			parsedRecType,
			parsedItemId,
			parsedReviewerId,
			parsedReviewDetail,
			parsedReviewDate,
			parsedUpCount,
			parsedDownCount,
			parsedStatus,
			"",
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateReviewRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Creator      string       `json:"creator"`
	RecType      string       `json:"recType"`
	ItemId       string       `json:"itemId"`
	ReviewerId   string       `json:"reviewerId"`
	ReviewDetail string       `json:"reviewDetail"`
	ReviewDate   string       `json:"reviewDate"`
	UpCount      string       `json:"upCount"`
	DownCount    string       `json:"downCount"`
}

func updateReviewHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updateReviewRequest
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

		parsedReviewerId := req.ReviewerId

		parsedReviewDetail := req.ReviewDetail

		parsedReviewDate := req.ReviewDate

		parsedUpCount := req.UpCount

		parsedDownCount := req.DownCount

		msg := types.NewMsgUpdateReview(
			req.Creator,
			id,
			parsedRecType,
			parsedItemId,
			parsedReviewerId,
			parsedReviewDetail,
			parsedReviewDate,
			parsedUpCount,
			parsedDownCount,
			"WAIT",
			"",
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteReviewRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteReviewHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deleteReviewRequest
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

		msg := types.NewMsgDeleteReview(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
