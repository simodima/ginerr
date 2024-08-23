package ginerr

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func translate(code int, t map[int]int) int {
	if statusCode, ok := t[code]; ok {
		return statusCode
	}

	return http.StatusInternalServerError
}

// AutoResponse automatically sends an HTTP response if there is an error
// in the gin.Context.
// translationMap is a translator to map internal error codes to HTTP codes
//
// Example:
//
//	AutoResponse(map[int]int{
//	  5001: 500,
//	  5002: 404,
//	  5003: 401,
//	})
func AutoResponse(translationMap map[int]int) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		code := -1
		// details := map[string]any{}
		translatableErrFound := false
		// var lastError error

		for _, err := range c.Errors {
			// lastError = err.Err

			// extracting the error code
			if codeErr, ok := err.Err.(interface {
				Error() string
				Code() int
			}); ok {
				code = codeErr.Code()
				translatableErrFound = true
			}

			if translatableErrFound {
				c.AbortWithStatusJSON(translate(code, translationMap), map[string]string{
					"code":    fmt.Sprint(code),
					"message": err.Error(),
				})
				break
			}
		}

		if len(c.Errors) > 0 {
			c.AbortWithStatusJSON(translate(code, translationMap), map[string]string{
				"code":    fmt.Sprint(code),
				"message": c.Errors.String(),
			})
		}
	}
}
