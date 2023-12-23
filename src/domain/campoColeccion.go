package domain

type CampoColeccion struct {
	id          int64
	coleccion   Coleccion
	campo       Campo
	propiedades map[string]bool
}

func NewCampoColeccion(coleccion Coleccion, campo Campo, propiedades map[string]bool) *CampoColeccion {
	return &CampoColeccion{
		campo:       campo,
		coleccion:   coleccion,
		propiedades: propiedades,
	}
}

func (cc *CampoColeccion) SetId(id int64) {
	cc.id = id
}

func (cc *CampoColeccion) Id() int64 {
	return cc.id
}

func (cc *CampoColeccion) Coleccion() string {
	return cc.coleccion.Nombre()
}

func (cc *CampoColeccion) ColeccionId() int64 {
	return cc.coleccion.Id()
}

func (cc *CampoColeccion) Nombre() string {
	return cc.campo.Nombre()
}

func (cc *CampoColeccion) CampoId() int64 {
	return cc.campo.Id()
}

func (cc *CampoColeccion) Descripcion() string {
	return cc.campo.Descripcion()
}

func (cc *CampoColeccion) EsCompuesto() bool {
	return cc.campo.EsCompuesto()
}

func (cc *CampoColeccion) SetObligatorio(obligatorio bool) {
	cc.propiedades["obligatorio"] = obligatorio
}

func (cc *CampoColeccion) SetEditable(editable bool) {
	cc.propiedades["editable"] = editable
}

func (cc *CampoColeccion) SetUnico(unico bool) {
	cc.propiedades["unico"] = unico
}

func (cc *CampoColeccion) Unico() bool {
	return cc.propiedades["unico"]
}

func (cc *CampoColeccion) Obligatorio() bool {
	return cc.propiedades["obligatorio"]
}

func (cc *CampoColeccion) Editable() bool {
	return cc.propiedades["editable"]
}
