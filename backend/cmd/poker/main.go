package main

import (
	"backend/internal/adapters/httpapi"
	"backend/internal/adapters/httpapi/gameapi"
	"backend/internal/adapters/httpapi/gamesapi"
	"backend/internal/adapters/httpapi/playersapi"
	"backend/internal/adapters/inmemoryrepo"
	"backend/internal/adapters/systemclock"
	"backend/internal/domain"
	"backend/internal/usecase/createplayerusecase"
	"backend/internal/usecase/findgameusecase"
	"backend/internal/usecase/startgameusecase"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"log"
	"net/url"
	"time"
)

func main() {
	// Configuration
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	tokenTTL := time.Hour * 24

	// Clock
	clock := systemclock.New()

	// Repositories
	playerRepository := inmemoryrepo.NewPlayerRepository()
	gameRepository := inmemoryrepo.NewGameRepository()

	// Domain
	playerIdGenerator := domain.NewUUIDPlayerIdGenerator()
	tokenService := domain.NewJwtTokenService(privateKey, clock, tokenTTL)

	// Use cases
	findGameUseCase := findgameusecase.New(inmemoryrepo.NewGameRepository())
	createPlayerUseCase := createplayerusecase.New(playerIdGenerator, playerRepository, tokenService)
	startGameUseCase := startgameusecase.New(gameRepository)

	// Http
	baseUrl, _ := url.Parse("http://") // todo

	playersApi := playersapi.New(baseUrl, createPlayerUseCase)
	gamesApi := gamesapi.NewGamesApi(baseUrl, startGameUseCase)
	gameApi := gameapi.NewGameApi(findGameUseCase)

	serverFactory := httpapi.NewServerFactory(playersApi, gamesApi, gameApi)
	server := serverFactory.NewServer()

	log.Printf("Starting to listen, addr: %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
