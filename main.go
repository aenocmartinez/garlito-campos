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
	r.PUT("/campos/quitar-subcampo", controller.QuitarSubcampo)

	r.GET("/colecciones", controller.ListarColecciones)
	r.GET("/colecciones/:id", controller.BuscarColeccionPorId)
	r.POST("/colecciones", controller.CrearColeccion)
	r.PUT("/colecciones", controller.ActualizarColeccion)
	r.DELETE("/colecciones/:id", controller.EliminarColeccion)
	r.PUT("/colecciones/agregar-campo", controller.AgregarCampoColeccion)
	r.PUT("/colecciones/quitar-campo", controller.QuitarCampoColeccion)

	r.Run(":2300")
}
