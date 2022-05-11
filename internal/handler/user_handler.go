package handler

import (
	"encoding/json"
	"fmt"
	"go-standard/internal/model"
	"net/http"
	"strconv"
)

type UserService interface {
	GetUsers(page int64, limit int64) ([]*model.User, error)
}

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	currentPage, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		fmt.Println(" page must be number: ", err.Error())
		return
	}
	fmt.Println("currentPage: ", currentPage)

	perPage := r.URL.Query().Get("perPage")
	limit, err := strconv.ParseInt(perPage, 10, 64)
	if err != nil {
		fmt.Println(" perPage must be number: ", err.Error())
		return
	}
	fmt.Println(limit)
	users, err := h.userService.GetUsers(currentPage, limit)
	for _, user := range users {
		fmt.Println("user: ", user)
		json.NewEncoder(w).Encode(user)
	}
}
