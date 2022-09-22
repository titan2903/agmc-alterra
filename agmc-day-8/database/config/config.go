package config

import (
	"agmc-day-8/internal/models"
	repositories "agmc-day-8/internal/repositories"
	"agmc-day-8/internal/services"
	"agmc-day-8/pkg/utils"
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
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
		// repo = repositories.NewRepositories(GetQuery(), ConnectDB())
		repo = repositories.NewRepositories(GetQuery())
	})

	return repo
}

var (
	qry     *gorm.DB
	oneSync sync.Once
)

func GetQuery() *gorm.DB {

	oneSync.Do(func() {
		// ! Connect to MySQL database
		// dbLogger := logger.New(
		// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		// 	logger.Config{
		// 		SlowThreshold: time.Second, // Slow SQL threshold
		// 		LogLevel:      logger.Info, // Log level
		// 		Colorful:      true,        // Disable color
		// 	},
		// )

		// gormConfig := &gorm.Config{
		// 	// enhance performance config
		// 	PrepareStmt:            true,
		// 	SkipDefaultTransaction: true,
		// 	Logger:                 dbLogger,
		// }

		// dsnMaster := utils.GoDotEnvVariable("DB_DSN")
		// dbMaster, errMaster := gorm.Open(mysql.Open(dsnMaster), gormConfig)

		//! ----------------------------------------------------------------------------------------------

		// ! Connect to PostgreSQL database
		dsnMaster := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
			utils.GoDotEnvVariable("DB_HOST"), utils.GoDotEnvVariable("DB_USER"), utils.GoDotEnvVariable("DB_PASSWORD"), utils.GoDotEnvVariable("DB_NAME"), utils.GoDotEnvVariable("DB_PORT"), utils.GoDotEnvVariable("SSL_MODE"),
		)
		dbMaster, errMaster := gorm.Open(postgres.Open(dsnMaster), &gorm.Config{})
		if errMaster != nil {
			log.Panic(errMaster)
		}

		//! ----------------------------------------------------------------------------------------------

		if errMaster = dbMaster.AutoMigrate(
			&models.User{},
			&models.Book{},
		); errMaster != nil {
			log.Fatal(errMaster)
		}

		fmt.Println("success connect to database")
		qry = dbMaster
	})

	return qry
}

// func EnvMongoURI() string {
// 	return utils.GoDotEnvVariable("MONGOURI")
// }

// func ConnectDB() *mongo.Client {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(EnvMongoURI()))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	//ping the database
// 	err = client.Ping(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("Connected to MongoDB")
// 	return client
// }

// //Client instance
// var DB *mongo.Client = ConnectDB()

// //getting database collections
// func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
// 	dbName := utils.GoDotEnvVariable("MONGODB_NAME")
// 	collection := client.Database(dbName).Collection(collectionName)
// 	return collection
// }
