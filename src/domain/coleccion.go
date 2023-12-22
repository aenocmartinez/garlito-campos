package domain

type Coleccion struct {
	id         int64
	nombre     string
	repository ColeccionRepository
}

func NewColeccion() *Coleccion {
	return &Coleccion{}
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
