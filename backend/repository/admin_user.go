package repository

import (
	"backend/model"
	"backend/query"
	"context"

	"gorm.io/gorm"
)

type AdminUserRepository struct {
	db *gorm.DB
}

func NewAdminUserRepository(db *gorm.DB) *AdminUserRepository {
	return &AdminUserRepository{
		db: db,
	}
}

func (u *AdminUserRepository) GetAdminUserInfoByID(ctx context.Context, adminUserID int64) (*model.AdminUser, error) {
	query := query.Use(u.db).AdminUser
	return query.WithContext(ctx).Where(query.ID.Eq(adminUserID)).First()
}
