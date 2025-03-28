package controller

import(
	"fmt"
	"github.com/gin-gonic/gin"
)

type MainController interface {
	initController() error
}

type MainControllerImpl struct {}

func (impl *MainControllerImpl) initController() {
	r := gin.Default()
	r.Run()
	fmt.Println("Controllers Initialized!")
}