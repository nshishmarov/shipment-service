package controller

import (
	"service/src/database"

	routing "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShipmentController struct {
	shipmentRepo *database.ShipmentRepository
}

func NewShipmentController(DB *gorm.DB) ShipmentController {
	sr := database.ShipmentRepository{DB: DB}
	return ShipmentController{&sr}
}

func (c *ShipmentController) ShipmentRoute(rg *routing.RouterGroup)  {
	router := rg.Group("shipments")
	router.GET("/", c.shipmentRepo.GetAllShipments)
}
