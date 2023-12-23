package usecase

import "garlito/src/view/dto"

type BuscarListaPorIdUseCase struct{}

func (uc *BuscarListaPorIdUseCase) Execute(id int64) (resp dto.ResponseHttp) {
	lista := listaRepository.BuscarListaPorId(id)
	if !lista.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	resp.Code = "200"
	resp.Data = dto.ListaDto{
		Id:      lista.Id(),
		Nombre:  lista.Nombre(),
		Valores: lista.Valores(),
	}
	return resp
}
