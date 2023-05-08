package users

import (
	"context"

	"github.com/LuizFelipe406/Database-With-GO/entities"
	"github.com/jmoiron/sqlx"
)

type userDbSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlx(writer *sqlx.DB, reader *sqlx.DB) UserRepositoryInterface {
	return &userDbSqlx{writer, reader}
}

func (u *userDbSqlx) GetAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User

	err := u.reader.SelectContext(ctx, &users, `
		SELECT * FROM users
	`)

	if err != nil {
		return nil, err
	}

	return users, nil
}
