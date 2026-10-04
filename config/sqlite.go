package config

import (
	"os"
	"projeto-teste/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func initSQLite () (*gorm.DB, error) {
	logger := InitLogger("SQLite")
	dbPath := "./database/main.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("creating new database")

		err = os.MkdirAll("./database", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil { 
		logger.ErrorF("open database connection fail: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(
		&model.Category{},
		&model.Subcategory{},
		&model.Product{},
		&model.LegalEntity{},
		&model.NaturalPerson{},
		&model.User{},
		&model.Client{},
		&model.Phone{},
		&model.Adress{},
		&model.SaleItem{},
		&model.Sale{},
		&model.Audit{},
	)

	if err != nil {
		logger.ErrorF("auto migration fail: %v", err)
		return nil, err
	}

	return db, nil
}