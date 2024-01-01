package usecase

import (
	"context"
	"github.com/echo-stream/api/infrastructure/repository"
	"github.com/google/uuid"
)

type AuthUsecase struct {
	authRepository repository.ApplicationRepository
}

func NewAuthUsecase(authRepository repository.ApplicationRepository) *AuthUsecase {
	return &AuthUsecase{
		authRepository: authRepository,
	}
}

func (a *AuthUsecase) FindAll(ctx context.Context) (map[uuid.UUID]string, error) {
	appMap, err := a.authRepository.FindAll(ctx)
	return appMap, err
}
