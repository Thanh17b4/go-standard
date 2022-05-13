package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-standard/internal/model"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type UserService interface {
	GetUsers(page int64, limit int64) ([]*model.User, error)
	CreateUser(u model.User) (int64, error)
	DeleteUser(userID int64) (int64, error)
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

func (h UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var user model.User
	err := json.Unmarshal(reqBody, &user)
	if err != nil {
		log.Fatalln("could not Unmarshal body request")
		return
	}
	fmt.Println(user)
	insertedID, err := h.userService.CreateUser(user)
	if err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
			"error": err.Error(),
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"insertedID": insertedID,
	})
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	id, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Printf("invalid id. ID should be number")
		return
	}
	deleteId, _ := h.userService.DeleteUser(id)
	json.NewEncoder(w).Encode(map[string]interface{}{
		" Deleted userID = ": deleteId,
	})
}
