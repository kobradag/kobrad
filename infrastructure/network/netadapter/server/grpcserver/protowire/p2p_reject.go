package protowire

import (
	"github.com/kobradag/kobrad/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KobradMessage_Reject) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "KobradMessage_Reject is nil")
	}
	return x.Reject.toAppMessage()
}

func (x *RejectMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RejectMessage is nil")
	}
	return &appmessage.MsgReject{
		Reason: x.Reason,
	}, nil
}

func (x *KobradMessage_Reject) fromAppMessage(msgReject *appmessage.MsgReject) error {
	x.Reject = &RejectMessage{
		Reason: msgReject.Reason,
	}
	return nil
}
