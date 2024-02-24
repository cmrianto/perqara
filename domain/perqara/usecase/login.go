package usecase

import (
	"context"
	"errors"
	"perqara/domain/payload"
	"perqara/domain/response"
	"perqara/lib/helper"

	"google.golang.org/grpc/codes"
)

func (s *service) Login(ctx context.Context, req *payload.LoginPayload) (*response.LoginResponse, error) {
	var (
		result response.LoginResponse
	)
	if req == nil {
		return &result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}
	if len(req.Username) == 0 && len(req.Password) == 0 {
		return &result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	users, err := s.perqaraRepo.GetUsers(ctx, &payload.GetUsersPayload{
		Username: &req.Username,
	})
	if err != nil {
		return &result, err
	}

	if len(users) == 0 {
		return &result, helper.Error(codes.Unauthenticated, "Unauthenticated: ", errors.New("unauthenticated"))
	}

	user := users[0]

	err = helper.VerifyPassword(req.Password, user.Password)
	if err != nil {
		return &result, err
	}

	var clientId *int64
	if user.ClientId.Valid && user.ClientId.Int64 > 0 {
		clientId = &user.ClientId.Int64
	}

	tokenVal, err := helper.GenerateToken(helper.TokenPayload{
		UserId:   user.Id,
		UserName: user.UserName,
		ClientId: clientId,
		Lifespan: s.config.Auth.Lifespan,
		Secret:   s.config.Auth.ApiSecret,
	})
	if err != nil {
		return &result, err
	}

	result.Token = tokenVal

	return &result, nil
}
