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

func (s *Store) CreateSessionForUser(session *types.Session) error {
	_, err := s.db.Exec("INSERT INTO sessions (user_id, session_token, expires_at) VALUES (?, ?, ?)", session.UserId, session.SessionToken, session.ExpiresAt)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) FindSessionBySessionId(sessionToken string) (*types.Session, error) {

	rows, err := s.db.Query("SELECT * FROM sessions WHERE session_token = ?", sessionToken)
	if (err != nil) {
		return nil, err
	}
	sess := new(types.Session)
	for rows.Next() {
		sess, err = scanRowIntoSession(rows)
		if err != nil {
			return nil, err
		}
	}
	if sess.ID == 0 {
		return nil, fmt.Errorf("session not found")
	}

	return sess, nil
}

func scanRowIntoSession(rows *sql.Rows) (*types.Session, error) {
	s := new(types.Session)
	err := rows.Scan(&s.ID, &s.UserId, &s.SessionToken, &s.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}