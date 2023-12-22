package usecase

import "garlito/src/view/dto"

type ListarColeccionesUseCase struct{}

func (uc *ListarColeccionesUseCase) Execute() (resp dto.ResponseHttp) {

	resp.Code = "200"
	resp.Data = coleccionRepository.ListarColeccion()

	return resp
}
