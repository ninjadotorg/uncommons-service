package models

// User : struct
type User struct {
	BaseModel
	RefID         string `gorm:"column:ref_id;unique;default:NULL" json:"ref_id"`
	Payload       string `gorm:"column:payload" json:"payload"`
	WalletAddress string `gorm:"column:wallet_address" json:"wallet_address"`
}

// TableName : user
func (u User) TableName() string {
	return "user"
}
