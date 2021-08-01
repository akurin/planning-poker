package domain

type PlayerIdGenerator interface {
	Generate() PlayerId
}
