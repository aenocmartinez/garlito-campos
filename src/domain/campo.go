package domain

type Campo interface {
	SetId(id int64)
	Id() int64
	SetNombre(nombre string)
	Nombre() string
	SetDescripcion(descripcion string)
	Descripcion() string
	EsCompuesto() bool
	Crear() bool
	Eliminar() bool
	Actualizar() bool
	Existe() bool
	SetRepository(repository CampoRepository)
}
