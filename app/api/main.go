package main

import (
	v1 "be_test_linkque/interface/extl/v1/routes"
	"be_test_linkque/utils/config"
	"be_test_linkque/utils/middlewares/log"
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4/v2"
)

func main() {
	config.LoadEnvVars()
	config.RegisterRequiredMiddleware()
	config.OpenMySQLPool()

	log.LoadLogger()
	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAuthorization,
		},
	}))

	// Handler for putting app request and response timestamp. This used for get elapsed time
	e.Use(ServiceRequestTime)

	dispatchMiddleware(e)
	v1.API(e)
	e.GET("/", func(c echo.Context) error {
		if err := c.Redirect(http.StatusFound, "/api/v1/health"); err != nil {
			return err
		}
		return nil
	})

	// Start server
	go func() {
		if err := e.Start(":" + os.Getenv("APP_PORT")); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}

// ServiceRequestTime middleware adds a `Server` header to the response.
func ServiceRequestTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Request().Header.Set("X-App-RequestTime", time.Now().Format(time.RFC3339))
		// lset log here
		return next(c)
	}
}

func dispatchMiddleware(e *echo.Echo) {
	middlewares := config.MiddlewareFactory.GetAll()
	for _, mf := range middlewares {
		if mf.IsPre() {
			e.Pre(mf.GetCallback())
		} else {
			e.Use(mf.GetCallback())
		}
	}
}
