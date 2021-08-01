package domain

type fakePlayerId string

func NewFakePlayerId(value string) PlayerId {
	return fakePlayerId(value)
}

func (i fakePlayerId) String() string {
	return string(i)
}
