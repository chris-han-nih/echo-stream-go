package controller

import (
	"github.com/echo-stream/api/application/usecase"
	"github.com/echo-stream/api/domain/models"
	"github.com/labstack/echo/v4"
)

type SlackController struct {
	SlackUsecase usecase.SlackUsecase
}

func (s *SlackController) Create(e echo.Context) error {
	slackMessage := new(models.Message)
	if err := e.Bind(slackMessage); err != nil {
		return e.JSON(400, err)
	}
	if err := e.Validate(slackMessage); err != nil {
		return err
	}
	err := s.SlackUsecase.Create(e.Request().Context(), slackMessage)
	if err != nil {
		return e.JSON(500, err)
	}
	return e.JSON(200, slackMessage)
}
