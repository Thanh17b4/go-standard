package service

import "go-standard/internal/model"

type UserRepo interface {
	List(page int64, limit int64) ([]*model.User, error)
	Create(user model.User) (int64, error)
	Delete(userID int64) (int64, error)
	Update(userId int64, user model.User) (*model.User, error)
}

type UserService struct {
	userRepo UserRepo
}

func NewUserService(userRepo UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s UserService) GetUsers(page int64, limit int64) ([]*model.User, error) {
	// @TODO anything
	return s.userRepo.List(page, limit)
}

func (s UserService) CreateUser(u model.User) (int64, error) {
	return s.userRepo.Create(u)
}
func (s UserService) DeleteUser(userID int64) (int64, error) {
	return s.userRepo.Delete(userID)
}
func (s UserService) UpdateUser(userId int64, user model.User) (*model.User, error) {
	return s.userRepo.Update(userId, user)

}
