package domain

type fakePlayerIdGenerator struct {
	value PlayerId
}

func NewFakePlayerIdGenerator(value PlayerId) PlayerIdGenerator {
	return fakePlayerIdGenerator{
		value: value,
	}
}

func (g fakePlayerIdGenerator) Generate() PlayerId {
	return g.value
}
