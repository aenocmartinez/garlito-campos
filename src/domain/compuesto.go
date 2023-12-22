package domain

import (
	"garlito/src/view/dto"
)

type Compuesto struct {
	id          int64
	nombre      string
	descripcion string
	subcampos   []Campo
	repository  CampoRepository
}

func NewCompuesto() *Compuesto {
	return &Compuesto{}
}

func (c *Compuesto) SetRepository(repository CampoRepository) {
	c.repository = repository
}

func (c *Compuesto) SetId(id int64) {
	c.id = id
}

func (c *Compuesto) Id() int64 {
	return c.id
}

func (c *Compuesto) SetNombre(nombre string) {
	c.nombre = nombre
}

func (c *Compuesto) Nombre() string {
	return c.nombre
}

func (c *Compuesto) SetDescripcion(descripcion string) {
	c.descripcion = descripcion
}

func (c *Compuesto) Descripcion() string {
	return c.descripcion
}

func (c *Compuesto) EsCompuesto() bool {
	return true
}

func (c *Compuesto) AgregarSubcampo(subcampo Campo) error {
	c.subcampos = append(c.subcampos, subcampo)
	return c.repository.ActualizarCampo(c)
}

func (c *Compuesto) QuitarSubcampo(campo Campo) bool {
	return false
}

func (c *Compuesto) Subcampos() []Campo {
	return c.subcampos
}

func (c *Compuesto) SubcamposDto() []dto.CampoDto {
	var subcampos []dto.CampoDto = []dto.CampoDto{}
	for _, subcampo := range c.Subcampos() {
		var subcampoDto dto.CampoDto
		if subcampo.EsCompuesto() {
			campoCompuesto, _ := subcampo.(*Compuesto)
			subcampoDto.Id = campoCompuesto.Id()
			subcampoDto.Nombre = campoCompuesto.Nombre()
			subcampoDto.Descripcion = campoCompuesto.Descripcion()
		} else {
			campoSimple, _ := subcampo.(*Simple)
			subcampoDto.Id = campoSimple.Id()
			subcampoDto.Nombre = campoSimple.Nombre()
			subcampoDto.Descripcion = campoSimple.Descripcion()
		}
		subcampos = append(subcampos, subcampoDto)
	}
	return subcampos
}

func (c *Compuesto) Crear() error {
	return c.repository.CrearCampo(c)
}

func (c *Compuesto) Eliminar() error {
	return c.repository.EliminarCampo(c.id)
}

func (c *Compuesto) Actualizar() error {
	return c.repository.ActualizarCampo(c)
}

func (c *Compuesto) Existe() bool {
	return c.id > 0
}
