package usecase

import "garlito/src/view/dto"

type QuitarCampoColeccionUseCase struct{}

func (uc *QuitarCampoColeccionUseCase) Execute(data dto.CampoColeccionDto) (resp dto.ResponseHttp) {

	coleccion := coleccionRepository.BuscarColeccionPorId(data.ColeccionId)
	if !coleccion.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	campo := campoRepository.BuscarCampoPorId(data.CampoId)
	if !campo.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	coleccion.SetRepository(coleccionRepository)
	err := coleccion.QuitarCampo(campo)
	if err != nil {
		resp.Code = "500"
		resp.Message = "server error internal"
		return resp
	}

	resp.Code = "200"
	resp.Message = "success"
	return resp
}
