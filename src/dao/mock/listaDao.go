package mock

import (
	"garlito/src/domain"
	"garlito/src/view/dto"
	"strings"
)

var (
	listas []string = []string{"denominaciones", "días de la semana", "nombre de los meses"}
)

type ListaDao struct {
	data []domain.Lista
}

func NewListaDao() *ListaDao {
	return &ListaDao{
		data: cargarLista(),
	}
}

func cargarLista() (data []domain.Lista) {
	for id, nombre := range listas {
		lista := domain.NewLista(nombre)
		lista.SetId(int64(id + 1))

		data = append(data, *lista)
	}
	return data
}

func (l *ListaDao) CrearLista(lista domain.Lista) (err error) {
	nuevaLista := domain.NewLista(lista.Nombre())
	nuevaLista.SetId(int64(len(l.data) + 1))
	nuevaLista.SetValores(lista.Valores())
	l.data = append(l.data, *nuevaLista)
	return err
}

func (l *ListaDao) ActualizarLista(lista domain.Lista) (err error) {
	var data []domain.Lista = []domain.Lista{}
	for _, item := range l.data {
		if item.Id() == lista.Id() {
			item.SetNombre(lista.Nombre())
			item.SetValores(lista.Valores())
		}

		data = append(data, item)
	}

	l.data = data

	return err
}

func (l *ListaDao) EliminarLista(id int64) (err error) {
	var data []domain.Lista = []domain.Lista{}
	for _, item := range l.data {
		if item.Id() == id {
			continue
		}

		data = append(data, item)
	}

	l.data = data

	return err
}

func (l *ListaDao) ListarLista() (lista []dto.ListaDto) {
	for _, item := range l.data {
		lista = append(lista, dto.ListaDto{
			Id:     item.Id(),
			Nombre: item.Nombre(),
		})
	}
	return lista
}

func (l *ListaDao) BuscarListaPorId(id int64) (lista domain.Lista) {
	for _, item := range l.data {
		if item.Id() == id {
			return item
		}
	}
	return lista
}

func (l *ListaDao) BuscarListaPorNombre(nombre string) (lista domain.Lista) {
	for _, item := range l.data {
		if strings.EqualFold(item.Nombre(), nombre) {
			return item
		}
	}
	return lista
}
