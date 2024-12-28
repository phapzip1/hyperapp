package model

// Address houses a users address information
type Post struct {
	Id string `validate:"required"`
	Title   string `validate:"required"`
	Author string `validate:"required"`
	Content  string `validate:"required"`
}