package data

import (
	"fmt"
	"log"
	"os"

	model "github.com/aleja/test-server/data/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//se utilizada la libreria para abrir postgres y se genere la conexion a la db
func Connect() *gorm.DB {
	//function to connect to the database with the information of the .env
	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)
	ConexionDd, err := gorm.Open(postgres.Open(pgConnString))
	if err != nil {
		log.Fatal(err)

	}
	//se utiliza el auto migrate para crear la tabla y las columnas correspondientes
	ConexionDd.AutoMigrate(&model.Superhero{})
	return ConexionDd
}
