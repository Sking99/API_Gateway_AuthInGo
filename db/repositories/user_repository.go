package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById(id string) (*models.User, error)
	Create(username string, email string, hashedPassword string) error
	GetAll() ([]*models.User, error)
	GetByEmail(email string) (*models.User, error)
	DeleteById(id int64) error
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

	var users []*models.User

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}

	fmt.Println("Users data:", users)

	return users, nil
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

func (u *UserRepositoryImpl) GetById(id string) (*models.User, error) {

	query := "SELECT id, username, password, email, created_at, updated_at FROM users WHERE id = ?"

	row := u.db.QueryRow(query, id)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Created_at, &user.Updated_at)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, username, password, email, created_at, updated_at FROM users WHERE email = ?"

	row := u.db.QueryRow(query, email)
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

	fmt.Println("User fetched successfully by email", user)

	return user, nil
}

func (u *UserRepositoryImpl) DeleteById(id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := u.db.Exec(query, id)

	if err != nil {
		fmt.Println("Error deleting user by id")
		return err
	}

	rowAffected, rowErr := result.RowsAffected()

	if rowErr != nil {
		fmt.Println("Error affecting the row")
		return rowErr
	}

	if rowAffected == 0 {
		fmt.Println("User not deleted")
		return nil
	}

	fmt.Println("User deleted successfully, Row affected:", rowAffected)

	return nil
}
