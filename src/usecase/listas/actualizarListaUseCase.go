package usecase

import "garlito/src/view/dto"

type ActualizarListaUseCase struct{}

func (uc *ActualizarListaUseCase) Execute(data dto.ListaDto) (resp dto.ResponseHttp) {

	lista := listaRepository.BuscarListaPorId(data.Id)
	if !lista.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	lista.SetRepository(listaRepository)
	lista.SetNombre(data.Nombre)
	lista.SetValores(data.Valores)

	if err := lista.Actualizar(); err != nil {
		resp.Code = "500"
		resp.Message = "server error internal"
		return resp
	}

	resp.Code = "200"
	resp.Message = "success"
	return resp
}
