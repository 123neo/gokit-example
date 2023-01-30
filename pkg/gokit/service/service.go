package service

import (
	"context"
	"database/sql"
	"errors"
	"gokit-example/pkg/gokit/dao"
	"gokit-example/pkg/gokit/model"

	"github.com/gofrs/uuid"
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
	pg     dao.UserSvcDao
}

func NewUsrService(logger *zap.Logger, conn *sql.DB) UserService {
	return &userServiceHandler{
		logger: logger,
		pg:     dao.NewPostgresClient(conn),
	}
}

func (usr *userServiceHandler) CreateUser(ctx context.Context, user *model.User) (string, error) {
	if _, errValidate := model.ValidateUser(user); errValidate != nil {
		usr.logger.Error("Validate Error", zap.Error(errValidate))
		return "", errValidate
	}

	uuid, err := uuid.NewV4()
	if err != nil {
		usr.logger.Warn("Error in generating uuid..")
	}
	id := uuid.String()
	user.ID = id

	// encrypting password

	if hashString, err := HashPassword(user.Password); err != nil {
		usr.logger.Error("Not able to encrypt password", zap.Error(err))
	} else {
		user.Password = hashString
	}

	if err := usr.pg.CreateUser(ctx, user); err != nil {
		usr.logger.Error("Create User Error:", zap.Error(err))
		return "", err
	}
	return user.ID, nil
}
