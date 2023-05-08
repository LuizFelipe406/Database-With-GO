package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LuizFelipe406/Database-With-GO/configs"
	"github.com/LuizFelipe406/Database-With-GO/repositories"
)

func main() {
	repo := repositories.New(repositories.Options{
		WriterSqlx: configs.GetWriterSqlx(),
		ReaderSqlx: configs.GetReaderSqlx(),
	})

	users, err := repo.User.GetAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
