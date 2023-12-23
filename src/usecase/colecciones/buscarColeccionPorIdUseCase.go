package usecase

import "garlito/src/view/dto"

type BuscarColeccionPorIdUseCase struct{}

func (uc *BuscarColeccionPorIdUseCase) Execute(id int64) (resp dto.ResponseHttp) {
	coleccion := coleccionRepository.BuscarColeccionPorId(id)
	if !coleccion.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	campos := []dto.CampoColeccionDto{}
	for _, campoColeccion := range coleccion.Campos() {
		campos = append(campos, dto.CampoColeccionDto{
			ColeccionId: campoColeccion.ColeccionId(),
			CampoId:     campoColeccion.CampoId(),
			Campo:       campoColeccion.Nombre(),
			Obligatorio: campoColeccion.Obligatorio(),
			Editable:    campoColeccion.Editable(),
			Unico:       campoColeccion.Unico(),
		})
	}

	resp.Code = "200"
	resp.Message = "success"
	resp.Data = dto.ColeccionDto{
		Id:     coleccion.Id(),
		Nombre: coleccion.Nombre(),
		Campos: campos,
	}
	return resp
}
