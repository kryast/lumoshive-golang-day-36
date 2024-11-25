package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type AllRepository struct {
	ProductRepo ProductRepoInterface
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		ProductRepo: NewProductRepository(db, log),
	}
}
