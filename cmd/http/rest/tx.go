// Deprecated by sdk 0.46.0, so get back from 0.45.5
package rest

import (
	//"bufio"
	//"context"
	//"errors"
	//"fmt"
	"net/http"
	//"os"

	//gogogrpc "github.com/gogo/protobuf/grpc"
	//"github.com/spf13/pflag"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	//"github.com/cosmos/cosmos-sdk/client/input"
	//cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//"github.com/cosmos/cosmos-sdk/types/rest"
	//"github.com/cosmos/cosmos-sdk/types/tx"
	//"github.com/cosmos/cosmos-sdk/types/tx/signing"
	//authsigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	clitx "github.com/cosmos/cosmos-sdk/client/tx"
)


// WriteGeneratedTxResponse writes a generated unsigned transaction to the
// provided http.ResponseWriter. It will simulate gas costs if requested by the
// BaseReq. Upon any error, the error will be written to the http.ResponseWriter.
// Note that this function returns the legacy StdTx Amino JSON format for compatibility
// with legacy clients.
// Deprecated: We are removing Amino soon.
func WriteGeneratedTxResponse(
	clientCtx client.Context, w http.ResponseWriter, br BaseReq, msgs ...sdk.Msg,
) {
	gasAdj, ok := ParseFloat64OrReturnBadRequest(w, br.GasAdjustment, flags.DefaultGasAdjustment)
	if !ok {
		return
	}

	gasSetting, err := flags.ParseGasSetting(br.Gas)
	if CheckBadRequestError(w, err) {
		return
	}

	txf := clitx.Factory{}.
		WithFees(br.Fees.String()).
		WithGasPrices(br.GasPrices.String()).
		WithAccountNumber(br.AccountNumber).
		WithSequence(br.Sequence).
		WithGas(gasSetting.Gas).
		WithGasAdjustment(gasAdj).
		WithMemo(br.Memo).
		WithChainID(br.ChainID).
		WithSimulateAndExecute(br.Simulate).
		WithTxConfig(clientCtx.TxConfig).
		WithTimeoutHeight(br.TimeoutHeight)

	if br.Simulate || gasSetting.Simulate {
		if gasAdj < 0 {
			WriteErrorResponse(w, http.StatusBadRequest, sdkerrors.ErrorInvalidGasAdjustment.Error())
			return
		}

		_, adjusted, err := clitx.CalculateGas(clientCtx, txf, msgs...)
		if CheckInternalServerError(w, err) {
			return
		}

		txf = txf.WithGas(adjusted)

		if br.Simulate {
			WriteSimulationResponse(w, clientCtx.LegacyAmino, txf.Gas())
			return
		}
	}

	tx, err := txf.BuildUnsignedTx(msgs...)
	if CheckBadRequestError(w, err) {
		return
	}

	stdTx, err := clitx.ConvertTxToStdTx(clientCtx.LegacyAmino, tx.GetTx())
	if CheckInternalServerError(w, err) {
		return
	}

	output, err := clientCtx.LegacyAmino.MarshalJSON(stdTx)
	if CheckInternalServerError(w, err) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(output)
}
