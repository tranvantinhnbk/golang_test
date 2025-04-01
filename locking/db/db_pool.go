package db

import (
	"fmt"
	"sync"
	"time"

	"golang_test/locking/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DBConfig holds database configuration
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

// DBConnection represents a database connection pool
type DBConnection struct {
	DB *gorm.DB
}

var (
	dbInstance *DBConnection
	once       sync.Once
)

// GetDBInstance returns a singleton database connection
func GetDBInstance(config DBConfig) (*DBConnection, error) {
	var err error
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			config.Host, config.Port, config.User, config.Password, config.DBName)

		// Configure GORM logger
		gormConfig := &gorm.Config{
			Logger: logger.Default.LogMode(logger.Error),
		}

		// Open database connection with connection pool settings
		db, err := gorm.Open(postgres.Open(dsn), gormConfig)
		if err != nil {
			return
		}

		// Get the underlying *sql.DB
		sqlDB, err := db.DB()
		if err != nil {
			return
		}

		// Set connection pool settings
		sqlDB.SetMaxIdleConns(10)           // Maximum number of idle connections
		sqlDB.SetMaxOpenConns(100)          // Maximum number of open connections
		sqlDB.SetConnMaxLifetime(time.Hour) // Maximum lifetime of a connection

		dbInstance = &DBConnection{DB: db}
	})

	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	return dbInstance, nil
}

// InitSchema initializes the database schema
func (db *DBConnection) InitSchema() error {
	// Auto migrate the Account model
	err := db.DB.AutoMigrate(&models.Account{})
	if err != nil {
		return fmt.Errorf("failed to migrate database schema: %v", err)
	}
	return nil
}

// Close closes the database connection
func (db *DBConnection) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get underlying *sql.DB: %v", err)
	}
	return sqlDB.Close()
}
