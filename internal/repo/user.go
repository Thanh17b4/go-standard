package repo

import (
	"database/sql"
	"fmt"
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

func (u *User) Create(user model.User) (int64, error) {
	result, err := u.db.Exec("INSERT INTO `users` (`name`, `email`, `activated`, `banned`, `protected`) VALUES (?, ?, ?, ?, ?)", user.Name, user.Email, user.Activated, user.Banned, user.Protected)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return lastInsertID, nil
}

func (u *User) Delete(userID int64) (int64, error) {
	_, err := u.db.Exec(" DELETE FROM `users` WHERE `id`=?", userID)
	if err != nil {
		fmt.Println(" err: ", err.Error())
	}
	return userID, nil
}
