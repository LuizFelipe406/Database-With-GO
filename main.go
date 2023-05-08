package main

import (
	"context"
	"fmt"
	"log"

	"github.com/LuizFelipe406/Database-With-GO/configs"
	"github.com/LuizFelipe406/Database-With-GO/entities"
	"github.com/LuizFelipe406/Database-With-GO/repositories"
)

func main() {
	repo := repositories.New(repositories.Options{
		WriterSqlx: configs.GetWriterSqlx(),
		ReaderSqlx: configs.GetReaderSqlx(),
	})

	err := repo.User.Create(context.Background(), entities.User{
		Username: "Luiz Dev",
		Name:     "Luiz Felipe Pereira",
		Email:    "email@example.com",
		Password: "12345678",
	})

	if err != nil {
		log.Fatal(err)
	}

	users, err := repo.User.GetAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}
