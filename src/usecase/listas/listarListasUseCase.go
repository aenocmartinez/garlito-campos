package usecase

import "garlito/src/view/dto"

type ListarListasUseCase struct{}

func (uc *ListarListasUseCase) Execute() (resp dto.ResponseHttp) {

	resp.Code = "200"
	resp.Data = listaRepository.ListarLista()
	return resp
}
