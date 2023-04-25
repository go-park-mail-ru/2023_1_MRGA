package app

import (
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func GetStructMessage(data *chatpc.Message) Message {
	return Message{
		SenderId:   uint(data.GetSenderId().GetValue()),
		ReceiverId: uint(data.GetReceiverId().GetValue()),
		Content:    data.GetContent(),
		SentAt:     data.GetSentAt().AsTime().Local(),
	}
}

func GetStructResentMessagesRequest(data *chatpc.ResentMessagesRequest) ResentMessagesRequest {
	return ResentMessagesRequest{
		UserId: uint(data.GetUserId().GetValue()),
	}
}

func GetGRPCMessage(data Message) *chatpc.Message {
	return &chatpc.Message{
		SenderId:   wrapperspb.UInt32(uint32(data.SenderId)),
		ReceiverId: wrapperspb.UInt32(uint32(data.ReceiverId)),
		Content:    data.Content,
		SentAt:     timestamppb.New(data.SentAt),
	}
}
