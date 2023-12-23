package dto

type ColeccionDto struct {
	Id     int64               `json:"id,omitempty"`
	Nombre string              `json:"nombre,omitempty"`
	Campos []CampoColeccionDto `json:"campos,omitempty"`
}
