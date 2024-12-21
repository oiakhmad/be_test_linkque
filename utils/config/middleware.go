package config

import (
	"be_test_linkque/utils/middlewares"
	"be_test_linkque/utils/middlewares/concrete"
)

var (
	MiddlewareFactory = middlewares.NewMiddlewareFactory()
)

func RegisterRequiredMiddleware() {
	MiddlewareFactory.Register(&concrete.RemoveTrailingSlash{}, true)
	MiddlewareFactory.Register(&concrete.LogMiddleware{}, true)
}
