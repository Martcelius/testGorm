package contact

type Contact struct {
	ID          uint `json:"customerId" gorm:"primary_key"`
	CountryCode int  `json:"countryCode"`
	MobileNo    uint `json:"mobileNo"`
	CustID      uint
}
