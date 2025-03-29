package controller

import (
	"service/src/database"

	routing "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShipmentRouter struct {
	shipmentController *database.ShipmentController
}

func NewShipmentController(db *gorm.DB) ShipmentRouter {
	sr := database.ShipmentController{DB: db}
	return ShipmentRouter{&sr}
}

func (c *ShipmentRouter) ShipmentRoute(rg *routing.RouterGroup)  {
	router := rg.Group("/shipments")
	router.GET("/get", c.shipmentController.GetAllShipments)
}
