package concrete

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RemoveTrailingSlash struct {
	MiddlewareProto
}

func (rt *RemoveTrailingSlash) GetCallback() func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.RemoveTrailingSlash()
}

func (rt *RemoveTrailingSlash) GetName() string {
	return "engine.middleware.provider.RemoveTrailingSlash"
}

func (rt *RemoveTrailingSlash) IsPre() bool {
	return true
}
