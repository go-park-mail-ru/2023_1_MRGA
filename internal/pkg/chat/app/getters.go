package app

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/chat/app/constants"
	chatpc "github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/chat"
)

func GetGRPCInitialChatData(data CreateChatRequest) *chatpc.CreateChatRequest {
	var grpcUserIds []uint32
	for _, userId := range data.UserIds {
		grpcUserIds = append(grpcUserIds, uint32(userId))
	}

	return &chatpc.CreateChatRequest{
		UserIds: grpcUserIds,
	}
}

func GetCreatedChatDataStruct(data *chatpc.CreateChatResponse) CreateChatResponse {
	return CreateChatResponse{
		ChatId: uint(data.GetChatId()),
	}
}

func GetGRPCChatMessage(data InitialMessageData, chatId uint) *chatpc.SendMessageRequest {
	return &chatpc.SendMessageRequest{
		Msg: &chatpc.Message{
			SenderId:   uint32(data.SenderId),
			Content:    data.Content,
			SentAt:     timestamppb.New(data.SentAt),
			ReadStatus: data.ReadStatus,
		},
		ChatId:      uint32(chatId),
		MessageType: string(data.MessageType),
		Path:        data.Path,
	}
}

func GetChatMessageStruct(data *chatpc.GetChatsListResponse) ChatMessage {
	var chatUserIds []uint
	for _, chatUserId := range data.GetChatUserIds() {
		chatUserIds = append(chatUserIds, uint(chatUserId))
	}

	return ChatMessage{
		Msg: MessageResponse{
			SenderId:    uint(data.GetMsg().GetSenderId()),
			Content:     data.GetMsg().GetContent(),
			SentAt:      data.GetMsg().GetSentAt().AsTime().Local().Format("15:04 02.01.2006"),
			ReadStatus:  data.GetMsg().GetReadStatus(),
			MessageType: constants.MessageType(data.GetMessageType()),
			Path:        data.GetPath(),
		},
		ChatId:      uint(data.GetChatId()),
		ChatUserIds: chatUserIds,
	}
}

func GetMessageDataStruct(data *chatpc.GetChatResponse) MessageData {
	return MessageData{
		Msg: MessageResponseWithId{
			SenderId:    uint(data.GetMsg().GetSenderId()),
			Content:     data.GetMsg().GetContent(),
			SentAt:      data.GetMsg().GetSentAt().AsTime().Local().Format("15:04 02.01.2006"),
			ReadStatus:  data.GetMsg().GetReadStatus(),
			MessageType: constants.MessageType(data.GetMessageType()),
			Path:        data.GetPath(),
			MsgId:       uint(data.GetMsgId()),
		},
	}
}
