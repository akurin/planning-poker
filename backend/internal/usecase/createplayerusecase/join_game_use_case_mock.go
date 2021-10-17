package createplayerusecase

import (
	"backend/internal/domain"
)

type fakeUseCase struct {
	execute func() (Result, error)
}

func (e fakeUseCase) Execute(_ string) (Result, error) {
	return e.execute()
}

func Mock(opts ...FakeUseCaseOption) UseCase {
	defaultExecute := func() (Result, error) {
		return NewResult(domain.NewFakePlayerId("some"), "some"), nil
	}
	result := fakeUseCase{defaultExecute}
	for _, opt := range opts {
		opt(&result)
	}
	return result
}

type FakeUseCaseOption func(useCase *fakeUseCase)

func WithResult(result Result) FakeUseCaseOption {
	return func(useCase *fakeUseCase) {
		useCase.execute = func() (Result, error) {
			return result, nil
		}
	}
}

func WithError(err error) FakeUseCaseOption {
	return func(useCase *fakeUseCase) {
		useCase.execute = func() (Result, error) {
			return Result{}, err
		}
	}
}
