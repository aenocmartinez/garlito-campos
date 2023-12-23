package controller

import (
	"garlito/src/infraestructure/util"
	usecase "garlito/src/usecase/colecciones"
	"garlito/src/view/dto"
	formrequest "garlito/src/view/form-request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListarColecciones(c *gin.Context) {
	listarColecciones := usecase.ListarColeccionesUseCase{}
	resp := listarColecciones.Execute()

	c.JSON(http.StatusOK, resp)
}

func BuscarColeccionPorId(c *gin.Context) {
	idParam := c.Param("id")
	if !util.EsNumero(idParam) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param invalid"})
		return
	}

	id := util.ConverToInt64(idParam)

	buscarColeccion := usecase.BuscarColeccionPorIdUseCase{}
	resp := buscarColeccion.Execute(id)

	c.JSON(http.StatusOK, resp)
}

func CrearColeccion(c *gin.Context) {
	var req formrequest.CrearColeccionFormRquest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	crearColeccion := usecase.CrearColeccionUseCase{}
	resp := crearColeccion.Execute(req.Nombre)

	code, _ := strconv.Atoi(resp.Code)
	c.JSON(code, resp)
}

func ActualizarColeccion(c *gin.Context) {
	var req formrequest.ActualizarColeccionFormRquest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	crearColeccion := usecase.ActualizarColeccionUseCase{}
	resp := crearColeccion.Execute(dto.ColeccionDto{
		Id:     req.Id,
		Nombre: req.Nombre,
	})

	code, _ := strconv.Atoi(resp.Code)
	c.JSON(code, resp)
}

func EliminarColeccion(c *gin.Context) {
	idParam := c.Param("id")
	if !util.EsNumero(idParam) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param invalid"})
		return
	}

	id := util.ConverToInt64(idParam)

	crearColeccion := usecase.EliminarColeccionUseCase{}
	resp := crearColeccion.Execute(id)

	code, _ := strconv.Atoi(resp.Code)
	c.JSON(code, resp)
}

func AgregarCampoColeccion(c *gin.Context) {
	var req formrequest.AgregarCampoColeccionFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	agregarCampoColeccion := usecase.AgregarCampoColeccionUseCase{}
	resp := agregarCampoColeccion.Execute(dto.CampoColeccionDto{
		ColeccionId: req.ColeccionId,
		CampoId:     req.CampoId,
		Obligatorio: req.Obligatorio,
		Editable:    req.Editable,
		Unico:       req.Unico,
	})

	code, _ := strconv.Atoi(resp.Code)
	c.JSON(code, resp)
}

func QuitarCampoColeccion(c *gin.Context) {
	var req formrequest.AgregarCampoColeccionFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	quitarCampoColeccion := usecase.QuitarCampoColeccionUseCase{}
	resp := quitarCampoColeccion.Execute(dto.CampoColeccionDto{
		ColeccionId: req.ColeccionId,
		CampoId:     req.CampoId,
	})

	code, _ := strconv.Atoi(resp.Code)
	c.JSON(code, resp)
}
