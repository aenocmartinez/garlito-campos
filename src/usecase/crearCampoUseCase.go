package usecase

import "garlito/src/view/dto"

type CrearCampoUseCase struct{}

func (uc *CrearCampoUseCase) Execute(campoDto dto.CampoDto) (resp dto.ResponseHttp) {

	campo := campoRepository.BuscarCampoPorNombre(campoDto.Nombre)
	if campo.Existe() {
		resp.Code = "200"
		resp.Message = "the resource already exists"
		return resp
	}

	campo.SetRepository(campoRepository)
	campo.SetNombre(campoDto.Nombre)
	campo.SetDescripcion(campoDto.Descripcion)

	err := campo.Crear()
	if err != nil {
		resp.Code = "500"
		resp.Message = "error internal"
		return resp
	}

	resp.Code = "201"
	resp.Message = "success"
	return resp
}
