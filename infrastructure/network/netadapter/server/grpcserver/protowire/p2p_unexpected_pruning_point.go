package protowire

import "github.com/kobradag/kobrad/app/appmessage"

func (x *HarbidMessage_UnexpectedPruningPoint) toAppMessage() (appmessage.Message, error) {
	return &appmessage.MsgUnexpectedPruningPoint{}, nil
}

func (x *HarbidMessage_UnexpectedPruningPoint) fromAppMessage(_ *appmessage.MsgUnexpectedPruningPoint) error {
	return nil
}
