package middleware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	d2dContext "github.com/scalent-io/orchestration-framework/pkg/context"
	"github.com/scalent-io/orchestration-framework/pkg/log"
)

type MiddlewareConfig struct {
	AuthServiceEndpoint string
}

type Middleware interface {
	Access() gin.HandlerFunc
	Cors() gin.HandlerFunc
}

type MiddlewareImpl struct {
}

func NewMiddlewareImpl() (*MiddlewareImpl, error) {

	return &MiddlewareImpl{}, nil
}

func (m *MiddlewareImpl) Access() gin.HandlerFunc {
	return func(c *gin.Context) {

		log.Info("internal>middleware: access middleware started", EMPTY_STRING)

		c.Writer.Header().Set("Content-Type", "application/json")
		ctx := c.Request.Context()
		var reqId string
		reqId = c.Request.Header.Get(d2dContext.REQUEST_ID)

		if reqId == EMPTY_STRING {
			log.Info("internal>middleware: RequestID", EMPTY_STRING)
			// generate request id
			reqId = uuid.New().String()
			ctx = context.WithValue(ctx, d2dContext.ContextKey(d2dContext.REQUEST_ID), reqId)
			// send the request ID in response header
			c.Writer.Header().Add("requestID", reqId)
		} else {
			// create context with values
			ctx = context.WithValue(c, d2dContext.ContextKey(d2dContext.REQUEST_ID), reqId)
			ctx = context.WithValue(ctx, d2dContext.ContextKey(d2dContext.REQUEST_ID), reqId)
			// send the request ID in response header
			c.Writer.Header().Add("requestID", reqId)
		}

		log.Info("internal>middleware: auth middleware started", reqId)
		log.Info("internal>middleware: RequestID", c.Request.URL.Path)

		token := c.Request.Header.Get(d2dContext.TOKEN)

		if len(token) == 0 {
			cookieData, err := c.Request.Cookie("token")
			if err != nil {
				log.Error("create guest session", reqId)
			} else {
				token = cookieData.Value
			}
		}

		session := ""

		// add the token and the session data
		ctxWithAuthData := context.WithValue(ctx, d2dContext.ContextKey(d2dContext.TOKEN), token)
		ctxWithAuthData = context.WithValue(ctxWithAuthData, d2dContext.ContextKey(d2dContext.SESSION_DATA), session)

		c.Request = c.Request.WithContext(ctxWithAuthData)

		// serve the request to the next handler
		c.Next()
	}
}

func (m *MiddlewareImpl) Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
