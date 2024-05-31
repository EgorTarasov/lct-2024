package repository

import (
	"errors"
)

var (
	ErrVkUserNotFound       = errors.New("vk user account not found")
	ErrUserNotFound         = errors.New("can't find user")
	ErrUserNotCreated       = errors.New("can't create user record")
	ErrEmailNotCreated      = errors.New("can't create email record")
	ErrUserPasswordNotFound = errors.New("unable to get users password hash")
	ErrEmailUpdate          = errors.New("unable to update email usage")
	ErrVkUserSave           = errors.New("pg err SaveVkUserData")
	ErrVkUserUpdate         = errors.New("pg err UpdateVkUserData")
)
