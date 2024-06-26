package models

import (
	"context"
	"errors"
	"net/mail"
	"reqship-api/helpers/auth"
	"reqship-api/helpers/db"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:users"`

	ID       int64 `bun:",pk,autoincrement:"`
	Username string
	Email    string
	Password string
}

func (u *User) SignUp() (err error) {
	db := db.Init()
	ctx := context.Background()
	defer db.Close()

	if u.Email == "" || u.Password == "" || u.Username == "" {
		return errors.New("invalid data provided")
	}

	if _, err = mail.ParseAddress(u.Email); err != nil {
		return errors.New("invalid email address")
	}

	if len(u.Username) < 5 {
		return errors.New("username must be at least 5 characters")
	}

	if len(u.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	users := []User{}
	count, err := db.NewSelect().Model(&users).Where("username = ?", u.Username).ScanAndCount(ctx)
	if err != nil {
		return
	}
	if count > 0 {
		return errors.New("account with username already exists")
	}

	users = []User{}
	count, err = db.NewSelect().Model(&users).Where("email = ?", u.Email).ScanAndCount(ctx)
	if err != nil {
		return
	}
	if count > 0 {
		return errors.New("account with email already exists")
	}

	hashed_password, err := auth.Hash(u.Password)
	if err != nil {
		return
	}

	u.Password = string(hashed_password)

	_, err = db.NewInsert().Model(u).Exec(ctx)
	return
}

type LoginUser struct {
	Username string
	Password string
}

func (u *LoginUser) Login() (err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	var users []User

	var count int

	if _, err = mail.ParseAddress(u.Username); err != nil {
		count, err = db.NewSelect().Model(&users).Where("username = ?", u.Username).ScanAndCount(ctx)

	} else {
		count, err = db.NewSelect().Model(&users).Where("email = ?", u.Username).ScanAndCount(ctx)
	}

	if err != nil {
		return
	}
	if count == 0 {
		return errors.New("cannot find user")
	}
	err = auth.CheckHash(users[0].Password, u.Password)
	return
}

func (u *User) DeleteUser() (err error) {
	db := db.Init()
	defer db.Close()
	ctx := context.Background()

	_, err = db.NewDelete().Model(u).WherePK().Exec(ctx)
	return
}

type UserLoginResponse struct {
	Token string
}
