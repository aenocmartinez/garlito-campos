package usecase

import (
	"garlito/src/domain"
	"garlito/src/view/dto"
)

type BuscarCampoPorIdUseCase struct{}

func (uc *BuscarCampoPorIdUseCase) Execute(id int64) (resp dto.ResponseHttp) {
	campo := campoRepository.BuscarCampoPorId(id)
	if !campo.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	var subcampos []dto.CampoDto
	if campo.EsCompuesto() {
		campoCompuesto, _ := campo.(*domain.Compuesto)
		subcampos = campoCompuesto.SubcamposDto()
	}

	resp.Code = "200"
	resp.Data = dto.CampoDto{
		Id:          campo.Id(),
		Nombre:      campo.Nombre(),
		Descripcion: campo.Descripcion(),
		Subcampos:   subcampos,
	}

	return resp
}
