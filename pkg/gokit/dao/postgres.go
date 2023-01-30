package dao

import (
	"context"
	"database/sql"
	"gokit-example/pkg/gokit/model"
)

type postgresClient struct {
	db *sql.DB
}

func NewPostgresClient(db *sql.DB) UserSvcDao {
	return &postgresClient{
		db: db,
	}
}

func (p *postgresClient) CreateUser(ctx context.Context, user *model.User) error {
	sqlStatement := `
	INSERT INTO users (user_id, first_name, last_name, email, contact, password)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := p.db.ExecContext(ctx, sqlStatement, user.ID, user.FirstName, user.LastName, user.Email, user.Contact, user.Password)
	if err != nil {
		return err
	}

	return nil
}
