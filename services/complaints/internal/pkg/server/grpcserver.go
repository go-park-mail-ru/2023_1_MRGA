package server

import (
	"context"

	compRepo "github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/internal/pkg"
	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/services/complaints/pkg/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/services/proto/complaints"
)

type GRPCServer struct {
	CompRepo compRepo.CompRepo
}

func NewGPRCServer(compRepo compRepo.CompRepo) *GRPCServer {
	return &GRPCServer{
		CompRepo: compRepo,
	}
}

func (s *GRPCServer) Complain(ctx context.Context, req *complaints.UserId) (*complaints.Response, error) {
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

		return &complaints.Response{
			Banned: false,
		}, nil
	}

	err = s.CompRepo.IncrementComplaint(userId)
	if err != nil {
		return nil, err
	}

	if count == 4 {
		return &complaints.Response{
			Banned: true,
		}, nil
	}
	return &complaints.Response{
		Banned: false,
	}, nil
}

func (s *GRPCServer) CheckBanned(ctx context.Context, req *complaints.UserId) (*complaints.Response, error) {
	userId := uint(req.UserId)
	count, err := s.CompRepo.CheckCountComplaint(userId)
	if err != nil {
		return nil, err
	}

	if count < 5 {
		return &complaints.Response{
			Banned: false,
		}, nil
	}

	return &complaints.Response{
		Banned: true,
	}, nil
}
