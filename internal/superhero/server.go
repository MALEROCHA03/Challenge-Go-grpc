package superheroserver

import (
	"context"
	"log"

	superherov1 "github.com/aleja/test-models/pkg/superhero/v1"
	c "github.com/aleja/test-server/data"
	model "github.com/aleja/test-server/data/models"
)

type SuperheroServer struct {
	superherov1.SuperheroApiServer
}

func (server *SuperheroServer) Crear(ctx context.Context, request *superherov1.Superhero) (*superherov1.Result, error) {
	connection := c.Connect()
	response := &superherov1.Result{}
	hero := model.Superhero{
		Nombre:    request.Nombre,
		Habilidad: request.Habilidad,
		Casa:      request.Casa.Enum().String(),
		Atributos: model.Atributos{

			Debilidad: request.Atributos.Debilidad,
			Ataque:    request.Atributos.Ataque,
			Vida:      request.Atributos.Vida,
			Genero:    request.Atributos.Genero.Enum().String(),
		},
	}
	//verificamos si esta pasando los datos correctamente o es un error
	transaction := connection.Save(&hero)
	if transaction == nil {
		response.Error = "errorrrrrrr"
		return response, transaction.Error
	}
	response.Success = true
	log.Println("transaction", response)
	return response, nil

}

func (server *SuperheroServer) Actulizar(ctx context.Context, request *superherov1.Superhero) (*superherov1.Result, error) {

	return &superherov1.Result{}, nil
}
func (server *SuperheroServer) Lista(ctx context.Context, request *superherov1.Pagina) (*superherov1.ListaSuperhero, error) {

	return &superherov1.ListaSuperhero{}, nil
}

func (server *SuperheroServer) Eliminar(ctx context.Context, request *superherov1.Pagina) (*superherov1.ListaSuperhero, error) {

	return &superherov1.ListaSuperhero{}, nil
}
