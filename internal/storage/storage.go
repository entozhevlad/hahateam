package storage

import "errors"

var (
	ErrWithLogin    = errors.New("trouble with login")
	ErrWithPassword = errors.New("trouble with password")
	ErrWithCompany  = errors.New("trouble with company")
)
