package domain

import "garlito/src/view/dto"

type CampoRepository interface {
	CrearCampo(campo Campo) (err error)
	EliminarCampo(id int64) (err error)
	ActualizarCampo(campo Campo) (err error)
	AgregarSuncampo(campo Campo, subcampo Campo) (err error)
	QuitarSuncampo(campo Campo, subcampo Campo) (err error)
	SubCampos(campo Campo) []Campo
	ListarCampos() []dto.CampoDto
	BuscarCampoPorId(id int64) Campo
	BuscarCampo(criteria string) []Campo
	BuscarCampoPorNombre(nombre string) Campo
}
