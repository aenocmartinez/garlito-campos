package usecase

import (
	"garlito/src/domain"
	"garlito/src/view/dto"
)

type QuitarSubcampoUseCase struct{}

func (uc *QuitarSubcampoUseCase) Execute(campoId, subcampoId int64) (resp dto.ResponseHttp) {

	campo := campoRepository.BuscarCampoPorId(campoId)
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
	err := campoCompuesto.QuitarSubcampo(subcampo)
	if err != nil {
		resp.Code = "500"
		resp.Message = "error server internal"
		return resp
	}

	resp.Code = "200"
	resp.Message = "success"
	return resp
}
