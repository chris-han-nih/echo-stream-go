package auth

import (
	"context"
	"github.com/echo-stream/api/infrastructure/repository"
	"github.com/echo-stream/database/ent"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
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
		log.Infof("%v", appMap)
		for k, v := range appMap {
			if k.String() == clientId {
				if v == clientSecret {
					c.Request().Header.Set("X-Application-Id", k.String())
					return next(c)
				}
			}
		}
		return echo.ErrUnauthorized
	}
}
