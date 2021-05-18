/*
 * @Author: Kim
 * @Date: 2021-05-18 17:00:41
 * @LastEditTime: 2021-05-18 18:48:03
 * @LastEditors: Kim
 * @Description: v
 * @FilePath: /template_go/internal/middleware/tracer.go
 */
package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
