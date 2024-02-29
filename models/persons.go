package models

type Person struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by Person to `persons`
func (Person) TableName() string {
	return "persons"
}
