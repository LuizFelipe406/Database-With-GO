package users

import (
	"context"

	"github.com/LuizFelipe406/Database-With-GO/entities"
	"github.com/jmoiron/sqlx"
)

type userSqlxRepo struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlxRepo(writer, reader *sqlx.DB) UserRepositoryInterface {
	return &userSqlxRepo{writer, reader}
}

func (u *userSqlxRepo) GetAll(ctx context.Context) ([]entities.User, error) {
	var users []entities.User

	err := u.reader.SelectContext(ctx, &users, `
		SELECT * FROM users
	`)

	if err != nil {
		return nil, err
	}

	return users, nil
}
