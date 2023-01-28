package limiter

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// 限流器所需要方法
type LimiterIface interface {
	Key(c *gin.Context) string
	GetBucket(key string) (*ratelimit.Bucket, bool)
	AddBuckets(rules ...LimiterBucketRule) LimiterIface
}

// 存储令牌桶与键值对名称映射关系
type Limiter struct {
	LimiterBuckets map[string]*ratelimit.Bucket
}

// 存储令牌桶的一些相应规则属性
type LimiterBucketRule struct {
	Key          string        //自定义键值对名称。
	FillInterval time.Duration //间隔多久时间放 N 个令牌。
	Capacity     int64         // 令牌桶的容量。
	Quantum      int64         //每次到达间隔时间后所放的具体令牌数量。
}

