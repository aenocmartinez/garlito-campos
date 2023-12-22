package main

import (
	"garlito/src/view/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	r.GET("/campos", controller.ListarCampos)
	r.GET("/campos/:id", controller.BuscarCampoPorId)
	r.POST("/campos", controller.CrearCampo)
	r.PUT("/campos", controller.ActualizarCampo)
	r.DELETE("/campos/:id", controller.EliminarCampo)
	r.PUT("/campos/agregar-subcampo", controller.AgregarSubcampo)

	r.Run(":2300")
}
