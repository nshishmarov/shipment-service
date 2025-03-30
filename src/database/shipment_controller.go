package database

import (
	"net/http"
	"service/src/api/dto"
	"service/src/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ShipmentController struct {
	DB *gorm.DB
}

// Migration
func (c *ShipmentController) Migrate() {
	c.DB.AutoMigrate(&model.ShipmentEntity{})
}

func (sr *ShipmentController) GetAllShipments(ctx *gin.Context) {
	var shipment model.ShipmentEntity
	var shipments = sr.DB.Find(&shipment)
	ctx.IndentedJSON(http.StatusOK, shipments)
}

func (sr *ShipmentController) CreateShipment(ctx *gin.Context) {
	var shipment dto.Shipment
	if err := ctx.ShouldBindJSON(&shipment); err != nil {
		ctx.Error(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	ctx.JSON(http.StatusOK, shipment)
	sr.DB.Create(&shipment)
}