package formrequest

type CrearColeccionFormRquest struct {
	Nombre string `json:"nombre" binding:"required"`
}

type ActualizarColeccionFormRquest struct {
	Id     int64  `json:"id" binding:"required"`
	Nombre string `json:"nombre" binding:"required"`
}

type AgregarCampoColeccionFormRequest struct {
	ColeccionId int64 `json:"coleccion_id" binding:"required"`
	CampoId     int64 `json:"campo_id" binding:"required"`
	Obligatorio bool  `json:"es_obligatorio"`
	Editable    bool  `json:"es_editable"`
	Unico       bool  `json:"es_unico"`
}
