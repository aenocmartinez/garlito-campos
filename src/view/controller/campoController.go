package controller

import (
	"garlito/src/infraestructure/util"
	usecase "garlito/src/usecase/campos"
	"garlito/src/view/dto"
	formrequest "garlito/src/view/form-request"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CrearCampo(c *gin.Context) {
	var req formrequest.CrearCampoFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	crearCampo := usecase.CrearCampoUseCase{}
	resp := crearCampo.Execute(dto.CampoDto{
		Nombre:      req.Nombre,
		Descripcion: req.Descripcion,
		EsCompuesto: req.EsCompuesto,
	})

	c.JSON(http.StatusCreated, resp)
}

func ActualizarCampo(c *gin.Context) {
	var req formrequest.ActualizarCampoFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	actualizarCampo := usecase.ActualizarCampoUseCase{}
	resp := actualizarCampo.Execute(dto.CampoDto{
		Id:          req.Id,
		Nombre:      req.Nombre,
		Descripcion: req.Descripcion,
	})

	c.JSON(http.StatusOK, resp)
}

func EliminarCampo(c *gin.Context) {
	idParam := c.Param("id")
	if !util.EsNumero(idParam) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param invalid"})
		return
	}

	id := util.ConverToInt64(idParam)

	eliminarCampo := usecase.EliminarCampoUseCase{}
	resp := eliminarCampo.Execute(id)
	c.JSON(http.StatusOK, resp)
}

func ListarCampos(c *gin.Context) {
	listarCampos := usecase.ListarCamposUseCase{}
	resp := listarCampos.Execute()
	c.JSON(http.StatusOK, resp)
}

func BuscarCampoPorId(c *gin.Context) {
	idParam := c.Param("id")
	if !util.EsNumero(idParam) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param invalid"})
		return
	}

	id := util.ConverToInt64(idParam)

	buscarCampo := usecase.BuscarCampoPorIdUseCase{}
	resp := buscarCampo.Execute(id)
	code, _ := strconv.Atoi(resp.Code)
	if resp.Code != "200" {
		c.JSON(code, resp)
		return
	}

	c.JSON(code, resp)
}

func AgregarSubcampo(c *gin.Context) {
	var req formrequest.AgregarSubcampoFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	agregarSubcampo := usecase.AgregarSubcampoUseCase{}
	resp := agregarSubcampo.Execute(req.CampoId, req.SubcampoId)

	code, _ := strconv.Atoi(resp.Code)
	if resp.Code != "200" {
		c.JSON(code, resp)
		return
	}

	c.JSON(code, resp)
}

func QuitarSubcampo(c *gin.Context) {
	var req formrequest.AgregarSubcampoFormRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	quitarSubcampo := usecase.QuitarSubcampoUseCase{}
	resp := quitarSubcampo.Execute(req.CampoId, req.SubcampoId)

	code, _ := strconv.Atoi(resp.Code)
	if resp.Code != "200" {
		c.JSON(code, resp)
		return
	}

	c.JSON(code, resp)
}
