package main

import (
	"backend/internal/adapters/httpapi"
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/usecase/findgame"
	"log"
)

func main() {
	findGameUseCase := findgame.New(inmemoryrepo.NewGameRepository())
	gameApi := httpapi.NewGameApi(findGameUseCase)
	serverFactory := httpapi.NewServerFactory(gameApi)
	server := serverFactory.NewServer()

	log.Printf("Starting to listen, addr: %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
