package models

type Person struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
