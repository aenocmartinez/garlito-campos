package usecase

import (
	"garlito/src/domain"
	"garlito/src/view/dto"
)

type AgregarSubcampoUseCase struct{}

func (uc *AgregarSubcampoUseCase) Execute(campoId, subcampoId int64) (resp dto.ResponseHttp) {

	var campo domain.Campo = campoRepository.BuscarCampoPorId(campoId)
	if !campo.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	subcampo := campoRepository.BuscarCampoPorId(subcampoId)
	if !subcampo.Existe() {
		resp.Code = "404"
		resp.Message = "resource not found"
		return resp
	}

	// campoCompuesto, _ := campo.(*domain.Compuesto)
	// campoCompuesto.AgregarSubcampo(subcampo)

	return resp
}
