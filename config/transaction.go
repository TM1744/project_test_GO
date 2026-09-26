package config

import (
	"context"
	"gorm.io/gorm"
)

type TransactionManager struct {
	Database   *gorm.DB
	ContextKey *string
}

func (tm *TransactionManager) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	return tm.Database.Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, tm.ContextKey, tx)
		return fn(txCtx)
	})
}

func (tm *TransactionManager) GetDB(context context.Context) *gorm.DB {
	if tx, ok := context.Value(tm.ContextKey).(*gorm.DB); ok {
		return tx
	}
	return tm.Database
}
