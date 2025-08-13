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
	userID := r.PathValue("id")
	user, err := uc.UserService.GetUserById(userID)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching user by ID", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var payload dto.UserRegisterDTO

	if jsonErr := utils.ReadJsonRequest(r, &payload); jsonErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid request payload", jsonErr)
		return
	}
	if validationErr := utils.Validator.Struct(payload); validationErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Validation error", validationErr)
		return
	}

	if err := uc.UserService.CreateUser(&payload); err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", nil)
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
	users, err := uc.UserService.GetAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Users fetched successfully", users)
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
