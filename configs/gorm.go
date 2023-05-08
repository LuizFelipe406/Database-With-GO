package configs

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func GetReaderGorm() *gorm.DB {
	reader, err := gorm.Open("mysql", DB_CONNECTION)

	if err != nil {
		panic("failer to connect to database")
	}

	return reader
}

func GetWriterGorm() *gorm.DB {
	reader, err := gorm.Open("mysql", DB_CONNECTION)

	if err != nil {
		panic("failer to connect to database")
	}

	return reader
}
