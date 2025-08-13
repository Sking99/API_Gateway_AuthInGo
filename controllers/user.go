package controllers

import (
	"AuthInGo/dto"
	"AuthInGo/services"
	"AuthInGo/utils"
	"net/http"
	"strconv"
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
	userID := r.URL.Query().Get("id")
	if userID == "" {
		userID = r.Context().Value("userId").(string)
	}
	// userID, conErr := strconv.ParseInt(r.PathValue("id"), 10, 64)
	// if conErr != nil {
	// 	utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid user ID", conErr)
	// 	return
	// }

	user, err := uc.UserService.GetUserById(userID)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error fetching user by ID", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "User fetched successfully", user)
}

func (uc *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	payload := r.Context().Value("useRegisterPayload").(dto.UserRegisterDTO)

	if err := uc.UserService.CreateUser(&payload); err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusInternalServerError, "Error creating user", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusCreated, "User created successfully", nil)
}

func (uc *UserController) LoginUser(w http.ResponseWriter, r *http.Request) {

	payload := r.Context().Value("userLoginPayload").(dto.UserLoginDTO)

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
	userID, conErr := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if conErr != nil {
		utils.WriteJsonErrorResponse(w, http.StatusBadRequest, "Invalid user ID", conErr)
		return
	}
	err := uc.UserService.DeleteUserById(userID)
	if err != nil {
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}
	w.Write([]byte("User deleted successfully"))
}
