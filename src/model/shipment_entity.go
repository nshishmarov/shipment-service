package model

import "gorm.io/gorm"

type ShipmentEntity struct {
	gorm.Model
	ShipmentId uint64 `gorm:"primaryKey" `
	ShipmentName string `gorm:"type:varchar(100)"`
	ShipmentDate string `gorm:"type:varchar(100)"`
}

// Specify Table Name
func (ShipmentEntity) TableName() string {
    return "shipments"
}

