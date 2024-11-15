package protowire

import (
	"github.com/kobradag/kobrad/app/appmessage"
	"github.com/pkg/errors"
)

type converter interface {
	toAppMessage() (appmessage.Message, error)
}

// ToAppMessage converts a KobradMessage to its appmessage.Message representation
func (x *KobradMessage) ToAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "kobradMessage is nil")
	}
	converter, ok := x.Payload.(converter)
	if !ok {
		return nil, errors.Errorf("received invalid message")
	}
	appMessage, err := converter.toAppMessage()
	if err != nil {
		return nil, err
	}
	return appMessage, nil
}

// FromAppMessage creates a KobradMessage from a appmessage.Message
func FromAppMessage(message appmessage.Message) (*KobradMessage, error) {
	payload, err := toPayload(message)
	if err != nil {
		return nil, err
	}
	return &KobradMessage{
		Payload: payload,
	}, nil
}

func toPayload(message appmessage.Message) (isKobradMessage_Payload, error) {
	p2pPayload, err := toP2PPayload(message)
	if err != nil {
		return nil, err
	}
	if p2pPayload != nil {
		return p2pPayload, nil
	}

	rpcPayload, err := toRPCPayload(message)
	if err != nil {
		return nil, err
	}
	if rpcPayload != nil {
		return rpcPayload, nil
	}

	return nil, errors.Errorf("unknown message type %T", message)
}

