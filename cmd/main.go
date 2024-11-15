package main

import (
	"flag"
	"github.com/joho/godotenv"
	"initial-project-go/di"
	config2 "initial-project-go/di/config"
	"initial-project-go/di/database"
	"initial-project-go/entity/migrater"
	"initial-project-go/repository"
	"log"
)

func main() {
	withMigrate := flag.Bool("with-migrate", false, "Run the application with migrations")
	flag.Parse()

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
		panic(err)
	}

	db, err := database.InitDatabase()
	if err != nil {
		panic(err)
	}

	if *withMigrate {
		appConfig := config2.GetConfig()
		err = migrater.AutoMigrate(db)
		if err != nil {
			log.Println("Migration failed", err)
			panic(err)
		}

		encryptorRepository := repository.ProvideEncryptorRepository(db, appConfig)
		_, err := encryptorRepository.GetPassphrase()
		if err != nil {
			log.Println("Failed to get passphrase", err)
			panic(err)
		}
		return
	}

	err = di.InitApplication()
	if err != nil {
		log.Println("Failed to initialize application")
		panic(err)
	}
}
