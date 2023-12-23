package domain

type Coleccion struct {
	id         int64
	nombre     string
	campos     []CampoColeccion
	repository ColeccionRepository
}

type Propiedad struct {
}

func NewColeccion() *Coleccion {
	return &Coleccion{
		campos: []CampoColeccion{},
	}
}

func (c *Coleccion) SetRepository(repository ColeccionRepository) {
	c.repository = repository
}

func (c *Coleccion) SetId(id int64) {
	c.id = id
}

func (c *Coleccion) Id() int64 {
	return c.id
}

func (c *Coleccion) SetNombre(nombre string) {
	c.nombre = nombre
}

func (c *Coleccion) SetCampos(campos []CampoColeccion) {
	c.campos = campos
}

func (c *Coleccion) Nombre() string {
	return c.nombre
}

func (c *Coleccion) Crear() error {
	return c.repository.CrearColeccion(*c)
}

func (c *Coleccion) Eliminar() error {
	return c.repository.EliminarColeccion(c.id)
}

func (c *Coleccion) Actualizar() error {
	return c.repository.ActualizarColeccion(*c)
}

func (c *Coleccion) Existe() bool {
	return c.id > 0
}

func (c *Coleccion) AgregarCampo(campo Campo, propiedades map[string]bool) error {
	c.campos = append(c.campos, *NewCampoColeccion(*c, campo, propiedades))
	return c.repository.ActualizarColeccion(*c)
}

func (c *Coleccion) QuitarCampo(campo Campo) error {
	var campos []CampoColeccion = []CampoColeccion{}
	for _, campoColeccion := range c.campos {
		if campoColeccion.CampoId() == campo.Id() {
			continue
		}
		campos = append(campos, campoColeccion)
	}

	c.campos = campos

	return c.repository.ActualizarColeccion(*c)
}

func (c *Coleccion) Campos() []CampoColeccion {
	return c.campos
}
