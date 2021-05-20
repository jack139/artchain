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

type createBidRequest struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Creator   string       `json:"creator"`
	RecType   string       `json:"recType"`
	AuctionId string       `json:"auctionId"`
	BidNo     string       `json:"bidNo"`
	ItemId    string       `json:"itemId"`
	BuyerId   string       `json:"buyerId"`
	BidPrice  string       `json:"bidPrice"`
	BidTime   string       `json:"bidTime"`
}

func createBidHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createBidRequest
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

		parsedBidNo := req.BidNo

		parsedItemId := req.ItemId

		parsedBuyerId := req.BuyerId

		parsedBidPrice := req.BidPrice

		parsedBidTime := req.BidTime

		msg := types.NewMsgCreateBid(
			req.Creator,
			parsedRecType,
			parsedAuctionId,
			parsedBidNo,
			parsedItemId,
			parsedBuyerId,
			parsedBidPrice,
			parsedBidTime,
			"",
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateBidRequest struct {
	BaseReq   rest.BaseReq `json:"base_req"`
	Creator   string       `json:"creator"`
	RecType   string       `json:"recType"`
	AuctionId string       `json:"auctionId"`
	BidNo     string       `json:"bidNo"`
	ItemId    string       `json:"itemId"`
	BuyerId   string       `json:"buyerId"`
	BidPrice  string       `json:"bidPrice"`
	BidTime   string       `json:"bidTime"`
}

func updateBidHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updateBidRequest
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

		parsedBidNo := req.BidNo

		parsedItemId := req.ItemId

		parsedBuyerId := req.BuyerId

		parsedBidPrice := req.BidPrice

		parsedBidTime := req.BidTime

		msg := types.NewMsgUpdateBid(
			req.Creator,
			id,
			parsedRecType,
			parsedAuctionId,
			parsedBidNo,
			parsedItemId,
			parsedBuyerId,
			parsedBidPrice,
			parsedBidTime,
			"",
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteBidRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteBidHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deleteBidRequest
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

		msg := types.NewMsgDeleteBid(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
