package concrete

import (
	logs "be_test_linkque/utils/middlewares/log"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogMiddleware struct {
	MiddlewareProto
}

func (hv *LogMiddleware) GetCallback() func(next echo.HandlerFunc) echo.HandlerFunc {
	return middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			form, err := c.MultipartForm()
			if err != nil {
				return false
			}

			if len(form.File["media"]) > 0 {
				return true
			}
			return false
		},
		Handler: func(ctx echo.Context, req, res []byte) {
			responseTime := time.Now()
			requestTime := time.Now()
			fields := []zapcore.Field{
				zap.String("unique_id", uuid.New().String()),
				zap.String("request", string(req)),
				zap.String("response", string(res)),
			}
			if req != nil {
				fields = append(fields, zap.String("request_time", requestTime.String()))
			}

			if res != nil {
				fields = append(fields, zap.String("response_time", responseTime.String()))
				processingTime := time.Since(requestTime)
				fields = append(fields, zap.Int("processing_time_nano_second", int(processingTime.Nanoseconds())))
			}
			logs.Logger.Info("log global request and response ", fields...)
		},
	})
}

func (hv *LogMiddleware) GetName() string {
	return "engine.middleware.core.LogMiddleware"
}

func (hv *LogMiddleware) IsPre() bool {
	return false
}
