package user

import "context"

type Service interface {
	GetUserById(ctx context.Context, id int) *User
	GetAllUsers(ctx context.Context) []*User
}
