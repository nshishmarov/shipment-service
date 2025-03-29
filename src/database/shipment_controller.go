package database

import (
	"net/http"
	"service/src/api/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShipmentController struct {
	DB *gorm.DB
}

func (sr *ShipmentController) GetAllShipments(ctx *gin.Context) {
	var shipments = []dto.Shipment{
		{ShipmentName: "test", ShipmentDate: "test"},
	}
	ctx.IndentedJSON(http.StatusOK, shipments)
}