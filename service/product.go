package service

import (
	"context"

	"github.com/negadive/oneline/customErrors"

	"github.com/negadive/oneline/authorizer"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"gorm.io/gorm"
)

type IProductService interface {
	Store(ctx context.Context, actorId *uint, product *model.Product) error
	GetOne(ctx context.Context, productId *uint) (*model.Product, error)
	FindAll(ctx context.Context) (*[]*model.Product, error)
	FindAllByUser(ctx context.Context, owner_id *uint) (*[]*model.Product, error)
	Delete(ctx context.Context, actorId *uint, productId *uint) error
	Update(ctx context.Context, actorId *uint, productId *uint, product *model.Product) error
}

type ProductService struct {
	dBCon          *gorm.DB
	productAuthzer authorizer.IProductAuthorizer
	productRepo    repository.IProductRepository
}

func NewProductService(
	dbCon *gorm.DB,
	productAuthzer authorizer.IProductAuthorizer,
	productRepo repository.IProductRepository,
) IProductService {
	return &ProductService{
		dBCon:          dbCon,
		productAuthzer: productAuthzer,
		productRepo:    productRepo,
	}
}

func (s *ProductService) Store(ctx context.Context, actorId *uint, product *model.Product) error {
	tx := s.dBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	if !s.productAuthzer.CanCreate(product, actorId) {
		return customErrors.NewForbiddenUser("not product owner")
	}
	if err := s.productRepo.Store(ctx, tx, product); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (s *ProductService) GetOne(ctx context.Context, productId *uint) (*model.Product, error) {
	tx := s.dBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	product, err := s.productRepo.FindById(ctx, tx, productId)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return product, nil
}

func (s *ProductService) FindAll(ctx context.Context) (*[]*model.Product, error) {
	tx := s.dBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	products, err := s.productRepo.FindAll(ctx, tx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return products, nil
}

func (s *ProductService) FindAllByUser(ctx context.Context, owner_id *uint) (*[]*model.Product, error) {
	tx := s.dBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	products, err := s.productRepo.FindAllOwnerByUser(ctx, tx, owner_id)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	return products, nil
}

func (s *ProductService) Update(ctx context.Context, actorId *uint, productId *uint, product *model.Product) error {
	tx := s.dBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	oldProduct, err := s.productRepo.FindById(ctx, tx, productId)
	if err != nil {
		tx.Rollback()
		return customErrors.NewNotFoundError("product")
	}
	if !s.productAuthzer.CanEdit(oldProduct, actorId) {
		tx.Rollback()
		return customErrors.NewForbiddenUser("not product owner")
	}

	if err := s.productRepo.Update(ctx, tx, productId, product); err != nil {
		tx.Rollback()
		return err
	}
	new_product, err := s.productRepo.FindById(ctx, tx, productId)
	if err != nil {
		tx.Rollback()
		return err
	}
	*product = *new_product

	return nil
}

func (s *ProductService) Delete(ctx context.Context, actorId *uint, productId *uint) error {
	tx := s.dBCon.Session(&gorm.Session{SkipDefaultTransaction: true})
	defer tx.Commit()

	oldProduct, err := s.productRepo.FindById(ctx, tx, productId)
	if err != nil {
		tx.Rollback()
		return err
	}
	if !s.productAuthzer.CanDelete(oldProduct, actorId) {
		tx.Rollback()
		return customErrors.NewForbiddenUser("not product owner")
	}

	if err := s.productRepo.Delete(ctx, tx, productId); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
