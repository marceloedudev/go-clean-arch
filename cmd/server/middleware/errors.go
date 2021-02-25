package middleware

import (
	"fmt"
	"go-clean-arch/pkg/http_errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HandleErrors middleware
func HandleErrors() gin.HandlerFunc {
	return func(c *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				switch err.(type) {
				case *http_errors.HttpException:
					{
						errors := err.(*http_errors.HttpException)
						errors.Error = http.StatusText(errors.Status)
						errors.Timestamp = time.Now()
						errors.Path = c.Request.URL.Path
						c.JSON(errors.Status, errors)
						return
					}
				default:
					{
						c.JSON(http.StatusInternalServerError, &http_errors.HttpException{
							Message:   fmt.Sprintf("%v", err),
							Status:    http.StatusInternalServerError,
							Error:     http.StatusText(http.StatusInternalServerError),
							Timestamp: time.Now(),
							Path:      c.Request.URL.Path,
						})
						return
					}
				}
			}
		}()

		c.Next()
	}
}
