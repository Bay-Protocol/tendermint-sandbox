package main

import (
	"fmt"
	"strconv"

	"github.com/tendermint/tendermint/abci/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
)

type SandboxApp struct {
	types.BaseApplication
}

var _ abcitypes.Application = (*SandboxApp)(nil)

func NewSandboxApp() *SandboxApp {
	return &SandboxApp{}
}

func (app *SandboxApp) DeliverTx(req abcitypes.RequestDeliverTx) abcitypes.ResponseDeliverTx {
	code, gas := app.estimateTx(req.Tx)
	fmt.Printf("DeliverTx Gas = %d\n", gas)
	return abcitypes.ResponseDeliverTx{Code: code, GasWanted: gas, GasUsed: gas}
}

func (app *SandboxApp) CheckTx(req abcitypes.RequestCheckTx) abcitypes.ResponseCheckTx {
	code, gas := app.estimateTx(req.Tx)
	fmt.Printf("CheckTx Gas = %d\n", gas)
	return abcitypes.ResponseCheckTx{Code: code, GasWanted: gas, Priority: gas}
}

func (SandboxApp) EndBlock(req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return abcitypes.ResponseEndBlock{}
}

func (app *SandboxApp) estimateTx(tx []byte) (code uint32, gas int64) {
	value, err := strconv.ParseInt(string(tx[:]), 10, 64)
	if err != nil {
		return 1, 0
	}
	if value <= 0 {
		return 2, 0
	}
	return 0, value
}
