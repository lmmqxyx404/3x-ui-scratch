package database

import (
	"io/fs"
	"log"
	"os"
	"path"
	"x-ui-scratch/config"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDB(dbPath string) error {
	dir := path.Dir(dbPath)
	err := os.MkdirAll(dir, fs.ModePerm)
	if err != nil {
		return err
	}

	var gormLogger logger.Interface

	if config.IsDebug() {
		gormLogger = logger.Default
	} else {
		gormLogger = logger.Discard
	}

	c := &gorm.Config{
		Logger: gormLogger,
	}

	db, err = gorm.Open(sqlite.Open(dbPath), c)
	if err != nil {
		return err
	}

	if err := initModels(); err != nil {
		return err
	}
	// TODO
	return nil
}

func initModels() error {
	// TODO: 添加可用的 model
	models := []interface{}{}
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			log.Printf("Error auto migrating model: %v", err)
			return err
		}
	}
	return nil
}
