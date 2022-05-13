package main

import (
	"fmt"
	"github.com/gorilla/mux"
	mysqlDB "go-standard/internal/db"
	"go-standard/internal/handler"
	"go-standard/internal/repo"
	"go-standard/internal/service"
	"log"
	"net/http"
)

func main() {
	sqlDns := "root:nhatminh21@tcp(165.22.245.167:13306)/backend"

	db, err := mysqlDB.NewDB(sqlDns)
	if err != nil {
		fmt.Println("can not connect to database:", err.Error())
	}

	userRepo := repo.NewUser(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := mux.NewRouter()
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", userHandler.DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":5000", r))
}
