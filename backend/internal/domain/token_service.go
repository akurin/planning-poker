package domain

type TokenService interface {
	IssueToken(playerId PlayerId) (string, error)
}
