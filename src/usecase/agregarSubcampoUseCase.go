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

	if !campo.EsCompuesto() {
		resp.Code = "200"
		resp.Message = "the resource is not composite"
		return resp
	}

	campoCompuesto, _ := campo.(*domain.Compuesto)

	campoCompuesto.SetRepository(campoRepository)
	err := campoCompuesto.AgregarSubcampo(subcampo)
	if err != nil {
		resp.Code = "500"
		resp.Message = "error server internal"
		return resp
	}

	resp.Code = "201"
	resp.Message = "success"
	return resp
}
