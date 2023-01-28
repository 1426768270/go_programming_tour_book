package middleware

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"blog-service/pkg/limiter"
	"github.com/gin-gonic/gin"
)

// 限流
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc{
	return func(c *gin.Context){
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			// TakeAvailable 方法，它会占用存储桶中立即可用的令牌的数量，返回值为删除的令牌数，如果没有可用的令牌，将会返回 0，也就是已经超出配额了
			count := bucket.TakeAvailable(1)
			if count == 0 {
				response := app.NewResponse(c)
				response.ToErrorResponse(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
