package models

import "time"

type Post struct {
	Id int `gorm:"primaryKey:not null:autoIncrement" json:"id"`
	Name string `json:"name"`
	Time float64 `gorm:"not null" json:"time"`
	ImageUrl string `gorm:"not null" json:"image_url"`
	CreatedAt time.Time `gorm:"autoCreateTime:int" json:"created_at"`	
	UpdatedAt time.Time `gorm:"autoUpdateTime:int" json:"updated_at"`
}
