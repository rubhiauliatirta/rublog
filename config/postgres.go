package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "rublog"
)

var Db *gorm.DB

func init() {
	var err error
	var psqlInfo string

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	} else {
		psqlInfo = url
	}

	Db, err = gorm.Open("postgres", psqlInfo)
	//Db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Failed to connect database")
		panic(err)
	}

	fmt.Println("Successfully connected to database!")
}
