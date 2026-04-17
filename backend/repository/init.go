package repository

import "gorm.io/gorm"

type Repositories struct {
	UserRepository      *UserRepository
	AdminUserRepository *AdminUserRepository
}

// 统一构造所有repository
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository:      NewUserRepository(db),
		AdminUserRepository: NewAdminUserRepository(db),
	}
}
