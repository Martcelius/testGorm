package customer

import (
	"testGorm/contact"
)

type Customer struct {
	ID       uint              `json:"customerId" gorm:"primary_key"`
	Name     string            `json:"name"`
	Contacts []contact.Contact `json:"contacts" gorm:"foreignKey:CustId"`
}
