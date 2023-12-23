package dto

type CampoColeccionDto struct {
	Coleccion   string `json:"coleccion,omitempty"`
	ColeccionId int64  `json:"coleccion_id"`
	Campo       string `json:"campo,omitempty"`
	CampoId     int64  `json:"campo_id"`
	Obligatorio bool   `json:"obligatorio"`
	Editable    bool   `json:"editable"`
	Unico       bool   `json:"unico"`
}
