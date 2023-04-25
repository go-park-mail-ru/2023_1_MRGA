package service

import (
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/chat/internal/pkg/repository"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func GetGRPCMessage(data repository.Message) *chatpc.Message {
	return &chatpc.Message{
		SenderId:   wrapperspb.UInt32(uint32(data.SenderId)),
		ReceiverId: wrapperspb.UInt32(uint32(data.ReceiverId)),
		Content:    data.Content,
		SentAt:     timestamppb.New(data.SentAt),
	}
}

func GetStructMessage(data *chatpc.Message) repository.Message {
	return repository.Message{
		SenderId:   uint(data.GetSenderId().GetValue()),
		ReceiverId: uint(data.GetReceiverId().GetValue()),
		Content:    data.GetContent(),
		SentAt:     data.GetSentAt().AsTime(),
	}
}
