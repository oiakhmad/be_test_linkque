package concrete

import (
	"github.com/labstack/echo/v4"
)

type MiddlewareProto interface {
	// get middleware callback
	GetCallback() func(next echo.HandlerFunc) echo.HandlerFunc

	// get associated middleware name
	GetName() string

	// is this middleware want to
	// run before routing process?
	IsPre() bool
}
