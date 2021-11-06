package main

import (
	"backend/internal/adapters/httpapi"
	"backend/internal/adapters/httpapi/gameapi"
	"backend/internal/adapters/httpapi/gamesapi"
	"backend/internal/adapters/httpapi/sessions"
	"backend/internal/adapters/httpapi/signupapi"
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/domain"
	"backend/internal/usecase/createplayerusecase"
	"backend/internal/usecase/findgameusecase"
	"backend/internal/usecase/startgameusecase"
	"log"
	"net/url"
)

func main() {
	// Repositories
	playerRepository := inmemoryrepo.NewPlayerRepository()
	gameRepository := inmemoryrepo.NewGameRepository()

	// Domain
	playerIdGenerator := domain.NewUUIDPlayerIdGenerator()

	// Use cases
	findGameUseCase := findgameusecase.New(inmemoryrepo.NewGameRepository())
	createPlayerUseCase := createplayerusecase.New(playerIdGenerator, playerRepository)
	startGameUseCase := startgameusecase.New(gameRepository)

	// Http
	baseUrl, _ := url.Parse("http://") // todo

	playersApi := signupapi.New(createPlayerUseCase, sessions.NewFakeStore()) // todo
	gamesApi := gamesapi.NewGamesApi(baseUrl, startGameUseCase)
	gameApi := gameapi.NewGameApi(findGameUseCase)

	serverFactory := httpapi.NewServerFactory(playersApi, gamesApi, gameApi)
	server := serverFactory.NewServer()

	log.Printf("Starting to listen, addr: %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
