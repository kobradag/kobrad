package rpchandlers

import (
	"github.com/kobradag/kobrad/app/appmessage"
	"github.com/kobradag/kobrad/app/rpc/rpccontext"
	"github.com/kobradag/kobrad/domain/consensus/utils/constants"
	"github.com/kobradag/kobrad/infrastructure/network/netadapter/router"
)

// HandleGetCoinSupply handles the respectively named RPC command
func HandleGetCoinSupply(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	if !context.Config.UTXOIndex {
		errorMessage := &appmessage.GetCoinSupplyResponseMessage{}
		errorMessage.Error = appmessage.RPCErrorf("Method unavailable when kobrad is run without --utxoindex")
		return errorMessage, nil
	}

	circulatingLeorSupply, err := context.UTXOIndex.GetCirculatingLeorSupply()
	if err != nil {
		return nil, err
	}

	response := appmessage.NewGetCoinSupplyResponseMessage(
		constants.MaxLeor,
		circulatingLeorSupply,
	)

	return response, nil
}
