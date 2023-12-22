package formrequest

type AgregarSubcampoFormRequest struct {
	CampoId    int64 `json:"campoId" binding:"required"`
	SubcampoId int64 `json:"subcampoId" binding:"required"`
}
