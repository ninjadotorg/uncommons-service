package models

// User : struct
type User struct {
	BaseModel
	UserID    string `gorm:"column:userID;unique;default:NULL" json:"userID"`
	Payload   string `gorm:"column:payload" json:"payload"`
	Introduce string `gorm:"column:introduce" json:"introduce"`
	Name      string `gorm:"column:name" json:"name"`
	Built     string `gorm:"column:built" json:"built"`
	Github    string `gorm:"column:github" json:"github"`
	Twitter   string `gorm:"column:twitter" json:"twitter"`
	Address   string `gorm:"column:address" json:"address"`
}

// TableName : user
func (u User) TableName() string {
	return "user"
}
