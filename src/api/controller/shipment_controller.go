package controller

import (
	"service/src/database"

	routing "github.com/gin-gonic/gin"
)

type ShipmentController struct {
	shipmentRepo *database.ShipmentRepository
}

func (c *ShipmentController) ShipmentRoute(rg *routing.RouterGroup)  {
	router := rg.Group("shipments")
	router.GET("/", c.shipmentRepo.GetAllShipments)
}
