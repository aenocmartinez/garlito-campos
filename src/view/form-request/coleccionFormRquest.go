package formrequest

type CrearColeccionFormRquest struct {
	Nombre string `json:"nombre" binding:"required"`
}

type ActualizarColeccionFormRquest struct {
	Id     int64  `json:"id" binding:"required"`
	Nombre string `json:"nombre" binding:"required"`
}
