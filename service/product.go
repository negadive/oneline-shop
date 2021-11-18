package service

import (
	"context"

	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"gorm.io/gorm"
)

type IProductService interface {
	GetProductRepo() repository.IProductRepository
	Store(ctx context.Context, product *model.Product) error
	GetOne(ctx context.Context, product_id *uint) (*model.Product, error)
	FindAll(ctx context.Context) (*[]model.Product, error)
	FindAllByUser(ctx context.Context, owner_id *uint) (*[]model.Product, error)
	Delete(ctx context.Context, product_id *uint) error
	Update(ctx context.Context, product_id *uint, product *model.Product) error
}

type ProductService struct {
	ProductRepo repository.IProductRepository
}

func NewProductService(product_repo repository.IProductRepository) IProductService {
	return &ProductService{
		ProductRepo: product_repo,
	}
}

func (s *ProductService) GetProductRepo() repository.IProductRepository {
	return s.ProductRepo
}

func (s *ProductService) Store(ctx context.Context, product *model.Product) error {
	if err := s.ProductRepo.Store(ctx, product); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) GetOne(ctx context.Context, product_id *uint) (*model.Product, error) {
	product, err := s.ProductRepo.FindById(ctx, product_id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) FindAll(ctx context.Context) (*[]model.Product, error) {
	products, err := s.ProductRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) FindAllByUser(ctx context.Context, owner_id *uint) (*[]model.Product, error) {
	products, err := s.ProductRepo.FindAllOwnerByUser(ctx, owner_id)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (s *ProductService) Update(ctx context.Context, product_id *uint, product *model.Product) error {
	if !s.ProductRepo.IsExists(ctx, product_id) {
		return gorm.ErrRecordNotFound
	}
	if err := s.ProductRepo.Update(ctx, product_id, product); err != nil {
		return err
	}
	new_product, err := s.ProductRepo.FindById(ctx, product_id)
	if err != nil {
		return err
	}
	*product = *new_product

	return nil
}

func (s *ProductService) Delete(ctx context.Context, product_id *uint) error {
	if !s.ProductRepo.IsExists(ctx, product_id) {
		return gorm.ErrRecordNotFound
	}
	if err := s.ProductRepo.Delete(ctx, product_id); err != nil {
		return err
	}

	return nil
}
