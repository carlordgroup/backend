package data

import (
	"carlord/ent"
	"context"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	*ent.Account
	ctx context.Context
}

func NewAccount(ctx context.Context, e *ent.Account) *Account {
	return &Account{Account: e, ctx: ctx}
}

// Verify the user information with raw password and bcrypt
func (a *Account) Verify(rawPassword string) bool {
	bcryptPass, err := base64.StdEncoding.DecodeString(a.Password)
	if err != nil {
		return false
	}
	err = bcrypt.CompareHashAndPassword(bcryptPass, []byte(rawPassword))
	return err == nil
}

func (a *Account) Admin() bool {
	return a.IsAdmin
}

func (a *Account) User() (*User, error) {
	user, err := a.QueryUser().Only(a.ctx)
	return NewUser(a.ctx, user), err
}
