package database

import (
	"fmt"
	"log"

	"github.com/justin-jiajia/note/config"
	"github.com/justin-jiajia/note/model"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	var db *gorm.DB
	var err error

	switch cfg.DBType {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cfg.DBPath), &gorm.Config{})
	default:
		log.Fatal("Unsupported database type")
	}

	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// Auto migrate the schema
	if err = autoMigrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	return db
}

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Note{},
		&model.NoteHistory{},
	)
}
