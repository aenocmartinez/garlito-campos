package mock

import (
	"garlito/src/domain"
	"garlito/src/view/dto"
	"math/rand"
)

var (
	names        = []string{"Autor", "Título", "Técnica", "Dimensiones", "Estado conservación", "Número registro"}
	descriptions = []string{"Comentario 1", "Comentario 2", "Comentario 3", "Comentario 4", "Comentario 5"}
)

type CampoDao struct {
	data []domain.Simple
}

func NewCampoDao() *CampoDao {
	return &CampoDao{
		data: loaderData(),
	}
}

// func generateRandomName() string {
// 	name := names[rand.Intn(len(names))]
// 	return name
// }

func generateRandomDescription() string {
	description := descriptions[rand.Intn(len(descriptions))]
	return description
}

func loaderData() (data []domain.Simple) {
	for i := 1; i <= len(names); i++ {
		campo := domain.NewSimple()

		campo.SetId(int64(i))
		// campo.SetNombre(generateRandomName())
		campo.SetNombre(names[i-1])
		campo.SetDescripcion(generateRandomDescription())

		data = append(data, *campo)
	}
	return data
}

func (c *CampoDao) CrearCampo(campo domain.Campo) bool {
	id := len(c.data) + 1

	simple, _ := campo.(*domain.Simple)
	simple.SetId(int64(id))
	simple.SetNombre(campo.Nombre())
	simple.SetDescripcion(campo.Descripcion())

	c.data = append(c.data, *simple)
	return true
}

func (c *CampoDao) EliminarCampo(id int64) bool {
	data := []domain.Simple{}
	for _, item := range c.data {
		if item.Id() == id {
			continue
		}
		data = append(data, item)
	}

	c.data = data
	return true
}

func (c *CampoDao) ActualizarCampo(campo domain.Campo) bool {
	data := []domain.Simple{}
	for _, item := range c.data {
		if item.Id() == campo.Id() {
			item.SetNombre(campo.Nombre())
			item.SetDescripcion(campo.Descripcion())
		}
		data = append(data, item)
	}

	c.data = data
	return true
}

func (c *CampoDao) AgregarSuncampo(campo domain.Campo, subcampo domain.Campo) bool {
	return false
}

func (c *CampoDao) QuitarSuncampo(campo domain.Campo, subcampo domain.Campo) bool {
	return false
}

func (c *CampoDao) SubCampos(campo domain.Campo) (subcampos []domain.Campo) {
	return subcampos
}

func (c *CampoDao) ListarCampos() (campos []dto.CampoDto) {
	campos = []dto.CampoDto{}
	for _, item := range c.data {
		campos = append(campos, dto.CampoDto{
			Id:          item.Id(),
			Nombre:      item.Nombre(),
			Descripcion: item.Descripcion(),
		})
	}
	return campos
}

func (c *CampoDao) BuscarCampoPorId(id int64) (campo domain.Campo) {
	for _, item := range c.data {
		if item.Id() != id {
			continue
		}
		return &item
	}
	return &domain.Simple{}
}

func (c *CampoDao) BuscarCampo(criteria string) (campos []domain.Campo) {
	return campos
}

func (c *CampoDao) BuscarCampoPorNombre(nombre string) (campo domain.Campo) {
	return &domain.Simple{}
}
