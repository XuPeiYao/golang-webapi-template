package common

import (
	"github.com/gofiber/fiber/v2"
)

type LogObjectConstructor func() any
type CtxRequestResponseCallbackFunc func(ctx *fiber.Ctx, data map[string]any) error
type CtxRequestCallbackFunc CtxRequestResponseCallbackFunc
type CtxResponseCallbackFunc CtxRequestResponseCallbackFunc

type FiberRequestResponseHookMiddleware struct {
	RequestFunc  CtxRequestCallbackFunc
	ResponseFunc CtxResponseCallbackFunc
}

func NewFiberRequestResponseHookMiddleware(
	requestFunc CtxRequestCallbackFunc,
	responseFunc CtxResponseCallbackFunc,
) *FiberRequestResponseHookMiddleware {
	return &FiberRequestResponseHookMiddleware{
		RequestFunc:  requestFunc,
		ResponseFunc: responseFunc,
	}
}

func (this *FiberRequestResponseHookMiddleware) Handler() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		data := make(map[string]any)

		err := this.RequestFunc(ctx, data)

		if err != nil {
			return err
		}

		err = ctx.Next()

		if err != nil {
			return err
		}

		err = this.ResponseFunc(ctx, data)

		return err
	}
}
