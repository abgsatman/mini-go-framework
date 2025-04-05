package userapi

import (
	"context"
	"errors"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*User, error)
	GetUserById(ctx context.Context, id int64) (*User, error)
	CreateUser(ctx context.Context, user *User) error
}

type InMemoryUserRepository struct {
	users map[int64]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int64]*User),
	}
}

func (r *InMemoryUserRepository) GetAllUsers(ctx context.Context) ([]*User, error) {
	var users []*User
	return users, nil
}

func (r *InMemoryUserRepository) GetUserById(ctx context.Context, id int64) (*User, error) {
	user, ok := r.users[int64(id)]
	if !ok {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (r *InMemoryUserRepository) CreateUser(ctx context.Context, user *User) error {
	r.users[int64(user.ID)] = user

	return nil
}
