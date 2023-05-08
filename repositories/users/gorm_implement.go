package users

import (
	"context"
	"errors"
	"fmt"

	"github.com/LuizFelipe406/Database-With-GO/entities"
	"github.com/jinzhu/gorm"
)

type userGormRepo struct {
	writer *gorm.DB
	reader *gorm.DB
}

func NewGormRepo(writer, reader *gorm.DB) UserRepositoryInterface {
	return &userGormRepo{writer, reader}
}

func (r *userGormRepo) Create(ctx context.Context, newUser entities.User) error {
	r.writer.Table("users").Create(&newUser)

	if r.writer.Error != nil {
		fmt.Println(r.writer.Error)
		return errors.New("não foi possível cadastrar o usuario")
	}

	return nil
}

func (repo *userGormRepo) GetAll(ctx context.Context) ([]entities.User, error) {
	var user []entities.User
	repo.reader.Table("users").Find(&user)

	if repo.reader.Error != nil {
		fmt.Println(repo.reader.Error)
		return nil, errors.New("usuário não encontrado")
	}

	return user, nil
}
