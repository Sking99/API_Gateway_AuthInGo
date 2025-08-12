package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById() (*models.User, error)
	Create(username string, email string, hashedPassword string) error
	GetAll() ([]*models.User, error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {

	query := "SELECT * FROM users;"

	rows, err := u.db.Query(query)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			return nil, err
		} else {
			fmt.Println("Error scanning user")
			return nil, err
		}
	}

	var users []models.User

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at)
		if err != nil {
			return nil, err
		}

		users = append(users, *user)

	}

	fmt.Println("Users data:", users)

	return nil, nil
}

func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) error {
	query := "INSERT into users (username, email, password) values (?, ?, ?);"

	result, err := u.db.Exec(query, username, email, hashedPassword)

	if err != nil {
		fmt.Println("Error inserting the row")
		return err
	}

	rowAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error affecting the row")
		return rowErr
	}

	if rowAffected == 0 {
		fmt.Println("User not created")
		return nil
	}

	fmt.Println("User craeted successfully, Row affected:", rowAffected)

	return nil
}

func (u *UserRepositoryImpl) GetById() (*models.User, error) {
	fmt.Println("User repository called")

	query := "SELECT id, username, password, email, created_at, updated_at FROM users WHERE id = ?"

	row := u.db.QueryRow(query, 1)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found")
			return nil, err
		} else {
			fmt.Println("Error scanning user")
			return nil, err
		}
	}

	fmt.Println("User fetched successfully", user)

	return user, nil
}
