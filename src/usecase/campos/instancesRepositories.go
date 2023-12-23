package usecase

import (
	"garlito/src/dao/mock"
	"garlito/src/domain"
)

var campoRepository domain.CampoRepository = mock.NewCampoDao()
