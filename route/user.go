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
	userRepo := repository.NewUserRepository()
	userAuthzer := authorizer.NewUserAuthorizer(userRepo)
	userService := service.NewUserService(db, userAuthzer, userRepo)
	userHandler := handler.NewUserHandler(
		userService,
		validate,
	)

	return userHandler
}

func User(app *fiber.App, userHandler handler.IUserHandler) {
	user := app.Group("/users")

	user.Post("/", userHandler.Register)
	user.Patch("/:id", userHandler.Update)

}
