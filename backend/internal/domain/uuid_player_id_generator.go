package domain

import "github.com/google/uuid"

type uuidPlayerIdGenerator struct {
}

func NewUUIDPlayerIdGenerator() PlayerIdGenerator {
	return uuidPlayerIdGenerator{}
}

func (g uuidPlayerIdGenerator) Generate() PlayerId {
	return uuid.New()
}
