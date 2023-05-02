package app

import (
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GetMessageStruct(data *chatpc.SendMessageRequest) Message {
	return Message{
		ChatId:   uint(data.GetChatId()),
		SenderId: uint(data.GetMsg().GetSenderId()),
		Content:  data.GetMsg().GetContent(),
		SentAt:   data.GetMsg().GetSentAt().AsTime().Local(),
	}
}

func GetInitialUserStruct(data *chatpc.GetChatsListRequest) GetChatsListRequest {
	return GetChatsListRequest{
		UserId: uint(data.GetUserId()),
	}
}

func GetGrpcChatMessage(data Message) *chatpc.GetChatsListResponse {
	return &chatpc.GetChatsListResponse{
		Msg: &chatpc.Message{
			SenderId:   uint32(data.SenderId),
			Content:    data.Content,
			SentAt:     timestamppb.New(data.SentAt),
			ReadStatus: data.ReadStatus,
		},
		ChatId: uint32(data.ChatId),
	}
}

func GetGrpcInitialChatData(data CreateChatResponse) *chatpc.CreateChatResponse {
	return &chatpc.CreateChatResponse{
		ChatId: uint32(data.ChatId),
	}
}

func GetInitialChatStruct(data *chatpc.GetChatRequest) GetChatRequest {
	return GetChatRequest{
		ChatId: uint(data.GetChatId()),
		UserId: uint(data.GetUserId()),
	}
}

func GetGrpcMessage(data Message) *chatpc.GetChatResponse {
	return &chatpc.GetChatResponse{
		Msg: &chatpc.Message{
			SenderId:   uint32(data.SenderId),
			Content:    data.Content,
			SentAt:     timestamppb.New(data.SentAt),
			ReadStatus: data.ReadStatus,
		},
		MsgId: uint32(data.Id),
	}
}