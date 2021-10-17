package domain

type fakeTokenService struct {
	issueToken func() (string, error)
}

func NewFakeTokenService(opts ...FakeTokenServiceOption) fakeTokenService {
	defaultIssueTokenFunc := func() (string, error) {
		return "some-token", nil
	}

	result := fakeTokenService{defaultIssueTokenFunc}
	for _, opt := range opts {
		opt(&result)
	}
	return result
}

func (s fakeTokenService) IssueToken(_ PlayerId) (string, error) {
	return s.issueToken()
}

type FakeTokenServiceOption func(service *fakeTokenService)

func WithToken(token string) FakeTokenServiceOption {
	return func(service *fakeTokenService) {
		service.issueToken = func() (string, error) {
			return token, nil
		}
	}
}

func WithError(err error) FakeTokenServiceOption {
	return func(service *fakeTokenService) {
		service.issueToken = func() (string, error) {
			return "", err
		}
	}
}
