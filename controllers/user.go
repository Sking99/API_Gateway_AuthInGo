package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController {
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetch user called in user controller")
	uc.UserService.GetUserById()
	w.Write([]byte("User profile fetch endpoint"))
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Register user called in user controller")
	uc.UserService.CreateUser()
	w.Write([]byte("User registration endpoint"))
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	var payload dto.UserLoginDTO

	if jsonErr := utils.ReadJsonRequest(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request payload", jsonErr)
		return
	}

	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error", validationErr)
		return
	}

	jwtToken, err := uc.UserService.LoginUser(&payload)

	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error logging in user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User logged in successfully", jwtToken)
}

func (uc *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all users called in user controller")
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}
	fmt.Println(w, "Fetched users: %+v", users)
}

func (uc *UserController) DeleteUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete user by ID called in user controller")
	id := 1 // This should be extracted from the request
	err := uc.UserService.DeleteUserById(id)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User deleted successfully"))
}
