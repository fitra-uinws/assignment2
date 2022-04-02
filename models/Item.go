package models

type Item struct {
	ItemID      uint   `gorm:"primaryKey" json:"lineItemId"`
	ItemCode    string `gorm:"type:varchar(191);not null;" json:"itemCode"`
	Description string `gorm:"type:varchar(1024);not null;" json:"description"`
	Quantity    uint   `gorm:"type:bigint;not null;" json:"quantity"`
	OrderID     uint   `json:"orderId"`
}
