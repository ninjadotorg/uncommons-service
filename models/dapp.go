package models

// Dapp : struct
type Dapp struct {
	BaseModel
	DappID          uint16 `gorm:"primary_key:yes;column:dappID" json:"dappID"`
	AuthorID        string `gorm:"column:authorID" json:"authorID"`
	Name            string `gorm:"column:name" json:"name"`
	Description     string `gorm:"column:description" json:"Description"`
	WhyWeWillLove   string `gorm:"column:why_we_will_love" json:"whyWeWillLove"`
	WhyYouWantBuild string `gorm:"column:why_you_want_build" json:"whyYouWantBuild"`
	WhatCool        string `gorm:"column:what_cool" json:"whatCool"`
	HowFar          string `gorm:"column:how_far" json:"howFar"`
	ShipDate        string `gorm:"column:ship_date" json:"shipDate"`
	Amount          string `gorm:"column:amount" json:"amount"`
}

// TableName : dapp
func (u Dapp) TableName() string {
	return "dapp"
}
