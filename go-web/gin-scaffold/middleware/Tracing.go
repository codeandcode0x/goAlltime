package middleware

import (
	"log"

	tracing "github.com/codeandcode0x/traceandtrace-go"
	"github.com/gin-gonic/gin"
)

// Tracing 中间件
func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("..... tracing header ", c.Request.Header)
		// add tracing
		pctx, cancel := tracing.AddHttpTracing(
			"ticket-manager",
			c.Request.URL.String()+" "+c.Request.Method,
			c.Request.Header,
			map[string]string{
				"component":      "gin-server",
				"httpMethod":     c.Request.Method,
				"httpUrl":        c.Request.URL.String(),
				"proto":          c.Request.Proto,
				"peerService":    c.HandlerName(),
				"httpStatusCode": "200",
				"spanKind":       "server",
			})
		defer cancel()

		// deliver parent context
		c.Set("parentCtx", pctx)
		c.Next()
	}
}
