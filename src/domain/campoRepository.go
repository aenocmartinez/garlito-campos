package domain

import "garlito/src/view/dto"

type CampoRepository interface {
	CrearCampo(campo Campo) bool
	EliminarCampo(id int64) bool
	ActualizarCampo(campo Campo) bool
	AgregarSuncampo(campo Campo, subcampo Campo) bool
	QuitarSuncampo(campo Campo, subcampo Campo) bool
	SubCampos(campo Campo) []Campo
	ListarCampos() []dto.CampoDto
	BuscarCampoPorId(id int64) Campo
	BuscarCampo(criteria string) []Campo
	BuscarCampoPorNombre(nombre string) Campo
}
