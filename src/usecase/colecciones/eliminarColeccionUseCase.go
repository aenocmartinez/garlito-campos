package usecase

import "garlito/src/view/dto"

type EliminarColeccionUseCase struct{}

func (uc *EliminarColeccionUseCase) Execute(id int64) (resp dto.ResponseHttp) {
	coleccion := coleccionRepository.BuscarColeccionPorId(id)
	if !coleccion.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	coleccion.SetRepository(coleccionRepository)

	if err := coleccion.Eliminar(); err != nil {
		resp.Code = "500"
		resp.Message = "error server internal"
		return resp
	}

	resp.Code = "200"
	resp.Message = "success"
	return resp
}
