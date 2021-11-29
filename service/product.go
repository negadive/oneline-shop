package service

import (
	"context"

	"github.com/negadive/oneline/custom_errors"

	"github.com/negadive/oneline/authorizer"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"gorm.io/gorm"
)

type IProductService interface {
	Store(ctx context.Context, actor_id *uint, product *model.Product) error
	GetOne(ctx context.Context, product_id *uint) (*model.Product, error)
	FindAll(ctx context.Context) (*[]model.Product, error)
	FindAllByUser(ctx context.Context, owner_id *uint) (*[]model.Product, error)
	Delete(ctx context.Context, actor_id *uint, product_id *uint) error
	Update(ctx context.Context, actor_id *uint, product_id *uint, product *model.Product) error
}

type ProductService struct {
	DBCon          *gorm.DB
	ProductAuthzer authorizer.IProductAuthorizer
	ProductRepo    repository.IProductRepository
}

func NewProductService(db_con *gorm.DB, productAuthzer authorizer.IProductAuthorizer, product_repo repository.IProductRepository) IProductService {
	return &ProductService{
		DBCon:          db_con,
		ProductAuthzer: productAuthzer,
		ProductRepo:    product_repo,
	}
}

func (s *ProductService) Store(ctx context.Context, actor_id *uint, product *model.Product) error {
	tx := s.DBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	if !s.ProductAuthzer.CanCreate(product, actor_id) {
		return custom_errors.NewForbiddenUser("not product owner")
	}
	if err := s.ProductRepo.Store(ctx, tx, product); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *ProductService) GetOne(ctx context.Context, product_id *uint) (*model.Product, error) {
	tx := s.DBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	product, err := s.ProductRepo.FindById(ctx, tx, product_id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return product, nil
}

func (s *ProductService) FindAll(ctx context.Context) (*[]model.Product, error) {
	tx := s.DBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	products, err := s.ProductRepo.FindAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return products, nil
}

func (s *ProductService) FindAllByUser(ctx context.Context, owner_id *uint) (*[]model.Product, error) {
	tx := s.DBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	products, err := s.ProductRepo.FindAllOwnerByUser(ctx, tx, owner_id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return products, nil
}

func (s *ProductService) Update(ctx context.Context, actor_id *uint, product_id *uint, product *model.Product) error {
	tx := s.DBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	old_product, err := s.ProductRepo.FindById(ctx, tx, product_id)
	if err != nil {
		tx.Rollback()
		return custom_errors.NewNotFoundError("product")
	}
	if !s.ProductAuthzer.CanEdit(old_product, actor_id) {
		tx.Rollback()
		return custom_errors.NewForbiddenUser("not product owner")
	}

	if err := s.ProductRepo.Update(ctx, tx, product_id, product); err != nil {
		tx.Rollback()
		return err
	}
	new_product, err := s.ProductRepo.FindById(ctx, tx, product_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	*product = *new_product

	return nil
}

func (s *ProductService) Delete(ctx context.Context, actor_id *uint, product_id *uint) error {
	tx := s.DBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	old_product, err := s.ProductRepo.FindById(ctx, tx, product_id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if !s.ProductAuthzer.CanDelete(old_product, actor_id) {
		tx.Rollback()
		return custom_errors.NewForbiddenUser("not product owner")
	}

	if err := s.ProductRepo.Delete(ctx, tx, product_id); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
