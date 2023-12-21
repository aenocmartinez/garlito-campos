package usecase

import "garlito/src/view/dto"

type EliminarCampoUseCase struct{}

func (uc *EliminarCampoUseCase) Execute(id int64) (resp dto.EesponseHttp) {

	campo := campoRepository.BuscarCampoPorId(id)
	if !campo.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	campo.SetRepository(campoRepository)
	exito := campo.Eliminar()
	if !exito {
		resp.Code = "500"
		resp.Message = "error server internal"
		return resp
	}

	resp.Code = "200"
	resp.Message = "success"
	return resp
}
