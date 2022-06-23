package models

import "time"

type Product struct {
	ID unit `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name string `jason:"name"`
	SerialNumer string `jason:"serial_number`
}