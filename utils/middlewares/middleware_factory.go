package middlewares

import (
	"be_test_linkque/utils/middlewares/concrete"
	"fmt"
)

type MiddlewareFactoryProto interface {
	// register new middleware handler.
	Register(handler concrete.MiddlewareProto, replace bool) error

	// get registered middleware handler
	// by name.
	Get(name string) (concrete.MiddlewareProto, error)

	// get all registered middleware handler.
	GetAll() map[string]concrete.MiddlewareProto
}

type MiddlewareFactory struct {
	Handlers map[string]concrete.MiddlewareProto
}

func NewMiddlewareFactory() MiddlewareFactoryProto {
	return &MiddlewareFactory{
		Handlers: make(map[string]concrete.MiddlewareProto, 0),
	}
}

func (mf *MiddlewareFactory) Register(handler concrete.MiddlewareProto, replace bool) error {
	if _, ok := mf.Handlers[handler.GetName()]; ok {
		if !replace {
			return fmt.Errorf("Middleware handler with key '%s' exists.", handler.GetName())
		}

		mf.Handlers[handler.GetName()] = handler
	}

	mf.Handlers[handler.GetName()] = handler
	return nil
}

func (mf *MiddlewareFactory) Get(name string) (concrete.MiddlewareProto, error) {
	handler, ok := mf.Handlers[name]

	if !ok {
		return nil, fmt.Errorf("Middleware handler with key '%s' not found.", name)
	}

	return handler, nil
}

func (mf *MiddlewareFactory) GetAll() map[string]concrete.MiddlewareProto {
	return mf.Handlers
}