func toP2PPayload(message appmessage.Message) (isKobradMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.MsgAddresses:
		payload := new(KobradMessage_Addresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlock:
		payload := new(KobradMessage_Block)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestBlockLocator:
		payload := new(KobradMessage_RequestBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockLocator:
		payload := new(KobradMessage_BlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAddresses:
		payload := new(KobradMessage_RequestAddresses)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDBlocks:
		payload := new(KobradMessage_RequestIBDBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextHeaders:
		payload := new(KobradMessage_RequestNextHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneHeaders:
		payload := new(KobradMessage_DoneHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestRelayBlocks:
		payload := new(KobradMessage_RequestRelayBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestTransactions:
		payload := new(KobradMessage_RequestTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTransactionNotFound:
		payload := new(KobradMessage_TransactionNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvRelayBlock:
		payload := new(KobradMessage_InvRelayBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgInvTransaction:
		payload := new(KobradMessage_InvTransactions)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPing:
		payload := new(KobradMessage_Ping)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPong:
		payload := new(KobradMessage_Pong)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTx:
		payload := new(KobradMessage_Transaction)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVerAck:
		payload := new(KobradMessage_Verack)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgVersion:
		payload := new(KobradMessage_Version)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReject:
		payload := new(KobradMessage_Reject)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointUTXOSet:
		payload := new(KobradMessage_RequestPruningPointUTXOSet)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointUTXOSetChunk:
		payload := new(KobradMessage_PruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgUnexpectedPruningPoint:
		payload := new(KobradMessage_UnexpectedPruningPoint)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocator:
		payload := new(KobradMessage_IbdBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHash:
		payload := new(KobradMessage_IbdBlockLocatorHighestHash)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlockLocatorHighestHashNotFound:
		payload := new(KobradMessage_IbdBlockLocatorHighestHashNotFound)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockHeadersMessage:
		payload := new(KobradMessage_BlockHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointUTXOSetChunk:
		payload := new(KobradMessage_RequestNextPruningPointUtxoSetChunk)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDonePruningPointUTXOSetChunks:
		payload := new(KobradMessage_DonePruningPointUtxoSetChunks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedData:
		payload := new(KobradMessage_BlockWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointAndItsAnticone:
		payload := new(KobradMessage_RequestPruningPointAndItsAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgDoneBlocksWithTrustedData:
		payload := new(KobradMessage_DoneBlocksWithTrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDBlock:
		payload := new(KobradMessage_IbdBlock)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestHeaders:
		payload := new(KobradMessage_RequestHeaders)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPoints:
		payload := new(KobradMessage_PruningPoints)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestPruningPointProof:
		payload := new(KobradMessage_RequestPruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgPruningPointProof:
		payload := new(KobradMessage_PruningPointProof)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgReady:
		payload := new(KobradMessage_Ready)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgTrustedData:
		payload := new(KobradMessage_TrustedData)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgBlockWithTrustedDataV4:
		payload := new(KobradMessage_BlockWithTrustedDataV4)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestNextPruningPointAndItsAnticoneBlocks:
		payload := new(KobradMessage_RequestNextPruningPointAndItsAnticoneBlocks)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestIBDChainBlockLocator:
		payload := new(KobradMessage_RequestIBDChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgIBDChainBlockLocator:
		payload := new(KobradMessage_IbdChainBlockLocator)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.MsgRequestAnticone:
		payload := new(KobradMessage_RequestAnticone)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}

func toRPCPayload(message appmessage.Message) (isKobradMessage_Payload, error) {
	switch message := message.(type) {
	case *appmessage.GetCurrentNetworkRequestMessage:
		payload := new(KobradMessage_GetCurrentNetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCurrentNetworkResponseMessage:
		payload := new(KobradMessage_GetCurrentNetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockRequestMessage:
		payload := new(KobradMessage_SubmitBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitBlockResponseMessage:
		payload := new(KobradMessage_SubmitBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateRequestMessage:
		payload := new(KobradMessage_GetBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockTemplateResponseMessage:
		payload := new(KobradMessage_GetBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedRequestMessage:
		payload := new(KobradMessage_NotifyBlockAddedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyBlockAddedResponseMessage:
		payload := new(KobradMessage_NotifyBlockAddedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BlockAddedNotificationMessage:
		payload := new(KobradMessage_BlockAddedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesRequestMessage:
		payload := new(KobradMessage_GetPeerAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetPeerAddressesResponseMessage:
		payload := new(KobradMessage_GetPeerAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashRequestMessage:
		payload := new(KobradMessage_GetSelectedTipHashRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSelectedTipHashResponseMessage:
		payload := new(KobradMessage_GetSelectedTipHashResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryRequestMessage:
		payload := new(KobradMessage_GetMempoolEntryRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntryResponseMessage:
		payload := new(KobradMessage_GetMempoolEntryResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoRequestMessage:
		payload := new(KobradMessage_GetConnectedPeerInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetConnectedPeerInfoResponseMessage:
		payload := new(KobradMessage_GetConnectedPeerInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerRequestMessage:
		payload := new(KobradMessage_AddPeerRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.AddPeerResponseMessage:
		payload := new(KobradMessage_AddPeerResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionRequestMessage:
		payload := new(KobradMessage_SubmitTransactionRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.SubmitTransactionResponseMessage:
		payload := new(KobradMessage_SubmitTransactionResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedRequestMessage:
		payload := new(KobradMessage_NotifyVirtualSelectedParentChainChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentChainChangedResponseMessage:
		payload := new(KobradMessage_NotifyVirtualSelectedParentChainChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentChainChangedNotificationMessage:
		payload := new(KobradMessage_VirtualSelectedParentChainChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockRequestMessage:
		payload := new(KobradMessage_GetBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockResponseMessage:
		payload := new(KobradMessage_GetBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkRequestMessage:
		payload := new(KobradMessage_GetSubnetworkRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetSubnetworkResponseMessage:
		payload := new(KobradMessage_GetSubnetworkResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockRequestMessage:
		payload := new(KobradMessage_GetVirtualSelectedParentChainFromBlockRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentChainFromBlockResponseMessage:
		payload := new(KobradMessage_GetVirtualSelectedParentChainFromBlockResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksRequestMessage:
		payload := new(KobradMessage_GetBlocksRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlocksResponseMessage:
		payload := new(KobradMessage_GetBlocksResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountRequestMessage:
		payload := new(KobradMessage_GetBlockCountRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockCountResponseMessage:
		payload := new(KobradMessage_GetBlockCountResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoRequestMessage:
		payload := new(KobradMessage_GetBlockDagInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBlockDAGInfoResponseMessage:
		payload := new(KobradMessage_GetBlockDagInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictRequestMessage:
		payload := new(KobradMessage_ResolveFinalityConflictRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ResolveFinalityConflictResponseMessage:
		payload := new(KobradMessage_ResolveFinalityConflictResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsRequestMessage:
		payload := new(KobradMessage_NotifyFinalityConflictsRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyFinalityConflictsResponseMessage:
		payload := new(KobradMessage_NotifyFinalityConflictsResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictNotificationMessage:
		payload := new(KobradMessage_FinalityConflictNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.FinalityConflictResolvedNotificationMessage:
		payload := new(KobradMessage_FinalityConflictResolvedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesRequestMessage:
		payload := new(KobradMessage_GetMempoolEntriesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesResponseMessage:
		payload := new(KobradMessage_GetMempoolEntriesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownRequestMessage:
		payload := new(KobradMessage_ShutDownRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.ShutDownResponseMessage:
		payload := new(KobradMessage_ShutDownResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersRequestMessage:
		payload := new(KobradMessage_GetHeadersRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetHeadersResponseMessage:
		payload := new(KobradMessage_GetHeadersResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedRequestMessage:
		payload := new(KobradMessage_NotifyUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyUTXOsChangedResponseMessage:
		payload := new(KobradMessage_NotifyUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UTXOsChangedNotificationMessage:
		payload := new(KobradMessage_UtxosChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedRequestMessage:
		payload := new(KobradMessage_StopNotifyingUtxosChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingUTXOsChangedResponseMessage:
		payload := new(KobradMessage_StopNotifyingUtxosChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesRequestMessage:
		payload := new(KobradMessage_GetUtxosByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetUTXOsByAddressesResponseMessage:
		payload := new(KobradMessage_GetUtxosByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressRequestMessage:
		payload := new(KobradMessage_GetBalanceByAddressRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalanceByAddressResponseMessage:
		payload := new(KobradMessage_GetBalanceByAddressResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreRequestMessage:
		payload := new(KobradMessage_GetVirtualSelectedParentBlueScoreRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetVirtualSelectedParentBlueScoreResponseMessage:
		payload := new(KobradMessage_GetVirtualSelectedParentBlueScoreResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedRequestMessage:
		payload := new(KobradMessage_NotifyVirtualSelectedParentBlueScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualSelectedParentBlueScoreChangedResponseMessage:
		payload := new(KobradMessage_NotifyVirtualSelectedParentBlueScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualSelectedParentBlueScoreChangedNotificationMessage:
		payload := new(KobradMessage_VirtualSelectedParentBlueScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanRequestMessage:
		payload := new(KobradMessage_BanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.BanResponseMessage:
		payload := new(KobradMessage_BanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanRequestMessage:
		payload := new(KobradMessage_UnbanRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.UnbanResponseMessage:
		payload := new(KobradMessage_UnbanResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoRequestMessage:
		payload := new(KobradMessage_GetInfoRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetInfoResponseMessage:
		payload := new(KobradMessage_GetInfoResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideRequestMessage:
		payload := new(KobradMessage_NotifyPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyPruningPointUTXOSetOverrideResponseMessage:
		payload := new(KobradMessage_NotifyPruningPointUTXOSetOverrideResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.PruningPointUTXOSetOverrideNotificationMessage:
		payload := new(KobradMessage_PruningPointUTXOSetOverrideNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.StopNotifyingPruningPointUTXOSetOverrideRequestMessage:
		payload := new(KobradMessage_StopNotifyingPruningPointUTXOSetOverrideRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondRequestMessage:
		payload := new(KobradMessage_EstimateNetworkHashesPerSecondRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.EstimateNetworkHashesPerSecondResponseMessage:
		payload := new(KobradMessage_EstimateNetworkHashesPerSecondResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedRequestMessage:
		payload := new(KobradMessage_NotifyVirtualDaaScoreChangedRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyVirtualDaaScoreChangedResponseMessage:
		payload := new(KobradMessage_NotifyVirtualDaaScoreChangedResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.VirtualDaaScoreChangedNotificationMessage:
		payload := new(KobradMessage_VirtualDaaScoreChangedNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesRequestMessage:
		payload := new(KobradMessage_GetBalancesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetBalancesByAddressesResponseMessage:
		payload := new(KobradMessage_GetBalancesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateRequestMessage:
		payload := new(KobradMessage_NotifyNewBlockTemplateRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NotifyNewBlockTemplateResponseMessage:
		payload := new(KobradMessage_NotifyNewBlockTemplateResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.NewBlockTemplateNotificationMessage:
		payload := new(KobradMessage_NewBlockTemplateNotification)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesRequestMessage:
		payload := new(KobradMessage_GetMempoolEntriesByAddressesRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetMempoolEntriesByAddressesResponseMessage:
		payload := new(KobradMessage_GetMempoolEntriesByAddressesResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyRequestMessage:
		payload := new(KobradMessage_GetCoinSupplyRequest)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	case *appmessage.GetCoinSupplyResponseMessage:
		payload := new(KobradMessage_GetCoinSupplyResponse)
		err := payload.fromAppMessage(message)
		if err != nil {
			return nil, err
		}
		return payload, nil
	default:
		return nil, nil
	}
}
