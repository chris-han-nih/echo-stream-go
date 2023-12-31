package infrastructure

import (
	"fmt"
	"os"
)

func NewDatabase() (string, string) {
	driver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	return driver,
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
			user,
			password,
			host, port,
			database)
}
