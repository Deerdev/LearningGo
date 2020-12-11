package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/pkg/app"
	"github.com/go-programming-tour-book/blog-service/pkg/errcode"
	"github.com/go-programming-tour-book/blog-service/pkg/limiter"
)

// 接口限流
// ratelimit提供了一个简单又高效的令牌桶实现，可以帮助我们实现限流器的逻辑
/*
LimiterIface接口，用于定义当前限流器所必需的方法。

实际上限流器的形式有多种，可能某一类接口需要限流器 A，而另外一类接口需要限流器B，它们所采用的策略并不完全一致，因此我们需要声明LimiterIface这类通用接口，
保证其接口的设计。我们初步的在Iface接口中声明以下三个方法：

· Key：获取对应的限流器的键值对名称。
· GetBucket：获取令牌桶。
· AddBuckets：新增多个令牌桶。

定义Limiter 结构体，存储令牌桶与键值对名称的映射关系。定义LimiterBucketRule 结构体，存储令牌桶的一些相应规则属性，具体如下：

· Key：自定义键值对名称。
· FillInterval：间隔多久时间放N个令牌。
· Capacity：令牌桶的容量。
· Quantum：每次到达间隔时间后所放的具体令牌数量。

至此就完成了一个Limiter最基本的属性定义，接下来针对不同的情况，实现这个项目中的限流器。

(2)MethodLimiter.

前文编写的是一个简单的限流器，它的主要功能是对路由进行限流。因为在项目中，我们可能只需要对某一部分接口进行流量调控即可。打开pkg/limiter，
并新建method_limiter.go文件，写入如下代码：

在上述代码中，对LimiterIface接口实现了MethodLimiter限流器，主要逻辑是在Key方法中根据RequestURI切割出核心路由作为键值对名称，
并从GetBucket和AddBuckets中获取和设置Bucket的对应逻辑。
 */

// 带你快速了解：限流中的漏桶和令牌桶算法 https://mp.weixin.qq.com/s/4SzZEUTmjwAAsH5Qd2GWLQ
// 限流熔断是什么 https://mp.weixin.qq.com/s/aA0M6wL1mrAa0SZd8-Hv-w
func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
	return func(c *gin.Context) {
		/*
		在RateLimiter中间件中，需要注意的是入参应该为LimiterIface接口类型。这样一来，只要符合该接口类型的具体限流器实现都可以传入并使用。
		另外，TakeAvailable 方法会占用存储桶中立即可用的令牌的数量，返回值为删除的令牌数。如果没有可用的令牌，则返回 0，即已经超出配额了。
		这时将返回errcode.TooManyRequest状态，让客户端减缓请求速度。
		 */
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
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
