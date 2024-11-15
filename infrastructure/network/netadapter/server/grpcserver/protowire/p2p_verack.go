package protowire

import (
	"github.com/kobradag/kobrad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KobradMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KobradMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *KobradMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
