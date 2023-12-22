package mock

import (
	"garlito/src/domain"
	"garlito/src/view/dto"
	"strings"
)

var (
	nombres []string = []string{"Arte", "Historia"}
)

type ColeccionDao struct {
	data []domain.Coleccion
}

func NewColeccionDao() *ColeccionDao {
	return &ColeccionDao{
		data: loadData(),
	}
}

func loadData() (colecciones []domain.Coleccion) {
	colecciones = []domain.Coleccion{}
	for index, nombre := range nombres {
		coleccion := domain.NewColeccion()
		coleccion.SetId(int64(index + 1))
		coleccion.SetNombre(nombre)
		colecciones = append(colecciones, *coleccion)
	}
	return colecciones
}

func (c *ColeccionDao) CrearColeccion(coleccion domain.Coleccion) (err error) {
	coleccion.SetId(int64(len(c.data) + 1))
	c.data = append(c.data, coleccion)
	return err
}

func (c *ColeccionDao) EliminarColeccion(id int64) (err error) {
	var colecciones []domain.Coleccion = []domain.Coleccion{}
	for _, coleccion := range c.data {
		if coleccion.Id() == id {
			continue
		}
		colecciones = append(colecciones, coleccion)
	}

	c.data = colecciones
	return err
}

func (c *ColeccionDao) ActualizarColeccion(coleccion domain.Coleccion) (err error) {
	var colecciones []domain.Coleccion = []domain.Coleccion{}
	for _, item := range c.data {
		if item.Id() == coleccion.Id() {
			item.SetNombre(coleccion.Nombre())
		}
		colecciones = append(colecciones, item)
	}

	c.data = colecciones

	return err
}

func (c *ColeccionDao) ListarColeccion() (colecciones []dto.ColeccionDto) {
	for _, coleccion := range c.data {
		colecciones = append(colecciones, dto.ColeccionDto{
			Id:     coleccion.Id(),
			Nombre: coleccion.Nombre(),
		})
	}
	return colecciones
}

func (c *ColeccionDao) BuscarColeccionPorNombre(nombre string) (coleccion domain.Coleccion) {
	for _, coleccion := range c.data {
		if strings.EqualFold(coleccion.Nombre(), nombre) {
			return coleccion
		}
	}
	return *domain.NewColeccion()
}

func (c *ColeccionDao) BuscarColeccionPorId(id int64) (coleccion domain.Coleccion) {
	for _, coleccion := range c.data {
		if coleccion.Id() == id {
			return coleccion
		}
	}
	return *domain.NewColeccion()
}
