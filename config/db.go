package configs

import (
	"fmt"
	"log"
	"os"

	// "strings"
	"time"
	"txrnxp-whats-happening/internal/database/tables"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	DB *gorm.DB
}

func NewDatabase(env *Config) (*Database, error) {

	// appEnv := strings.ToLower(os.Getenv("ENVIRONMENT"))
	appEnv := env.Environment

	var dsn string
	switch appEnv {
	case "local":
		dsn = fmt.Sprintf(
			"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
			// os.Getenv("DB_HOST"),
			os.Getenv("DB_USER_LOCAL"),
			os.Getenv("DB_PASSWORD_LOCAL"),
			os.Getenv("DB_NAME_LOCAL"),
			// os.Getenv(fmt.Sprintf("DB_USER_%s", os.Getenv("ENVIRONMENT"))),
			// os.Getenv(fmt.Sprintf("DB_PASSWORD_%s", os.Getenv("ENVIRONMENT"))),
			// os.Getenv(fmt.Sprintf("DB_NAME_%s", os.Getenv("ENVIRONMENT"))),
		)
	case "staging":
		dsn = os.Getenv("DATABASE_URL")
		if dsn == "" {
			log.Fatal("DATABASE_URL is not set")
			os.Exit(2)
		}
	default:
		dsn = os.Getenv("DATABASE_URL_PROD")
		if dsn == "" {
			log.Fatal("DATABASE_URL is not set")
			os.Exit(2)
		}
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:      logger.Default.LogMode(logger.Info),
		PrepareStmt: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		os.Exit(2)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to configure database connection: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	db.Logger = logger.Default.LogMode(logger.Error)

	// DB := Dbinstance{
	// 	Db: db,
	// }

	log.Println("running migrations")
	db.AutoMigrate(&tables.WhatsHappening{})

	return &Database{DB: db}, nil

	// dsn := fmt.Sprintf(
	// 	"host=db user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
	// 	env.DBUser, env.DBPassword, env.DBName, env.DBPort,
	// )

	// // initialize GORM with custom configurations
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to connect to database: %w", err)
	// }

	// configure connection pooling
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to configure database connection: %w", err)
	// }
	// sqlDB.SetMaxIdleConns(10)
	// sqlDB.SetMaxOpenConns(100)
	// sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// log.Println("Database connection established")
	// return &Database{DB: db}, nil
}

func (db *Database) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get database connection: %w", err)
	}
	return sqlDB.Close()
}
