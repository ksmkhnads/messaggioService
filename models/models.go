package models

type Message struct {
	ID        uint   `gorm:"primary_key"`
	Content   string `gorm:"not null"`
	Processed bool   `gorm:"default:false"`
}
