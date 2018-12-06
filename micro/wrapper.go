package micro

import (
	"github.com/zm-dev/gerrors"
	"github.com/micro/go-micro/server"
	"context"
)

func NewHandlingErrorWrapper(serviceName string) server.HandlerWrapper {
	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			return handlingError(fn(ctx, req, rsp), serviceName)
		}
	}
}

func handlingError(err error, serviceName string) error {
	if err == nil {
		return nil
	}
	if ge, ok := err.(*gerrors.GlobalError); ok {
		if ge.ServiceName == "" {
			ge.ServiceName = serviceName
		}
		return err
	}
	return gerrors.InternalServerError(10000, serviceName, err.Error(), "")
}
