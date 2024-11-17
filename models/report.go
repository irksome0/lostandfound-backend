package models

import "time"

type Report struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"report_id"`
	UserID     uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	User       User      `gorm:"foreignKey:UserID"`
	ItemID     uint      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	Item       Item      `gorm:"foreignKey:ItemID"`
	ReportDate time.Time `json:"report_date"`
	Status     string    `json:"report_status"`
}
