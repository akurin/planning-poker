package domain

type fakeTokenService struct {
	token string
}

func NewFakeTokenService(opts ...FakeTokenServiceOption) fakeTokenService {
	const defaultToken = "some"

	result := fakeTokenService{defaultToken}
	for _, opt := range opts {
		opt(&result)
	}
	return result
}

func (s fakeTokenService) IssueToken(_ PlayerId) string {
	return s.token
}

type FakeTokenServiceOption func(service *fakeTokenService)

func WithToken(token string) FakeTokenServiceOption {
	return func(service *fakeTokenService) {
		service.token = token
	}
}
