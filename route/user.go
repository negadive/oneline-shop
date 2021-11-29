package route

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/negadive/oneline/authorizer"
	"github.com/negadive/oneline/handler"
	"github.com/negadive/oneline/repository"
	"github.com/negadive/oneline/service"
	"gorm.io/gorm"
)

func setupUserHandler(db *gorm.DB, validate *validator.Validate) handler.IUserHandler {
	UserRepo := repository.NewUserRepository()
	UserAuthzer := authorizer.NewUserAuthorizer(UserRepo)
	UserService := service.NewUserService(db, UserAuthzer, UserRepo)
	UserHandler := handler.NewUserHandler(
		UserService,
		validate,
	)

	return UserHandler
}

func User(app *fiber.App, UserHandler handler.IUserHandler) {
	user := app.Group("/users")

	user.Post("/", UserHandler.Register)
	user.Patch("/:id", UserHandler.Update)

}
