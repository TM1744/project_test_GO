package config

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitDatabase() (*gorm.DB, error) {
	db, err := initSQLite()

	if err != nil {
		return nil, fmt.Errorf("failed to initialize the SQLite database : %v", err)
	}

	return db, nil
}

func InitLogger(prefix string) *Logger {
	logger := newLogger(prefix)
	return logger
}

func InitGin() *gin.Engine {
	return initGin()
}

func InitTransactionManager(db *gorm.DB, key *string) *TransactionManager {
	return &TransactionManager{Database: db, ContextKey: key}
}
