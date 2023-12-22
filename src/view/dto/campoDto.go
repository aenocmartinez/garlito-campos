package dto

type CampoDto struct {
	Id          int64      `json:"id"`
	Nombre      string     `json:"nombre"`
	Descripcion string     `json:"descripcion"`
	Subcampos   []CampoDto `json:"subcampos,omitempty"`
	EsCompuesto bool       `json:"es_compuesto"`
}
