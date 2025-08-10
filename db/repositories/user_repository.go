package db

import "fmt"

// import "database/sql"

type UserRepository interface {
	Create() error
}

type UserRepositoryImpl struct {
	// db *sql.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Create() error {
	fmt.Println("User repository called")
	return nil
}
