package users

import (
	"database/sql"
	"fmt"

	"github.com/gokuls-codes/go-echo-starter/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}


func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if (err != nil) {
		return nil, err
	}
	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if (err != nil) {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	u := new(types.User)
	err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *Store) CreateUser(user *types.User) error {
	_, err := s.db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", user.Name, user.Email, user.Password)
	return err
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
	if (err != nil) {
		return nil, err
	}
	user := new(types.User)
	for rows.Next() {
		user, err = scanRowIntoUser(rows)
		if (err != nil) {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
	
}