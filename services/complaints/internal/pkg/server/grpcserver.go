package server

import (
	"context"

	compRepo "github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/internal/pkg"
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/pkg/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaintProto"
	tracejaeger "github.com/go-park-mail-ru/2023_1_MRGA.git/utils/trace_jaeger"
)

type GRPCServer struct {
	CompRepo compRepo.CompRepo
}

func NewGPRCServer(compRepo compRepo.CompRepo) *GRPCServer {
	return &GRPCServer{
		CompRepo: compRepo,
	}
}

func (s *GRPCServer) Complain(ctx context.Context, req *complaintProto.UserId) (*complaintProto.Response, error) {
	_, span := tracejaeger.NewSpan(ctx, "complaintsServer", "Complain", nil)
	defer span.End()

	userId := uint(req.UserId)
	count, err := s.CompRepo.CheckCountComplaint(userId)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		complaint := dataStruct.Complaint{Count: 1, UserId: userId}
		err = s.CompRepo.SaveComplaint(complaint)
		if err != nil {
			return nil, err
		}

		return &complaintProto.Response{
			Banned: false,
		}, nil
	}

	err = s.CompRepo.IncrementComplaint(userId)
	if err != nil {
		return nil, err
	}

	if count == 4 {
		return &complaintProto.Response{
			Banned: true,
		}, nil
	}
	return &complaintProto.Response{
		Banned: false,
	}, nil
}

func (s *GRPCServer) CheckBanned(ctx context.Context, req *complaintProto.UserId) (*complaintProto.Response, error) {
	_, span := tracejaeger.NewSpan(ctx, "complaintsServer", "CheckBanned", nil)
	defer span.End()

	userId := uint(req.UserId)
	count, err := s.CompRepo.CheckCountComplaint(userId)
	if err != nil {
		return nil, err
	}

	if count < 5 {
		return &complaintProto.Response{
			Banned: false,
		}, nil
	}

	return &complaintProto.Response{
		Banned: true,
	}, nil
}
