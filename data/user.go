package data

import (
	"carlord/ent"
	"context"
)

type User struct {
	*ent.User
	ctx context.Context
}

func NewUser(ctx context.Context, e *ent.User) *User {
	return &User{User: e, ctx: ctx}
}
