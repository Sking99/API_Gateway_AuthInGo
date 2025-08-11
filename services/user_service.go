package services

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

type UserService interface {
	GetUserById() error
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
	u.userRepository.GetAll()
	return nil
}
