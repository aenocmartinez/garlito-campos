package domain

type Lista struct {
	id         int64
	nombre     string
	valores    []string
	repository ListaRepository
}

func NewLista(nombre string) *Lista {
	return &Lista{
		nombre: nombre,
	}
}

func (l *Lista) SetRepository(repository ListaRepository) {
	l.repository = repository
}

func (l *Lista) SetId(id int64) {
	l.id = id
}

func (l *Lista) Id() int64 {
	return l.id
}

func (l *Lista) SetNombre(nombre string) {
	l.nombre = nombre
}

func (l *Lista) Nombre() string {
	return l.nombre
}

func (l *Lista) SetValores(valores []string) {
	l.valores = valores
}

func (l *Lista) Valores() []string {
	return l.valores
}

func (l *Lista) Crear() error {
	return l.repository.CrearLista(*l)
}

func (l *Lista) Actualizar() error {
	return l.repository.ActualizarLista(*l)
}

func (l *Lista) Eliminar() error {
	return l.repository.EliminarLista(l.id)
}

func (l *Lista) Existe() bool {
	return l.id > 0
}
