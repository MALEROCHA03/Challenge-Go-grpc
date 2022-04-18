package main

//se coloca pakage main ya que si no es asi a la hora de corre el o run si no es package principal no lo corre
import (
	"context"

	superherov1 "github.com/aleja/test-models/pkg/superhero/v1"
	"google.golang.org/grpc"
)

//creamos la funcion crear super heroe , esperando los datos que se requieren del Superhero
func CreateHero(hero *superherov1.Superhero) (bool, string) {
	if hero == nil {
		return false, "enviaste un valor nulo."
	}

	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		return false, err.Error()
	}

	client := superherov1.NewSuperheroApiClient(conn)

	response, err := client.Crear(context.Background(), &superherov1.Superhero{
		Nombre:    hero.Nombre,
		Habilidad: hero.Habilidad,
		Casa:      hero.Casa,
	},
	)

	if err != nil {
		return false, err.Error()
	}

	return response.Success, response.Error
}

func main() {
	hero := superherov1.Superhero{
		Nombre:    "aleja",
		Habilidad: "agua",
		Casa:      superherov1.Casa_dc,
	}
	CreateHero(&hero)
}
