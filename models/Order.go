package models

import "time"

type Order struct {
	OrderID      uint      `gorm:"primaryKey" json:"orderId"`
	CustomerName string    `gorm:"type:varchar(191);not null;" json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `json:"items"`
}
