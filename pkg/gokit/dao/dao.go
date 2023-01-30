package dao

import (
	"context"
	"gokit-example/pkg/gokit/model"
)

type UserSvcDao interface {
	CreateUser(ctx context.Context, user *model.User) error
}
