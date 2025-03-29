package database

import (
	"net/http"
	"service/src/api/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShipmentRepository struct {
	DB *gorm.DB
}

func (sr *ShipmentRepository) GetAllShipments(ctx *gin.Context) {
	var shipments = []dto.Shipment{
		{ShipmentName: "test", ShipmentDate: "test"},
	}
	ctx.IndentedJSON(http.StatusOK, shipments)
}