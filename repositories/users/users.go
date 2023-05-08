package users

import (
	"context"

	"github.com/LuizFelipe406/Database-With-GO/entities"
)

type UserRepositoryInterface interface {
	GetAll(ctx context.Context) ([]entities.User, error)

	Create(ctx context.Context, newUser entities.User) error
}
