package models

import (
	"errors"
	"time"
)

var (
	ErrorNoRecord = errors.New("models: no matching record found")

	ErrInvalidCredentials = errors.New("models: invalid credentials")

	ErrDuplicateEmail = errors.New("models: duplicate email")
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type User struct {
	ID       int
	name     string
	email    string
	password []byte
	created  time.Time
	active   bool
}
