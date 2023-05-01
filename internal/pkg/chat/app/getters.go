package app

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/proto_services/proto_chat"
)

func GetGRPCInitialChatData(data CreateChatRequest) *chatpc.CreateChatRequest {
	var grpcUserIds []*wrapperspb.UInt32Value
	for _, userId := range data.UserIds {
		grpcUserIds = append(grpcUserIds, wrapperspb.UInt32(uint32(userId)))
	}

	return &chatpc.CreateChatRequest{
		UserIds: grpcUserIds,
	}
}

func GetCreatedChatDataStruct(data *chatpc.CreateChatResponse) CreateChatResponse {
	return CreateChatResponse{
		ChatId: uint(data.GetChatId().GetValue()),
	}
}

func GetGRPCChatMessage(data Message, chatId uint) *chatpc.SendMessageRequest {
	return &chatpc.SendMessageRequest{
		Msg: &chatpc.Message{
			SenderId:   wrapperspb.UInt32(uint32(data.SenderId)),
			Content:    data.Content,
			SentAt:     timestamppb.New(data.SentAt),
			ReadStatus: data.ReadStatus,
		},
		ChatId: wrapperspb.UInt32(uint32(chatId)),
	}
}

func GetChatMessageStruct(data *chatpc.GetChatsListResponse) ChatMessage {
	return ChatMessage{
		Msg: MessageResponse{
			SenderId:   uint(data.GetMsg().GetSenderId().GetValue()),
			Content:    data.GetMsg().GetContent(),
			SentAt:     data.GetMsg().GetSentAt().AsTime().Local().Format("15:04 02.01.2006"),
			ReadStatus: data.GetMsg().GetReadStatus(),
		},
		ChatId: uint(data.GetChatId().GetValue()),
	}
}

func GetMessageDataStruct(data *chatpc.GetChatResponse) MessageData {
	return MessageData{
		Msg: MessageResponse{
			SenderId:   uint(data.GetMsg().GetSenderId().GetValue()),
			Content:    data.GetMsg().GetContent(),
			SentAt:     data.GetMsg().GetSentAt().AsTime().Local().Format("15:04 02.01.2006"),
			ReadStatus: data.GetMsg().GetReadStatus(),
		},
		MsgId: uint(data.GetMsgId().GetValue()),
	}
}
