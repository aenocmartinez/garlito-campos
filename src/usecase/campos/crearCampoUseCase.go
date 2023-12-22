package usecase

import (
	"garlito/src/domain"
	"garlito/src/view/dto"
)

type CrearCampoUseCase struct{}

func (uc *CrearCampoUseCase) Execute(campoDto dto.CampoDto) (resp dto.ResponseHttp) {

	var nuevoCampo domain.Campo
	var campo domain.Campo = campoRepository.BuscarCampoPorNombre(campoDto.Nombre)
	if campo.Existe() {
		resp.Code = "200"
		resp.Message = "the resource already exists"
		return resp
	}

	nuevoCampo = domain.NewSimple()
	if campoDto.EsCompuesto {
		nuevoCampo = domain.NewCompuesto()
	}

	nuevoCampo.SetRepository(campoRepository)
	nuevoCampo.SetNombre(campoDto.Nombre)
	nuevoCampo.SetDescripcion(campoDto.Descripcion)

	err := nuevoCampo.Crear()
	if err != nil {
		resp.Code = "500"
		resp.Message = "error internal"
		return resp
	}

	resp.Code = "201"
	resp.Message = "success"
	return resp
}
