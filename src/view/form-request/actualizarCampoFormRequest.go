package formrequest

type ActualizarCampoFormRequest struct {
	Id          int64  `json:"id" binding:"required"`
	Nombre      string `json:"nombre" binding:"required"`
	Descripcion string `json:"descripcion"`
}
