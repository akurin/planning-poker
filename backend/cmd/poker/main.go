package main

import (
	"backend/internal/adapters/httpapi"
	"backend/internal/adapters/httpapi/gameapi"
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/usecase/findgameusecase"
	"log"
)

func main() {
	findGameUseCase := findgameusecase.New(inmemoryrepo.NewGameRepository())
	gameApi := gameapi.NewGameApi(findGameUseCase)
	serverFactory := httpapi.NewServerFactory(gameApi)
	server := serverFactory.NewServer()

	log.Printf("Starting to listen, addr: %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
