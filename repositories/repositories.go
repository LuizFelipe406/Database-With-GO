package repositories

import (
	"github.com/LuizFelipe406/Database-With-GO/repositories/users"
	"github.com/jmoiron/sqlx"
)

type Options struct {
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
}

type Container struct {
	User users.UserRepositoryInterface
}

func New(options Options) *Container {
	return &Container{
		User: users.NewSqlx(options.WriterSqlx, options.ReaderSqlx),
	}
}
