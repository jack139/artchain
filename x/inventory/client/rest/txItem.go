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

type createItemRequest struct {
	BaseReq        rest.BaseReq `json:"base_req"`
	Creator        string       `json:"creator"`
	RecType        string       `json:"recType"`
	ItemDesc       string       `json:"itemDesc"`
	ItemDetail     string       `json:"itemDetail"`
	ItemDate       string       `json:"itemDate"`
	ItemType       string       `json:"itemType"`
	ItemSubject    string       `json:"itemSubject"`
	ItemMedia      string       `json:"itemMedia"`
	ItemSize       string       `json:"itemSize"`
	ItemImage      string       `json:"itemImage"`
	AESKey         string       `json:"AESKey"`
	ItemBasePrice  string       `json:"itemBasePrice"`
	CurrentOwnerId string       `json:"currentOwnerId"`
	Status         string       `json:"status"`
}

func createItemHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createItemRequest
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

		parsedItemDesc := req.ItemDesc

		parsedItemDetail := req.ItemDetail

		parsedItemDate := req.ItemDate

		parsedItemType := req.ItemType

		parsedItemSubject := req.ItemSubject

		parsedItemMedia := req.ItemMedia

		parsedItemSize := req.ItemSize

		parsedItemImage := req.ItemImage

		parsedAESKey := req.AESKey

		parsedItemBasePrice := req.ItemBasePrice

		parsedCurrentOwnerId := req.CurrentOwnerId

		parsedStatus := req.Status

		msg := types.NewMsgCreateItem(
			req.Creator,
			parsedRecType,
			parsedItemDesc,
			parsedItemDetail,
			parsedItemDate,
			parsedItemType,
			parsedItemSubject,
			parsedItemMedia,
			parsedItemSize,
			parsedItemImage,
			parsedAESKey,
			parsedItemBasePrice,
			parsedCurrentOwnerId,
			parsedStatus,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type updateItemRequest struct {
	BaseReq        rest.BaseReq `json:"base_req"`
	Creator        string       `json:"creator"`
	RecType        string       `json:"recType"`
	ItemDesc       string       `json:"itemDesc"`
	ItemDetail     string       `json:"itemDetail"`
	ItemDate       string       `json:"itemDate"`
	ItemType       string       `json:"itemType"`
	ItemSubject    string       `json:"itemSubject"`
	ItemMedia      string       `json:"itemMedia"`
	ItemSize       string       `json:"itemSize"`
	ItemImage      string       `json:"itemImage"`
	AESKey         string       `json:"AESKey"`
	ItemBasePrice  string       `json:"itemBasePrice"`
	CurrentOwnerId string       `json:"currentOwnerId"`
	Status         string       `json:"status"`
}

func updateItemHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req updateItemRequest
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

		parsedItemDesc := req.ItemDesc

		parsedItemDetail := req.ItemDetail

		parsedItemDate := req.ItemDate

		parsedItemType := req.ItemType

		parsedItemSubject := req.ItemSubject

		parsedItemMedia := req.ItemMedia

		parsedItemSize := req.ItemSize

		parsedItemImage := req.ItemImage

		parsedAESKey := req.AESKey

		parsedItemBasePrice := req.ItemBasePrice

		parsedCurrentOwnerId := req.CurrentOwnerId

		parsedStatus := req.Status

		msg := types.NewMsgUpdateItem(
			req.Creator,
			id,
			parsedRecType,
			parsedItemDesc,
			parsedItemDetail,
			parsedItemDate,
			parsedItemType,
			parsedItemSubject,
			parsedItemMedia,
			parsedItemSize,
			parsedItemImage,
			parsedAESKey,
			parsedItemBasePrice,
			parsedCurrentOwnerId,
			parsedStatus,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type deleteItemRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Creator string       `json:"creator"`
}

func deleteItemHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
		if err != nil {
			return
		}

		var req deleteItemRequest
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

		msg := types.NewMsgDeleteItem(
			req.Creator,
			id,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
