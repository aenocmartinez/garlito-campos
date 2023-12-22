package usecase

import "garlito/src/view/dto"

type ActualizarColeccionUseCase struct{}

func (uc *ActualizarColeccionUseCase) Execute(coleccionDto dto.ColeccionDto) (resp dto.ResponseHttp) {

	coleccion := coleccionRepository.BuscarColeccionPorId(coleccionDto.Id)
	if !coleccion.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	coleccion.SetRepository(coleccionRepository)
	coleccion.SetNombre(coleccionDto.Nombre)

	if err := coleccion.Actualizar(); err != nil {
		resp.Code = "500"
		resp.Message = "server error internal"
		return resp
	}

	resp.Code = "200"
	resp.Message = "success"
	return resp
}
