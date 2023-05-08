package users

import (
	"context"
	"errors"
	"fmt"

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

func (u *userSqlxRepo) Create(ctx context.Context, newUser entities.User) error {
	_, err := u.writer.ExecContext(ctx, `
		INSERT INTO users (username, name, email, password) VALUES (?, ?, ?, ?)
	`, newUser.Username, newUser.Name, newUser.Email, newUser.Password)

	if err != nil {
		fmt.Println(err)
		return errors.New("não foi possível cadastrar o usuario")
	}

	return nil
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
