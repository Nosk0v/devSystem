package models

type User struct {
	Username string `db:"username"`
}

type CreateOrganizationInput struct {
	Name   string `json:"name" binding:"required"`
	Prefix string `json:"prefix" binding:"required"`
}
