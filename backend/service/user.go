package service

type IUserService interface {
	Login(*UserLoginInput)
}

type UserService struct {
}

func NewUserService() IUserService {
	return &UserService{}
}

type UserLoginInput struct {
	Username string
	Password string
}

func (u *UserService) Login(*UserLoginInput) {

}
