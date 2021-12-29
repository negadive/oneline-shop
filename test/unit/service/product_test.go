package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/negadive/oneline/authorizer"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"github.com/negadive/oneline/service"
)

func setupProductServiceDeps() (repository.IProductRepository, authorizer.IProductAuthorizer) {
	productRepo := repository.NewProductRepository()
	productAuthzer := authorizer.NewProductAuthorizer(productRepo)

	return productRepo, productAuthzer
}

func TestProductService_Store(t *testing.T) {
	db := setupTestDB()
	truncate_product(db)
	truncate_user(db)
	productRepo, productAuthzer := setupProductServiceDeps()
	productService := service.NewProductService(db, productAuthzer, productRepo)

	ctx := context.Background()

	user1 := &model.User{Email: "test", Name: "test", Password: "test"}
	user2 := &model.User{Email: "test2", Name: "test2", Password: "test2"}
	db.Create(user1)
	db.Create(user2)
	products := []*model.Product{
		{Name: "product", Price: 1000, OwnerID: user1.ID},
		{Name: "product2", Price: 1000, OwnerID: user1.ID},
	}

	type args struct {
		ctx     context.Context
		actorId *uint
		product *model.Product
	}
	tests := []struct {
		name    string
		s       service.IProductService
		args    args
		wantErr bool
	}{
		{
			name: "Success store",
			s:    productService,
			args: args{
				ctx:     ctx,
				actorId: &user1.ID,
				product: products[0],
			},
			wantErr: false,
		},
		{
			name: "Error unique product store",
			s:    productService,
			args: args{
				ctx:     ctx,
				actorId: &user1.ID,
				product: products[0],
			},
			wantErr: true,
		},
		{
			name: "Error store by other user",
			s:    productService,
			args: args{
				ctx:     ctx,
				actorId: &user2.ID,
				product: products[1],
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Store(tt.args.ctx, tt.args.actorId, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("ProductService.Store() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	truncate_product(db)
	truncate_user(db)
}

func TestProductService_GetOne(t *testing.T) {
	db := setupTestDB()
	truncate_product(db)
	truncate_user(db)
	productRepo, productAuthzer := setupProductServiceDeps()
	productService := service.NewProductService(db, productAuthzer, productRepo)

	user1 := &model.User{Email: "test", Name: "test", Password: "test"}
	user2 := &model.User{Email: "test2", Name: "test2", Password: "test2"}
	db.Create(user1)
	db.Create(user2)
	products := []*model.Product{
		{Name: "product", Price: 1000, OwnerID: user1.ID},
		{Name: "product2", Price: 1000, OwnerID: user1.ID},
	}
	db.CreateInBatches(products, 5)
	random_uint := uint(909033)

	ctx := context.Background()
	type args struct {
		ctx       context.Context
		productId *uint
	}
	tests := []struct {
		name    string
		s       service.IProductService
		args    args
		want    *model.Product
		wantErr bool
	}{
		{
			name: "Success get one",
			s:    productService,
			args: args{
				ctx:       ctx,
				productId: &products[0].ID,
			},
			want:    products[0],
			wantErr: false,
		},
		{
			name: "Success get another one",
			s:    productService,
			args: args{
				ctx:       ctx,
				productId: &products[1].ID,
			},
			want:    products[1],
			wantErr: false,
		},
		{
			name: "Error random id",
			s:    productService,
			args: args{
				ctx:       ctx,
				productId: &random_uint,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetOne(tt.args.ctx, tt.args.productId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.GetOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("ProductService.GetOne() = %v, want %v", got, tt.want)
				return
			}
			if tt.want != nil && got != nil {
				if got.ID != tt.want.ID {
					t.Errorf("ProductService.GetOne() = %v, want %v", got.ID, tt.want.ID)
				}
			}
		})
	}
	truncate_product(db)
	truncate_user(db)
}

func TestProductService_FindAll(t *testing.T) {
	db := setupTestDB()
	truncate_product(db)
	truncate_user(db)
	productRepo, productAuthzer := setupProductServiceDeps()
	productService := service.NewProductService(db, productAuthzer, productRepo)

	user1 := &model.User{Email: "test", Name: "test", Password: "test"}
	user2 := &model.User{Email: "test2", Name: "test2", Password: "test2"}
	db.Create(user1)
	db.Create(user2)
	products := []*model.Product{
		{Name: "product", Price: 1000, OwnerID: user1.ID},
		{Name: "product2", Price: 1000, OwnerID: user1.ID},
	}
	db.CreateInBatches(products, 5)

	ctx := context.Background()
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		s       service.IProductService
		args    args
		want    []*model.Product
		wantErr bool
	}{
		{
			name:    "Find all",
			s:       productService,
			args:    args{ctx: ctx},
			want:    products,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FindAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.FindAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want == nil && got != nil {
				t.Errorf("ProductService.GetOne() = %v, want %v", got, tt.want)
				return
			}
			if tt.want != nil && got != nil {
				for idx, product := range *got {
					if product.ID != tt.want[idx].ID {
						t.Errorf("ProductService.GetOne() = %v, want %v", product.ID, tt.want[idx].ID)
					}
				}
			}
		})
	}
}

func TestProductService_FindAllByUser(t *testing.T) {
	type args struct {
		ctx      context.Context
		owner_id *uint
	}
	tests := []struct {
		name    string
		s       service.IProductService
		args    args
		want    *[]model.Product
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.FindAllByUser(tt.args.ctx, tt.args.owner_id)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProductService.FindAllByUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductService.FindAllByUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProductService_Update(t *testing.T) {
	type args struct {
		ctx       context.Context
		actorId   *uint
		productId *uint
		product   *model.Product
	}
	tests := []struct {
		name    string
		s       service.IProductService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Update(tt.args.ctx, tt.args.actorId, tt.args.productId, tt.args.product); (err != nil) != tt.wantErr {
				t.Errorf("ProductService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestProductService_Delete(t *testing.T) {
	type args struct {
		ctx       context.Context
		actorId   *uint
		productId *uint
	}
	tests := []struct {
		name    string
		s       service.IProductService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Delete(tt.args.ctx, tt.args.actorId, tt.args.productId); (err != nil) != tt.wantErr {
				t.Errorf("ProductService.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
