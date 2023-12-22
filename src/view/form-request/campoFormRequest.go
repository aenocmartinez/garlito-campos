package formrequest

type CrearCampoFormRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion"`
	EsCompuesto bool   `json:"es_compuesto"`
}

type ActualizarCampoFormRequest struct {
	Id          int64  `json:"id" binding:"required"`
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion"`
}

type AgregarSubcampoFormRequest struct {
	CampoId    int64 `json:"campoId" binding:"required"`
	SubcampoId int64 `json:"subcampoId" binding:"required"`
}
