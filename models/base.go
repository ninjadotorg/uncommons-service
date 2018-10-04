package models

import (
	"time"
)

// BaseModel : struct
type BaseModel struct {
	ID           uint      `gorm:"primary_key" json:"id"`
	DateCreated  time.Time `gorm:"column:date_created" json:"date_created"`
	DateModified time.Time `gorm:"column:date_modified" json:"date_modified"`
}

// BeforeCreate : create time
func (m *BaseModel) BeforeCreate() (err error) {
	m.DateCreated = time.Now()
	m.DateModified = time.Now()
	return
}

// BeforeUpdate : update time
func (m *BaseModel) BeforeUpdate() (err error) {
	m.DateModified = time.Now()
	return
}
