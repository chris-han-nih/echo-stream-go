package usecase

import (
	"context"
	"github.com/echo-stream/api/domain/models"
	"github.com/echo-stream/api/infrastructure/repository"
)

type SlackUsecase struct {
	slackRepository repository.SlackRepository
}

func NewSlackUsecase(slackRepository repository.SlackRepository) *SlackUsecase {
	return &SlackUsecase{
		slackRepository: slackRepository,
	}
}

func (s *SlackUsecase) Create(ctx context.Context, m *models.Message) error {
	err := s.slackRepository.Create(ctx, m)
	return err
}
