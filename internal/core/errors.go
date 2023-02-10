package core

import "errors"

var (
	ErrProductNotFound = errors.New("product id not found")
	ErrProductExists   = errors.New("product already exists")
)
