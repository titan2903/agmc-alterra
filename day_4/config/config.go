package config

import (
	"day_4/models"
	repositories "day_4/repositories"
	"day_4/services"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	c            services.Services
	onController sync.Once
)

func GetController() services.Services {
	onController.Do(func() {
		c = services.NewServices(GetRepository())
	})

	return c
}

var (
	repo    repositories.Repositories
	oneRepo sync.Once
)

func GetRepository() repositories.Repositories {
	oneRepo.Do(func() {
		repo = repositories.NewRepositories(GetQuery())
	})

	return repo
}

var (
	qry     *gorm.DB
	oneSync sync.Once
)

func GetQuery() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	oneSync.Do(func() {
		dbLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: time.Second, // Slow SQL threshold
				LogLevel:      logger.Info, // Log level
				Colorful:      true,        // Disable color
			},
		)

		gormConfig := &gorm.Config{
			// enhance performance config
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
			Logger:                 dbLogger,
		}

		dsnMaster := os.Getenv("DB_DSN")
		dbMaster, errMaster := gorm.Open(mysql.Open(dsnMaster), gormConfig)
		if errMaster != nil {
			log.Panic(errMaster)
		}

		var err error
		if err = dbMaster.AutoMigrate(
			&models.User{},
			&models.Book{},
		); err != nil {
			log.Fatal(err)
		}

		fmt.Println("success connect to database")
		qry = dbMaster
	})

	return qry
}
