package container

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPostgres() error {
	dsn := "host=localhost user=postgres password=manab123 dbname=nebulo_storage port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to DB: %w", err)
	}
	return nil
}
