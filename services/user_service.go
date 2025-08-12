package services

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() error
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

func (u *UserServiceImpl) LoginUser() error {
	fmt.Println("Login user service called")
	// username := "testuser2"
	password := "hashedPassword123@example"

	response := utils.CheckPasswordHash(password, "$2a$10$VqkLETn926U/vS2gHAEg9ODHaA/DZKFpyfBH2VctzMB8bH0wpLwom")

	fmt.Println("Password match:", response)

	return nil
}
