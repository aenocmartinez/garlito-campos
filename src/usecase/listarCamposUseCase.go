package usecase

import "garlito/src/view/dto"

type ListarCamposUseCase struct{}

func (uc *ListarCamposUseCase) Execute() (resp dto.EesponseHttp) {

	resp.Code = "200"
	resp.Data = campoRepository.ListarCampos()

	return resp
}
