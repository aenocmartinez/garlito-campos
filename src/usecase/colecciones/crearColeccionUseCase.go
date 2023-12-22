package usecase

import "garlito/src/view/dto"

type CrearColeccionUseCase struct{}

func (uc *CrearColeccionUseCase) Execute(nombre string) (resp dto.ResponseHttp) {

	coleccion := coleccionRepository.BuscarColeccionPorNombre(nombre)
	if coleccion.Existe() {
		resp.Code = "200"
		resp.Message = "the resource already exists"
		return resp
	}

	coleccion.SetRepository(coleccionRepository)
	coleccion.SetNombre(nombre)
	err := coleccion.Crear()
	if err != nil {
		resp.Code = "500"
		resp.Message = "error server internal"
		return resp
	}

	resp.Code = "201"
	resp.Message = "success"
	return resp
}
