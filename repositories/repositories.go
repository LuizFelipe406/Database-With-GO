package repositories

import (
	"github.com/LuizFelipe406/Database-With-GO/repositories/users"
	"github.com/jinzhu/gorm"
	"github.com/jmoiron/sqlx"
)

type Options struct {
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
	WriterGorm *gorm.DB
	ReaderGorm *gorm.DB
}

type Container struct {
	User users.UserRepositoryInterface
}

func New(options Options) *Container {
	return &Container{
		User: users.NewSqlxRepo(options.WriterSqlx, options.ReaderSqlx), // Modify to users.NewGormRepo(options.WriterGorm, options.ReaderGorm) to change between implementations
	}
}
