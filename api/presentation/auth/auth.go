package middleware

import (
	"context"
	"github.com/echo-stream/api/infrastructure/repository"
	"github.com/echo-stream/database/ent"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"sync"
)

var once sync.Once

type Apps map[uuid.UUID]string

var appMap Apps

func InitAuth(ent *ent.Client) Apps {
	once.Do(func() {
		appMap, _ = repository.NewApplicationRepository(ent).FindAll(context.Background())
	})
	return appMap
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		clientId, clientSecret, ok := c.Request().BasicAuth()
		if !ok {
			return echo.ErrUnauthorized
		}
		for k, v := range appMap {
			if v == clientId {
				if k.String() == clientSecret {
					return next(c)
				}
			}
		}
		return echo.ErrUnauthorized
	}
}
