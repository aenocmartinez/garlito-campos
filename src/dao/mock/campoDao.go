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
	campos []domain.Campo
}

func NewCampoDao() *CampoDao {
	return &CampoDao{
		campos: loaderData(),
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

func loaderData() (data []domain.Campo) {
	for i := 1; i <= len(names); i++ {

		campo := domain.NewSimple()
		campo.SetId(int64(i))
		// campo.SetNombre(generateRandomName())
		campo.SetNombre(names[i-1])
		campo.SetDescripcion(generateRandomDescription())

		data = append(data, campo)
	}
	return data
}

func (c *CampoDao) CrearCampo(campo domain.Campo) (err error) {
	id := len(c.campos) + 1
	campo.SetId(int64(id))

	c.campos = append(c.campos, campo)

	return err
}

func (c *CampoDao) EliminarCampo(id int64) (err error) {
	var campos []domain.Campo = []domain.Campo{}
	for _, campo := range c.campos {
		if campo.Id() == id {
			continue
		}
		campos = append(campos, campo)
	}

	c.campos = campos

	return err
}

func (c *CampoDao) ActualizarCampo(campo domain.Campo) (err error) {
	var campos []domain.Campo = []domain.Campo{}
	for _, item := range c.campos {
		if item.Id() == campo.Id() {
			item.SetNombre(campo.Nombre())
			item.SetDescripcion(campo.Descripcion())
		}
		campos = append(campos, item)
	}

	c.campos = campos

	return err
}

func (c *CampoDao) SubCampos(campo domain.Campo) (subcampos []domain.Campo) {
	return subcampos
}

func (c *CampoDao) ListarCampos() (campos []dto.CampoDto) {
	campos = []dto.CampoDto{}
	for _, item := range c.campos {
		campos = append(campos, dto.CampoDto{
			Id:          item.Id(),
			Nombre:      item.Nombre(),
			Descripcion: item.Descripcion(),
			EsCompuesto: item.EsCompuesto(),
		})
	}
	return campos
}

func (c *CampoDao) BuscarCampoPorId(id int64) (campo domain.Campo) {
	for _, campo := range c.campos {
		if campo.Id() != id {
			continue
		}
		return campo
	}
	return &domain.Simple{}
}

func (c *CampoDao) BuscarCampo(criteria string) (campos []domain.Campo) {
	return campos
}

func (c *CampoDao) BuscarCampoPorNombre(nombre string) (campo domain.Campo) {
	for _, campo := range c.campos {
		if campo.Nombre() != nombre {
			continue
		}
		return campo
	}
	return &domain.Simple{}
}
