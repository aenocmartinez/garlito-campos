package usecase

import "garlito/src/view/dto"

type EliminarCampoUseCase struct{}

func (uc *EliminarCampoUseCase) Execute(id int64) (resp dto.ResponseHttp) {

	campo := campoRepository.BuscarCampoPorId(id)
	if !campo.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	campo.SetRepository(campoRepository)
	err := campo.Eliminar()
	if err != nil {
		resp.Code = "500"
		resp.Message = "error server internal"
		return resp
	}

	resp.Code = "200"
	resp.Message = "success"
	return resp
}
