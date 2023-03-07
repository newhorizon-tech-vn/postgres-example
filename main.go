package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBClient *gorm.DB

func initPostgresClient(username, password, host string, port int, database string) (*gorm.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", username, password, host, port, database)
	return gorm.Open(postgres.Open(dbURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}

func main() {
	username := "postgres_user_name"
	password := "mysecretpassword"
	host := "127.0.0.1"
	port := 5432
	database := "test_postgres_db"

	var err error
	DBClient, err = initPostgresClient(username, password, host, port, database)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		err = DBClient.AutoMigrate(&Category{})
		if err != nil {
			log.Fatalln(err)
		}

		err = DBClient.AutoMigrate(&Product{})
		if err != nil {
			log.Fatalln(err)
		}

		err = DBClient.AutoMigrate(&Item{})
		if err != nil {
			log.Fatalln(err)
		}

		err = DBClient.AutoMigrate(&Factory{})
		if err != nil {
			log.Fatalln(err)
		}

		err = DBClient.AutoMigrate(&ProductFactory{})
		if err != nil {
			log.Fatalln(err)
		}

		err = DBClient.AutoMigrate(&Workshop{})
		if err != nil {
			log.Fatalln(err)
		}
	*/

	// testHasMany()

	testManyToMany()
}
