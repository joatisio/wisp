package models

type Person struct {
	BaseModel
	FirstName string
	LastName  string

	Address string

	Avatar Media

	OwnerID   ID
	OwnerType string
}
