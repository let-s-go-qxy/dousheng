package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func DemoMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// pre-handle
		// ...
		c.Next(ctx)
	}
}
