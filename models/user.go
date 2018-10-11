package models

// User : struct
type User struct {
	BaseModel
	Payload       string `gorm:"column:payload" json:"payload"`
	WalletAddress string `gorm:"column:wallet_address" json:"wallet_address"`
}

// TableName : user
func (u User) TableName() string {
	return "user"
}
