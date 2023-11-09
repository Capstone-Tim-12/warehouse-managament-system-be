package userdb

import "context"

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (resp *User, err error)
	Create(ctx context.Context, req *UserDetail) (err error)
}
