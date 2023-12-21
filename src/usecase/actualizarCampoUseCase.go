package usecase

import (
	"garlito/src/view/dto"
)

type ActualizarCampoUseCase struct{}

func (uc *ActualizarCampoUseCase) Execute(campoDto dto.CampoDto) (resp dto.ResponseHttp) {
	campo := campoRepository.BuscarCampoPorId(campoDto.Id)
	if !campo.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	campo.SetRepository(campoRepository)
	campo.SetNombre(campoDto.Nombre)
	campo.SetDescripcion(campoDto.Descripcion)
	err := campo.Actualizar()
	if err != nil {
		resp.Code = "500"
		resp.Message = "error server internal"
		return resp
	}

	resp.Code = "200"
	resp.Message = "success"
	return resp
}
