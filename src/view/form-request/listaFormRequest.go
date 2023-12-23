package formrequest

type CrearListaFormRequest struct {
	Nombre  string `json:"nombre" binding:"required"`
	Valores string `json:"valores"`
}

type ActualizarListaFormRequest struct {
	Id      int64  `json:"id" binding:"required"`
	Nombre  string `json:"nombre" binding:"required"`
	Valores string `json:"valores"`
}

type AgregarValorListaFormRequest struct {
	Id    int64  `json:"lista_id" binding:"required"`
	Valor string `json:"valor" binding:"required"`
}
