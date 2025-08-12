package services

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser(payload *dto.UserLoginDTO) (string, error)
	GetAllUsers() ([]*models.User, error)
	DeleteUserById(id int) error
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("User service called")
	u.userRepository.GetById()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Create user service called")
	username := "testuser2"
	email := "test2@gmail.com"
	password, err := utils.GenerateHashedPassword("hashedPassword123@example")
	if err != nil {
		return err
	}
	u.userRepository.Create(username, email, password)
	return nil
}

func (u *UserServiceImpl) LoginUser(payload *dto.UserLoginDTO) (string, error) {
	// email := "test2@gmail.com"
	// password := "hashedPassword123@example"

	user, err := u.userRepository.GetByEmail(payload.Email)
	if err != nil {
		fmt.Println("Error fetching user by email:", err)
		return "", err
	}
	if user == nil {
		fmt.Println("User not found")
		return "", fmt.Errorf("user not found with email: %s", payload.Email)
	}

	isPasswordValid := utils.CheckPasswordHash(payload.Password, user.Password)
	if !isPasswordValid {
		fmt.Println("Password does not match")
		return "", fmt.Errorf("invalid password for user: %s", payload.Email)
	}

	jwtPayload := jwt.MapClaims{
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token valid for 24 hours
		"iat":      time.Now().Unix(),                     // Issued at time
		"nbf":      time.Now().Unix(),                     // Not before time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "TOKEN")))

	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}

	// fmt.Println("Generated JWT Token:", tokenString)

	return tokenString, nil
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User, error) {
	fmt.Println("Get all users service called")
	users, err := u.userRepository.GetAll()
	if err != nil {
		fmt.Println("Error fetching users:", err)
		return nil, err
	}
	fmt.Println("Fetched users in service:", users)
	return users, nil
}

func (u *UserServiceImpl) DeleteUserById(id int) error {
	fmt.Println("Delete user by ID service called")
	err := u.userRepository.DeleteById(id)
	if err != nil {
		fmt.Println("Error deleting user by ID:", err)
		return err
	}
	fmt.Println("User deleted successfully")
	return nil
}
