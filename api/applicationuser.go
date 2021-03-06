// +build !codeanalysis

package api

import (
	"context"

	applicationuser "github.com/NpoolPlatform/application-management/pkg/crud/application-user"
	appuser "github.com/NpoolPlatform/application-management/pkg/middleware/app-user"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/application"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) AddUsersToApplication(ctx context.Context, in *npool.AddUsersToApplicationRequest) (*npool.AddUsersToApplicationResponse, error) {
	resp, err := applicationuser.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("add user to app error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserFromApplication(ctx context.Context, in *npool.GetUserFromApplicationRequest) (*npool.GetUserFromApplicationResponse, error) {
	resp, err := applicationuser.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user from app error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetUsersFromApplication(ctx context.Context, in *npool.GetUsersFromApplicationRequest) (*npool.GetUsersFromApplicationResponse, error) {
	resp, err := applicationuser.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get users from app error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) RemoveUsersFromApplication(ctx context.Context, in *npool.RemoveUsersFromApplicationRequest) (*npool.RemoveUsersFromApplicationResponse, error) {
	resp, err := applicationuser.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("remove users from app error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) SetGALogin(ctx context.Context, in *npool.SetGALoginRequest) (*npool.SetGALoginResponse, error) {
	resp, err := applicationuser.SetGALogin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("set users user ga login verify: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) SetSMSLogin(ctx context.Context, in *npool.SetSMSLoginRequest) (*npool.SetSMSLoginResponse, error) {
	resp, err := applicationuser.SetSMSLogin(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("set users user sms login verify: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) AddUserLoginTime(ctx context.Context, in *npool.AddUserLoginTimeRequest) (*npool.AddUserLoginTimeResponse, error) {
	resp, err := applicationuser.AddUserLoginTime(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("add user login number error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateUserGAStatus(ctx context.Context, in *npool.UpdateUserGAStatusRequest) (*npool.UpdateUserGAStatusResponse, error) {
	resp, err := applicationuser.UpdateUserGAStatus(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update user ga status error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateUserKYCStatus(ctx context.Context, in *npool.UpdateUserKYCStatusRequest) (*npool.UpdateUserKYCStatusResponse, error) {
	resp, err := applicationuser.UpdateUserKYCStatus(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update user kyc status error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetApplicationUserDetail(ctx context.Context, in *npool.GetApplicationUserDetailRequest) (*npool.GetApplicationUserDetailResponse, error) {
	resp, err := appuser.GetApplicationUserDetail(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get applcation user detail error: %v", err)
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetUserAppID(ctx context.Context, in *npool.GetUserAppIDRequest) (*npool.GetUserAppIDResponse, error) {
	resp, err := applicationuser.GetUserAppID(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get user app id error: %v", err)
		return nil, status.Errorf(codes.FailedPrecondition, "fail to get user app id: %v", err.Error())
	}
	return resp, nil
}
