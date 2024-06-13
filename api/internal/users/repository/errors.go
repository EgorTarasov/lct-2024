package repository

import (
	"errors"
)

var (
	// ErrVkUserNotFound is the error when a VK user account is not found.
	ErrVkUserNotFound = errors.New("vk user account not found")

	// ErrUserNotFound is the error when a user account is not found.
	ErrUserNotFound = errors.New("can't find user")

	// ErrUserNotCreated is the error when a user record cannot be created.
	ErrUserNotCreated = errors.New("can't create user record")

	// ErrEmailNotCreated is the error when an email record cannot be created.
	ErrEmailNotCreated = errors.New("can't create email record")

	// ErrUserPasswordNotFound is the error when the user's password hash cannot be retrieved.
	ErrUserPasswordNotFound = errors.New("unable to get users password hash")

	// ErrEmailUpdate is the error when an email usage cannot be updated.
	ErrEmailUpdate = errors.New("unable to update email usage")

	// ErrVkUserSave is the error when VK user data cannot be saved.
	ErrVkUserSave = errors.New("pg err SaveVkUserData")

	// ErrVkUserUpdate is the error when VK user data cannot be updated.
	ErrVkUserUpdate = errors.New("pg err UpdateVkUserData")
)
