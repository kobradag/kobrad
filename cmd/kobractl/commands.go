package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/kobradag/kobrad/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.KobradMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.KobradMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.KobradMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.KobradMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.KobradMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.KobradMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.KobradMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.KobradMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.KobradMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.KobradMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.KobradMessage_BanRequest{}),
	reflect.TypeOf(protowire.KobradMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
