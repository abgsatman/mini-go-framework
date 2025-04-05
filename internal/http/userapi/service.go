package userapi

import "context"

type UserService interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserById(ctx context.Context, userId int64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (u userService) GetAllUsers(ctx context.Context) ([]*User, error) {
	return u.repo.GetAllUsers(ctx)
}

func (u userService) GetUserById(ctx context.Context, userId int64) (*User, error) {
	return u.repo.GetUserById(ctx, userId)
}

func (u userService) CreateUser(ctx context.Context, user *User) error {
	return u.repo.CreateUser(ctx, user)
}
