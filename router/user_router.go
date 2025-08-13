package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router {
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Get("/user/{id}", ur.userController.GetUserById)
	r.Post("/signup", ur.userController.CreateUser)
	r.Post("/login", ur.userController.LoginUser)
	r.Get("/users", ur.userController.GetAllUsers)
	r.Delete("/users/{id}", ur.userController.DeleteUserById)
}
