package domain

import "garlito/src/view/dto"

type CampoRepository interface {
	CrearCampo(campo Campo) (err error)
	EliminarCampo(id int64) (err error)
	ActualizarCampo(campo Campo) (err error)
	SubCampos(campo Campo) []Campo
	ListarCampos() []dto.CampoDto
	BuscarCampoPorId(id int64) Campo
	BuscarCampo(criteria string) []Campo
	BuscarCampoPorNombre(nombre string) Campo
}

type ColeccionRepository interface {
	CrearColeccion(coleccion Coleccion) (err error)
	EliminarColeccion(id int64) (err error)
	ActualizarColeccion(coleccion Coleccion) (err error)
	ListarColeccion() (colecciones []dto.ColeccionDto)
	BuscarColeccionPorNombre(nombre string) (coleccion Coleccion)
	BuscarColeccionPorId(id int64) (coleccion Coleccion)
}

type ListaRepository interface {
	CrearLista(lista Lista) (err error)
	ActualizarLista(lista Lista) (err error)
	EliminarLista(id int64) (err error)
	ListarLista() (lista []dto.ListaDto)
	BuscarListaPorId(id int64) (lista Lista)
	BuscarListaPorNombre(nombre string) (lista Lista)
}
