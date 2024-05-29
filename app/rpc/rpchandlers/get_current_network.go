package rpchandlers

import (
	"github.com/kobradag/kobrad/app/appmessage"
	"github.com/kobradag/kobrad/app/rpc/rpccontext"
	"github.com/kobradag/kobrad/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
