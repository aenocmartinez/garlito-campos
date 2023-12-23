package usecase

import (
	"garlito/src/dao/mock"
	"garlito/src/domain"
)

var coleccionRepository domain.ColeccionRepository = mock.NewColeccionDao()
var campoRepository domain.CampoRepository = mock.NewCampoDao()
