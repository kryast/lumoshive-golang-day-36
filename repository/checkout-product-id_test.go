package repository

import (
	"database/sql"
	"day-36/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetProductById(t *testing.T) {
	logger := zap.NewNop()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to create mock database: %v", err)
	}
	defer db.Close()

	repo := NewProductRepository(db, logger)

	expectedProduct := &model.Product{
		ID:       1,
		Name:     "Test Product",
		Detail:   "This is a test product",
		Price:    100.0,
		Category: "Electronics",
		PhotoURL: "http://example.com/product.jpg",
	}

	t.Run("Success", func(t *testing.T) {
		mock.ExpectQuery(`SELECT 
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
			pd.id = \$1`).
			WithArgs(1).
			WillReturnRows(sqlmock.NewRows([]string{"id", "product_name", "product_detail", "product_price", "category_name", "photo_url"}).
				AddRow(expectedProduct.ID, expectedProduct.Name, expectedProduct.Detail, expectedProduct.Price, expectedProduct.Category, expectedProduct.PhotoURL))

		product, err := repo.GetProductById(1)

		assert.NoError(t, err)
		assert.NotNil(t, product)
		assert.Equal(t, expectedProduct.Name, product.Name)
		assert.Equal(t, expectedProduct.Detail, product.Detail)
		assert.Equal(t, expectedProduct.Price, product.Price)
		assert.Equal(t, expectedProduct.Category, product.Category)
		assert.Equal(t, expectedProduct.PhotoURL, product.PhotoURL)
	})

	t.Run("ProductNotFound", func(t *testing.T) {
		mock.ExpectQuery(`SELECT 
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
			pd.id = \$1`).
			WithArgs(999).
			WillReturnError(sql.ErrNoRows)

		product, err := repo.GetProductById(999)

		assert.Error(t, err)
		assert.Nil(t, product)
		assert.Equal(t, "product not found", err.Error())

	})

}
