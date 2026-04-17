package service

import "backend/repository"

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(repositories *repository.Repositories) *UserService {
	return &UserService{
		userRepository: repositories.UserRepository,
	}
}

func (u *UserService) GetUserInfo() {

}
