package repository

import "gorm.io/gorm"

type Repositories struct {
	UserRepository *UserRepository
}

// 统一构造所有repository
func NewRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		UserRepository: NewUserRepository(db),
	}
}
