package service

import (
	bo "backend/bo/admin_user"
	"backend/repository"
	"context"
)

type AdminService struct {
	adminRepository *repository.AdminUserRepository
}

func NewAdminService(repositories *repository.Repositories) *AdminService {
	return &AdminService{
		adminRepository: repositories.AdminUserRepository,
	}
}

func (a *AdminService) GetAdminUserInfoByID(ctx context.Context, adminUserID int64) (*bo.AdminUser, error) {
	adminUserDO, err := a.adminRepository.GetAdminUserInfoByID(ctx, adminUserID)
	if err != nil {
		return nil, err
	}
	bo := &bo.AdminUser{
		Name: adminUserDO.Name,
	}
	return bo, nil
}
