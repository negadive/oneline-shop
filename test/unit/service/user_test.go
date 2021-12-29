package service

import (
	"context"
	"testing"

	"github.com/negadive/oneline/authorizer"
	"github.com/negadive/oneline/model"
	"github.com/negadive/oneline/repository"
	"github.com/negadive/oneline/service"
)

func setupUserServiceDeps() (repository.IUserRepository, authorizer.IUserAuthorizer) {
	userRepo := repository.NewUserRepository()
	userAuthzer := authorizer.NewUserAuthorizer(userRepo)

	return userRepo, userAuthzer
}

func TestUserService_Register(t *testing.T) {
	db := setupTestDB()
	truncate_user(db)
	userRepo, userAuthzer := setupUserServiceDeps()
	userService := service.NewUserService(db,
		userAuthzer,
		userRepo,
	)

	type args struct {
		ctx  context.Context
		user *model.User
	}
	tests := []struct {
		name    string
		s       service.IUserService
		args    args
		wantErr bool
	}{
		{
			name: "Error on invalid data",
			s:    userService,
			args: args{
				ctx: nil,
				user: &model.User{
					Email:    "hehe@gmail.com",
					Name:     "",
					Password: "",
				},
			},
			wantErr: true,
		},
		{
			name: "Success",
			s:    userService,
			args: args{
				ctx: nil,
				user: &model.User{
					Email:    "test@gmail.com",
					Name:     "test",
					Password: "test",
				},
			},
			wantErr: false,
		},
		{
			name: "Error duplicate",
			s:    userService,
			args: args{
				ctx: nil,
				user: &model.User{
					Email:    "test@gmail.com",
					Name:     "test",
					Password: "test",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Register(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.Register() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	truncate_user(db)
}

func TestUserService_Update(t *testing.T) {
	db := setupTestDB()
	truncate_user(db)
	userRepo, userAuthzer := setupUserServiceDeps()
	userService := service.NewUserService(db,
		userAuthzer,
		userRepo,
	)

	ctx := context.Background()
	user1 := &model.User{
		Email:    "test@gmail.com",
		Name:     "test",
		Password: "test",
	}
	userService.Register(ctx, user1)
	user2 := &model.User{
		Email:    "test2@gmail.com",
		Name:     "test2",
		Password: "test2",
	}
	userService.Register(ctx, user2)

	type args struct {
		ctx     context.Context
		actorId *uint
		user_id *uint
		user    *model.User
	}
	tests := []struct {
		name    string
		s       service.IUserService
		args    args
		wantErr bool
	}{
		{
			name: "Success Update",
			s:    userService,
			args: args{
				ctx:     ctx,
				actorId: &user1.ID,
				user_id: &user1.ID,
				user: &model.User{
					Name: "test update",
				},
			},
			wantErr: false,
		},
		{
			name: "Error updated by other user",
			s:    userService,
			args: args{
				ctx:     ctx,
				actorId: &user2.ID,
				user_id: &user1.ID,
				user: &model.User{
					Name: "test update",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.Update(tt.args.ctx, tt.args.actorId, tt.args.user_id, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.Update() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	truncate_user(db)
}
