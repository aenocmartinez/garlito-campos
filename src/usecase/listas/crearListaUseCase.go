package usecase

import "garlito/src/view/dto"

type CrearListaUseCase struct{}

func (uc *CrearListaUseCase) Execute(data dto.ListaDto) (resp dto.ResponseHttp) {

	lista := listaRepository.BuscarListaPorNombre(data.Nombre)
	if lista.Existe() {
		resp.Code = "200"
		resp.Message = "the resource already exists"
		return resp
	}

	lista.SetRepository(listaRepository)
	lista.SetNombre(data.Nombre)
	lista.SetValores(data.Valores)

	if err := lista.Crear(); err != nil {
		resp.Code = "500"
		resp.Message = "server error internal"
		return resp
	}

	resp.Code = "201"
	resp.Message = "success"
	return resp
}
