package repository

import (
	"context"
	"github.com/echo-stream/database/ent"
	"github.com/echo-stream/database/ent/application"
	"github.com/google/uuid"
)

type ApplicationRepository struct {
	client *ent.Client
}

func NewApplicationRepository(client *ent.Client) *ApplicationRepository {
	return &ApplicationRepository{
		client: client,
	}
}

type App struct {
	ApplicationId uuid.UUID `json:"application_id"`
	Secret        string    `json:"secret"`
}

func (a *ApplicationRepository) FindAll(ctx context.Context) (map[uuid.UUID]string, error) {
	var apps []App
	err := a.client.Application.
		Query().
		Select(application.FieldApplicationId, application.FieldSecret).
		Scan(ctx, &apps)
	if err != nil {
		return nil, err
	}
	var appMap = make(map[uuid.UUID]string)
	for _, app := range apps {
		appMap[app.ApplicationId] = app.Secret
	}
	return appMap, nil
}
