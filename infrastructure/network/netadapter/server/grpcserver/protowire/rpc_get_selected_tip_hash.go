package protowire

import (
	"github.com/kobradag/kobrad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KobradMessage_GetSelectedTipHashRequest) toAppMessage() (appmessage.Message, error) {
	return &appmessage.GetSelectedTipHashRequestMessage{}, nil
}

func (x *KobradMessage_GetSelectedTipHashRequest) fromAppMessage(_ *appmessage.GetSelectedTipHashRequestMessage) error {
	return nil
}

func (x *KobradMessage_GetSelectedTipHashResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KobradMessage_GetSelectedTipHashResponse is nil")
	}
	return x.GetSelectedTipHashResponse.toAppMessage()
}

func (x *KobradMessage_GetSelectedTipHashResponse) fromAppMessage(message *appmessage.GetSelectedTipHashResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.GetSelectedTipHashResponse = &GetSelectedTipHashResponseMessage{
		SelectedTipHash: message.SelectedTipHash,
		Error:           err,
	}
	return nil
}

func (x *GetSelectedTipHashResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetSelectedTipHashResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	if rpcErr != nil && len(x.SelectedTipHash) != 0 {
		return nil, errors.New("GetSelectedTipHashResponseMessage contains both an error and a response")
	}

	return &appmessage.GetSelectedTipHashResponseMessage{
		SelectedTipHash: x.SelectedTipHash,
		Error:           rpcErr,
	}, nil
}
