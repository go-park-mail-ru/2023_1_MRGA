package app

import (
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func GetMessageStruct(data *chatpc.SendMessageRequest) Message {
	return Message{
		ChatId:   uint(data.GetChatId().GetValue()),
		SenderId: uint(data.GetMsg().GetSenderId().GetValue()),
		Content:  data.GetMsg().GetContent(),
		SentAt:   data.GetMsg().GetSentAt().AsTime().Local(),
	}
}

func GetInitialUserStruct(data *chatpc.GetChatsListRequest) GetChatsListRequest {
	return GetChatsListRequest{
		UserId: uint(data.GetUserId().GetValue()),
	}
}

func GetGrpcChatMessage(data Message) *chatpc.GetChatsListResponse {
	return &chatpc.GetChatsListResponse{
		Msg: &chatpc.Message{
			SenderId:   wrapperspb.UInt32(uint32(data.SenderId)),
			Content:    data.Content,
			SentAt:     timestamppb.New(data.SentAt),
			ReadStatus: data.ReadStatus,
		},
		ChatId: wrapperspb.UInt32(uint32(data.ChatId)),
	}
}

func GetGrpcInitialChatData(data CreateChatResponse) *chatpc.CreateChatResponse {
	return &chatpc.CreateChatResponse{
		ChatId: wrapperspb.UInt32(uint32(data.ChatId)),
	}
}

func GetInitialChatStruct(data *chatpc.GetChatRequest) GetChatRequest {
	return GetChatRequest{
		ChatId: uint(data.GetChatId().GetValue()),
		UserId: uint(data.GetUserId().GetValue()),
	}
}

func GetGrpcMessage(data Message) *chatpc.GetChatResponse {
	return &chatpc.GetChatResponse{
		Msg: &chatpc.Message{
			SenderId:   wrapperspb.UInt32(uint32(data.SenderId)),
			Content:    data.Content,
			SentAt:     timestamppb.New(data.SentAt),
			ReadStatus: data.ReadStatus,
		},
		MsgId: wrapperspb.UInt32(uint32(data.Id)),
	}
}
