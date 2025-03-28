package service

import (
	"gorm.io/gorm"
)

type ShipmentService struct {
	db *gorm.DB
}

func (s *ShipmentService) getAllShipments() {
	
}