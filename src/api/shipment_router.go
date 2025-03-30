package api

import (
	"service/src/database"

	routing "github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShipmentRouter struct {
	shipmentController *database.ShipmentController
	rg *routing.RouterGroup
}

func NewShipmentRouter(db *gorm.DB, rg *routing.RouterGroup) ShipmentRouter {
	router := rg.Group("/shipment")
	sr := &database.ShipmentController{DB: db}
	return ShipmentRouter{sr, router}
}

func (c *ShipmentRouter) AddGetRoute()  {
	c.rg.GET("/all", c.shipmentController.GetAllShipments)
}

func (c *ShipmentRouter) AddPostRoute() {
	c.rg.POST("/", c.shipmentController.CreateShipment)
}
