package api

import (
	"errors"
	"go-lang-app/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authroizationHeader     = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authroizationHeader := c.Request.Header.Get(authroizationHeader)
		if len(authroizationHeader) == 0 {
			err := errors.New("missing authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		fields := strings.Fields(authroizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := errors.New("invalid authorization type")
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse(err))
			return
		}
		c.Set(authorizationPayloadKey, payload)
		c.Next()
	}
}
