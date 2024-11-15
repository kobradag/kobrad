package protowire

import (
	"github.com/kobradag/kobrad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KobradMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KobradMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *KobradMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
