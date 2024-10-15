package postgres

import (
	"context"
	"fmt"

	"letual/internal/models"
)

func (s *Storage) GetProducts(ctx context.Context) ([]*models.Product, error) {
	const fn = "GetProducts"

	const query = `
		SELECT id, name, description, price, image_url
		FROM products`

	products := make([]*models.Product, 0)

	rows, err := s.client.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.ImageUrl)
		if err != nil {
			return nil, fmt.Errorf("%s : %w", fn, err)
		}

		products = append(products, &product)
	}

	return products, nil
}

func (s *Storage) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	const fn = "GetProduct"

	const query = `
		SELECT id, name, description, price, image_url
		FROM products
		WHERE id = $1`

	var product models.Product

	err := s.client.QueryRow(ctx, query, id).
		Scan(&product.ID, &product.Title, &product.Description, &product.Price, &product.ImageUrl)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	return &product, nil
}

func (s *Storage) SaveProduct(ctx context.Context, product *models.Product) error {
	const fn = "SaveProduct"

	const query = `
		INSERT INTO products (name, description, price, image_url)
		VALUES ($1, $2, $3, $4)`

	_, err := s.client.Exec(ctx, query, product.Title, product.Description, product.Price, product.ImageUrl)
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}

	return nil
}
