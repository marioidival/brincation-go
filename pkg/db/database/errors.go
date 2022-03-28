package database

import "errors"

var (
	DbNotFound error = errors.New("not found")
)
