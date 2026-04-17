package service

import "backend/repository"

type Services struct {
	UserService  *UserService
	AdminService *AdminService
}

// 构建所有services
func NewServices(repositories *repository.Repositories) *Services {
	return &Services{
		UserService:  NewUserService(repositories),
		AdminService: NewAdminService(repositories),
	}
}
