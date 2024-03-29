package app

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/chat"
)

func GetMessageStruct(data *chatpc.SendMessageRequest) ChatMessage {
	return ChatMessage{
		Message: Message{
			ChatId:   uint(data.GetChatId()),
			SenderId: uint(data.GetMsg().GetSenderId()),
			Content:  data.GetMsg().GetContent(),
			SentAt:   data.GetMsg().GetSentAt().AsTime(),
		},
		MessageType: MessageType(data.GetMessageType()),
		Path:        data.GetPath(),
	}
}

func GetInitialUserStruct(data *chatpc.GetChatsListRequest) GetChatsListRequest {
	return GetChatsListRequest{
		UserId: uint(data.GetUserId()),
	}
}

func GetGrpcChatMessage(data MessageWithChatUsers) *chatpc.GetChatsListResponse {
	var uint32ChatUserIds []uint32
	for _, uintUserId := range data.ChatUserIds {
		uint32ChatUserIds = append(uint32ChatUserIds, uint32(uintUserId))
	}

	return &chatpc.GetChatsListResponse{
		Msg: &chatpc.Message{
			SenderId:   uint32(data.SenderId),
			Content:    data.Content,
			SentAt:     timestamppb.New(data.SentAt),
			ReadStatus: data.ReadStatus,
		},
		ChatId:      uint32(data.ChatId),
		MessageType: string(data.MessageType),
		Path:        data.Path,
		ChatUserIds: uint32ChatUserIds,
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

func GetInitialChatForParticipantsStruct(data *chatpc.GetChatParticipantsRequest) GetChatParticipantsRequest {
	return GetChatParticipantsRequest{
		ChatId: uint(data.GetChatId()),
		UserId: uint(data.GetUserId()),
	}
}

func GetGrpcMessage(data ChatMessage) *chatpc.GetChatResponse {
	return &chatpc.GetChatResponse{
		Msg: &chatpc.Message{
			SenderId:   uint32(data.SenderId),
			Content:    data.Content,
			SentAt:     timestamppb.New(data.SentAt),
			ReadStatus: data.ReadStatus,
		},
		MsgId:       uint32(data.Id),
		MessageType: string(data.MessageType),
		Path:        data.Path,
	}
}
