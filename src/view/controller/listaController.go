package controller

import (
	"garlito/src/infraestructure/util"
	usecase "garlito/src/usecase/listas"
	"garlito/src/view/dto"
	formrequest "garlito/src/view/form-request"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func ListarListas(c *gin.Context) {
	listarListasUseCase := usecase.ListarListasUseCase{}
	resp := listarListasUseCase.Execute()
	c.JSON(http.StatusOK, resp)
}

func BuscarListaPorId(c *gin.Context) {
	idParam := c.Param("id")
	if !util.EsNumero(idParam) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param invalid"})
		return
	}

	id := util.ConverToInt64(idParam)
	buscarLista := usecase.BuscarListaPorIdUseCase{}
	resp := buscarLista.Execute(id)

	c.JSON(http.StatusOK, resp)
}

func CrearLista(c *gin.Context) {
	var req formrequest.CrearListaFormRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	valores := strings.Split(req.Valores, ",")

	crearLista := usecase.CrearListaUseCase{}
	resp := crearLista.Execute(dto.ListaDto{
		Nombre:  req.Nombre,
		Valores: valores,
	})

	code, _ := strconv.Atoi(resp.Code)
	c.JSON(code, resp)
}

func ActualizarLista(c *gin.Context) {
	var req formrequest.ActualizarListaFormRequest
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	valores := strings.Split(req.Valores, ",")

	actualizarLista := usecase.ActualizarListaUseCase{}
	resp := actualizarLista.Execute(dto.ListaDto{
		Id:      req.Id,
		Nombre:  req.Nombre,
		Valores: valores,
	})

	code, _ := strconv.Atoi(resp.Code)
	c.JSON(code, resp)
}

func EliminarLista(c *gin.Context) {
	idParam := c.Param("id")
	if !util.EsNumero(idParam) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "param invalid"})
		return
	}

	id := util.ConverToInt64(idParam)
	eliminarLista := usecase.EliminarListaUseCase{}
	resp := eliminarLista.Execute(id)

	c.JSON(http.StatusOK, resp)
}
