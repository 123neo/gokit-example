package service

import (
	"context"
	"database/sql"
	"errors"
	"gokit-example/pkg/gokit/model"

	"go.uber.org/zap"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrCmdRepository     = errors.New("unable to command repository")
	ErrQueryRepository   = errors.New("unable to query repository")
)

type UserService interface {
	CreateUser(ctx context.Context, user *model.User) (string, error)
}

type userServiceHandler struct {
	logger *zap.Logger
	db     *sql.DB
}

func NewUsrService(logger *zap.Logger, conn *sql.DB) UserService {
	return &userServiceHandler{
		logger: logger,
		db:     conn,
	}
}

func (usr *userServiceHandler) CreateUser(ctx context.Context, user *model.User) (string, error) {
	if _, errValidate := model.ValidateUser(user); errValidate != nil {
		usr.logger.Error("Validate Error", zap.Error(errValidate))
		return "", errValidate
	}
	return "", nil
}
