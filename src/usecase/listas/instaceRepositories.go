package usecase

import (
	"garlito/src/dao/mock"
	"garlito/src/domain"
)

var listaRepository domain.ListaRepository = mock.NewListaDao()
