package models

import "time"

type Metadata struct {
	ID        uint       `gorm:"primary_key" json:"_id"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index" json:"-"`
}
