package repo

import (
	"database/sql"
	"go-standard/internal/model"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{db: db}
}

func (u *User) List(page int64, limit int64) ([]*model.User, error) {
	var users []*model.User
	offset := (page - 1) * limit
	result, err := u.db.Query("SELECT id, name, address FROM users LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	for result.Next() {
		u := &model.User{}
		err := result.Scan(&u.ID, &u.Name, &u.Address)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
