package repository

import (
	"database/sql"
	"fmt"

	"github.com/alanzorzi/crud-go/app/model"
	"github.com/alanzorzi/crud-go/app/repository/interfaces"
)

type userRepository struct {
	db *sql.DB
}

var _ interfaces.UserRepositoryInterface = &userRepository{}

func NewUserRepository(db *sql.DB) interfaces.UserRepositoryInterface {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserById(id string) ([]*model.User, error) {
	rows, err := r.db.Query("SELECT * FROM user_data WHERE id = '" + id + "' LIMIT 1")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*model.User{}

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) GetAllUsers() ([]*model.User, error) {
	rows, err := r.db.Query("SELECT * FROM user_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*model.User{}

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	_, err := r.db.Exec("INSERT INTO user_data (id, name, age, email, password) VALUES (?, ?, ?, ?, ?)",
		user.ID, user.Name, user.Age, user.Email, user.Password)
	return err
}

func (r *userRepository) UpdateUser(user *model.User) error {
	_, err := r.db.Exec(fmt.Sprintf("UPDATE  user_data set name = '%s', age = %d, email = '%s', password = '%s' WHERE id  = '%s'",
		user.Name, user.Age, user.Email, user.Password, user.ID))
	return err
}

func (r *userRepository) DeleteUser(id string) error {
	rows, err := r.db.Query(fmt.Sprintf("DELETE FROM user_data WHERE id = '%s'", id))
	if err != nil {
		return err
	}
	defer rows.Close()

	err = rows.Err()

	return err
}

func (r *userRepository) GetUserByEmailAndPassword(email string, password string) ([]*model.User, error) {
	rows, err := r.db.Query(fmt.Sprintf("SELECT * FROM user_data WHERE email = '%s' AND password = '%s'", email, password))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*model.User{}

	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
