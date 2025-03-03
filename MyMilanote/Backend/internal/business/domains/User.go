package domains

import "context"

type User struct {
	ID string
	//TODO:
}

type IUserService interface {
	UpdateUser(ctx context.Context, user *User) (*User, error)
	FindUsers(ctx context.Context, user *User) ([]User, error)
}

type IUserRepository interface {
	Update(ctx context.Context, user *User) (*User, error)
	Select(ctx context.Context, user *User) ([]User, error)
}
