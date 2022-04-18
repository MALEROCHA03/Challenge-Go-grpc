package main

import (
	"fmt"
	"log"
	"net"

	superherov1 "github.com/aleja/test-models/pkg/superhero/v1"
	superheroserver "github.com/aleja/test-server/internal/superhero"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
)

func main() {
	godotenv.Load()

	host := fmt.Sprintf("localhost:%d", 3000)
	lis, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	superheroServer := superheroserver.SuperheroServer{} //base de datas
	superherov1.RegisterSuperheroApiServer(grpcServer, &superheroServer)
	fmt.Println("server up on: ", host)
	grpcServer.Serve(lis)

}
