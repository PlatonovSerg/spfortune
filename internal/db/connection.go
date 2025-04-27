package db

import (
	`fmt`
	`log`

	`gorm.io/driver/sqlite`
	`gorm.io/gorm`
	`gorm.io/gorm/logger`
)

var DB *gorm.DB

func InitDB(dsn string) {
	var err error
	DB, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	fmt.Println("Connected to database")

	if err := RunMigrations(); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
}
