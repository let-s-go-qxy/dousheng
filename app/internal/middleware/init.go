package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

// InitMiddleWareForDefault 此处配置全局的middleware
func InitMiddleWareForDefault(h *server.Hertz) *server.Hertz {
	h.Use(Cors())
	return h
}
