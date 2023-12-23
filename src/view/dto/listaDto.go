package dto

type ListaDto struct {
	Id      int64    `json:"id"`
	Nombre  string   `json:"nombre"`
	Valores []string `json:"valores,omitempty"`
}
