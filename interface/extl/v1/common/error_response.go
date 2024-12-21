package common

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func NewGeneralBadRequest(err interface{}) *echo.HTTPError {
	return echo.NewHTTPError(
		http.StatusBadRequest,
		NewBadRequest(err, ""),
	)
}

func NewGeneralBadRequestWithMessage(err interface{}, message string) *echo.HTTPError {
	return echo.NewHTTPError(
		http.StatusBadRequest,
		NewBadRequest(err, message),
	)
}

func NewGeneralUnprocessableEntity(err interface{}) *echo.HTTPError {
	return echo.NewHTTPError(
		http.StatusUnprocessableEntity,
		NewUnprocessableEntity(err, ""),
	)
}

func NewGeneralUnprocessableEntityWithMessage(err interface{}, message string) *echo.HTTPError {
	return echo.NewHTTPError(
		http.StatusUnprocessableEntity,
		NewUnprocessableEntity(err, message),
	)
}

func NewGeneralInternalServerError(err interface{}) *echo.HTTPError {
	return echo.NewHTTPError(
		http.StatusInternalServerError,
		NewInternalServerError(err, ""),
	)
}

func NewGeneralInternalServerErrorWithMessage(err interface{}, message string) *echo.HTTPError {
	return echo.NewHTTPError(
		http.StatusInternalServerError,
		NewInternalServerError(err, message),
	)
}
