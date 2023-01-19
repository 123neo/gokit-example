package endpoint

import (
	"context"
	"gokit-example/pkg/gokit/model"
	"gokit-example/pkg/gokit/service"

	"github.com/go-kit/kit/endpoint"
)

type CreateUserRequest struct {
	User *model.User `json:"user"`
}

type CreateUserResponse struct {
	Key string      `json:"userKey"`
	Err interface{} `json:"error"`
}

// MakeCreateuserEndpoint constructs a create user endpoint wrapping the service
func MakeCreateUserEndpoint(svc service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		res, err := svc.CreateUser(ctx, req.User)
		return CreateUserResponse{res, err}, nil
	}
}
