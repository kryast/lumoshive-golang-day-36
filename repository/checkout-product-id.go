package repository

import (
	"database/sql"
	"day-36/model"
	"errors"

	"go.uber.org/zap"
)

type ProductRepoInterface interface {
	GetProductById(productId int) (*model.Product, error)
}

type ProductRepositoryDB struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewProductRepository(db *sql.DB, Log *zap.Logger) ProductRepoInterface {
	return &ProductRepositoryDB{
		DB:     db,
		Logger: Log,
	}
}

func (repo *ProductRepositoryDB) GetProductById(productId int) (*model.Product, error) {
	query := `
		SELECT 
			pd.id AS id,
			p.name AS product_name,
			p.detail AS product_detail,
			p.price AS product_price,
			c.name AS category_name,
			ph.photo_url
		FROM 
			product_details pd
		JOIN 
			products p ON p.id = pd.product_id
		JOIN 
			categories c ON c.id = pd.category_id
		JOIN 
			photos ph ON ph.id = pd.photo_id
		WHERE 
			pd.id = $1
	`

	rows := repo.DB.QueryRow(query, productId)

	var product model.Product

	err := rows.Scan(
		&product.ID,
		&product.Name,
		&product.Detail,
		&product.Price,
		&product.Category,
		&product.PhotoURL,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			repo.Logger.Warn("Product not found", zap.Int("productId", productId), zap.Error(err))
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	return &product, nil
}
