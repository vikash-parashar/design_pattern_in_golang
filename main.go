package main

import (
	"database/sql"
)

type UserRepository interface {
	Create(user User) error
	Read(id int) (User, error)
	Update(user User) error
	Delete(id int) error
}
type SQLUserRepository struct {
	db *sql.DB
}
type User struct {
	ID       int
	Name     string
	Email    string
	Password string // Note: Password should be securely hashed and not stored in plain text
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db}
}

func (r *SQLUserRepository) Create(user User) error {
	// Implement database insertion logic
	_, err := r.db.Exec("INSERT INTO users (name, email) VALUES (?, ?)", user.Name, user.Email)
	return err
}

func (r *SQLUserRepository) Read(id int) (User, error) {
	// Implement database retrieval logic
	var user User
	err := r.db.QueryRow("SELECT name, email FROM users WHERE id = ?", id).Scan(&user.Name, &user.Email)
	if err != nil {
		return User{}, err
	}
	user.ID = id
	return user, nil
}

func (r *SQLUserRepository) Update(user User) error {
	// Implement database update logic
	_, err := r.db.Exec("UPDATE users SET name = ?, email = ? WHERE id = ?", user.Name, user.Email, user.ID)
	return err
}

func (r *SQLUserRepository) Delete(id int) error {
	// Implement database deletion logic
	_, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
func CreateUser(repo UserRepository, user User) error {
	return repo.Create(user)
}

func GetUser(repo UserRepository, id int) (User, error) {
	return repo.Read(id)
}

func UpdateUser(repo UserRepository, user User) error {
	return repo.Update(user)
}

func DeleteUser(repo UserRepository, id int) error {
	return repo.Delete(id)
}
