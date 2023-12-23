package usecase

import "garlito/src/view/dto"

type EliminarListaUseCase struct{}

func (uc *EliminarListaUseCase) Execute(id int64) (resp dto.ResponseHttp) {

	lista := listaRepository.BuscarListaPorId(id)
	if !lista.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	lista.SetRepository(listaRepository)
	if err := lista.Eliminar(); err != nil {
		resp.Code = "500"
		resp.Message = "server error internal"
		return resp
	}

	resp.Code = "200"
	resp.Message = "success"
	return resp
}
