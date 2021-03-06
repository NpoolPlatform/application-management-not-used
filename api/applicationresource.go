// +build !codeanalysis

package api

import (
	"context"

	applicationresource "github.com/NpoolPlatform/application-management/pkg/crud/application-resource"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/application"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateResource(ctx context.Context, in *npool.CreateResourceRequest) (*npool.CreateResourceResponse, error) {
	if in.Info == nil {
		logger.Sugar().Errorf("invalid input params")
		return &npool.CreateResourceResponse{}, status.Errorf(codes.InvalidArgument, "invalid input params")
	}
	resp, err := applicationresource.Create(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("create resource error: %v", err)
		return &npool.CreateResourceResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetResource(ctx context.Context, in *npool.GetResourceRequest) (*npool.GetResourceResponse, error) {
	resp, err := applicationresource.Get(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get resource error: %v", err)
		return &npool.GetResourceResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetResources(ctx context.Context, in *npool.GetResourcesRequest) (*npool.GetResourcesResponse, error) {
	resp, err := applicationresource.GetAll(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get resources error: %v", err)
		return &npool.GetResourcesResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) UpdateResource(ctx context.Context, in *npool.UpdateResourceRequest) (*npool.UpdateResourceResponse, error) {
	if in.Info == nil {
		logger.Sugar().Errorf("invalid input params")
		return &npool.UpdateResourceResponse{}, status.Errorf(codes.InvalidArgument, "invalid input params")
	}
	resp, err := applicationresource.Update(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("update resource error: %v", err)
		return &npool.UpdateResourceResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) DeleteResource(ctx context.Context, in *npool.DeleteResourceRequest) (*npool.DeleteResourceResponse, error) {
	resp, err := applicationresource.Delete(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("delete resource error: %v", err)
		return &npool.DeleteResourceResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetResourceByCreator(ctx context.Context, in *npool.GetResourceByCreatorRequest) (*npool.GetResourceByCreatorResponse, error) {
	resp, err := applicationresource.GetResourceByCreator(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get resource by creator error: %v", err)
		return &npool.GetResourceByCreatorResponse{}, status.Errorf(codes.Internal, "internal server error: %v", err.Error())
	}
	return resp, nil
}

func (s *Server) GetResourceByName(ctx context.Context, in *npool.GetResourceByNameRequest) (*npool.GetResourceByNameResponse, error) {
	resp, err := applicationresource.GetResourceByName(ctx, in)
	if err != nil {
		logger.Sugar().Errorf("get resource by resource name error: %v", err)
		return &npool.GetResourceByNameResponse{}, status.Errorf(codes.FailedPrecondition, "internal server error: %v", err.Error())
	}
	return resp, nil
}
