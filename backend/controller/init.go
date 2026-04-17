package controller

import "backend/service"

type Controllers struct {
	PublicController *PublicController
}

// 构建所有controller
func NewControllers(services *service.Services) *Controllers {
	return &Controllers{
		PublicController: NewPublicController(),
	}
}
