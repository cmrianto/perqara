package usecase

import (
	"context"
	"perqara/domain/response"
)

func (s *service) GetUserContentTypes(ctx context.Context) (*response.GetUserContentTypesResponse, error) {
	var (
		result response.GetUserContentTypesResponse
	)

	// reqId, _ := strconv.ParseInt(req.Id, 10, 64)
	// users, err := s.perqaraRepo.GetUsers(ctx, &payload.GetUsersPayload{})
	// if err != nil {
	// 	return &result, err
	// }

	return &result, nil
}
