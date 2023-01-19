package endpoint

import (
	"gokit-example/pkg/gokit/service"

	"github.com/go-kit/kit/endpoint"
	"go.uber.org/zap"
)

type UserSvcSet struct {
	CreateUserEndpoint endpoint.Endpoint
}

func NewUserSvcSet(svc service.UserService, logger *zap.Logger) UserSvcSet {
	endpointLogger := logger.With(zap.String("location", "endpoint"))
	var createUserEndpoint endpoint.Endpoint
	{
		createUserEndpoint = MakeCreateUserEndpoint(svc)
		createUserEndpoint = LoggingMiddleware(endpointLogger.With(zap.String("method", "CreateUser")))(createUserEndpoint)

	}
	return UserSvcSet{
		CreateUserEndpoint: createUserEndpoint,
	}
}
