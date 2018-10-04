package models

// User : struct
type User struct {
	BaseModel
	UserID  string `gorm:"column:userID;unique;default:NULL" json:"userID"`
	Payload string `gorm:"column:payload" json:"payload"`
}

// TableName : user
func (u User) TableName() string {
	return "user"
}
