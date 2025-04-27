package db

func RunMigrations() error {
	return DB.AutoMigrate()
}
