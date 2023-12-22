package formrequest

type CrearCampoFormRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion"`
	EsCompuesto bool   `json:"es_compuesto"`
}
