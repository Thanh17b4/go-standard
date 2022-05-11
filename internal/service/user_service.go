package service

import "go-standard/internal/model"

type UserRepo interface {
	List(page int64, limit int64) ([]*model.User, error)
}

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s UserService) GetUsers(page int64, limit int64) ([]*model.User, error) {
	return s.userRepo.List(page, limit)
}
