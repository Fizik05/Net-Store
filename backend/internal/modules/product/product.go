package product

import (
	"context"
	"fmt"
	"log/slog"

	"letual/internal/models"
)

type Storage interface {
	GetProducts(ctx context.Context) ([]*models.Product, error)
	SaveProduct(ctx context.Context, product *models.Product) error
	GetProduct(ctx context.Context, id int) (*models.Product, error)
}

type Product struct {
	ctx     context.Context
	logger  *slog.Logger
	storage Storage
}

func NewProduct(ctx context.Context, log *slog.Logger, storage Storage) *Product {
	return &Product{
		ctx:     ctx,
		logger:  log,
		storage: storage,
	}
}

func (p *Product) GetProducts(ctx context.Context) ([]*models.Product, error) {
	const fn = "GetProducts"

	products, err := p.storage.GetProducts(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	return products, err
}

func (p *Product) GetProduct(ctx context.Context, id int) (*models.Product, error) {
	const fn = "GetProduct"

	product, err := p.storage.GetProduct(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("%s : %w", fn, err)
	}

	return product, err
}

func (p *Product) SaveProduct(ctx context.Context, product *models.Product) error {
	const fn = "SaveProduct"

	err := p.storage.SaveProduct(ctx, product)
	if err != nil {
		return fmt.Errorf("%s : %w", fn, err)
	}

	return err
}
