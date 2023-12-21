package domain

type Simple struct {
	id          int64
	nombre      string
	descripcion string
	repository  CampoRepository
}

func NewSimple() *Simple {
	return &Simple{}
}

func (s *Simple) SetRepository(repository CampoRepository) {
	s.repository = repository
}

func (s *Simple) SetId(id int64) {
	s.id = id
}

func (s *Simple) Id() int64 {
	return s.id
}

func (s *Simple) SetNombre(nombre string) {
	s.nombre = nombre
}

func (s *Simple) Nombre() string {
	return s.nombre
}

func (s *Simple) SetDescripcion(descripcion string) {
	s.descripcion = descripcion
}

func (s *Simple) Descripcion() string {
	return s.descripcion
}

func (s *Simple) EsCompuesto() bool {
	return false
}

func (s *Simple) Crear() bool {
	return s.repository.CrearCampo(s)
}

func (s *Simple) Eliminar() bool {
	return s.repository.EliminarCampo(s.id)
}

func (s *Simple) Actualizar() bool {
	return s.repository.ActualizarCampo(s)
}

func (s *Simple) Existe() bool {
	return s.id > 0
}
