package middleware

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
)

/*
在应用程序的运行过程中，经常会遇到一个让人头疼的问题，即假设应用A调用应用B，应用B调用应用C，如果应用C出现问题，
则在没有任何约束的情况下仍持续调用，就会导致应用A、B、C均出现问题。这就是十分常见的上下游应用的相互影响所导致的连环反应，最终使得整个集群应用出现一定规模的不可用

为了避免出现这种情况，最简单的一个约束点，就是统一在应用程序中针对所有请求都进行一个最基本的超时时间控制
 */

func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 调用了context.WithTimeout方法来设置当前context的超时时间，并重新赋给gin.Context。当当前请求运行到指定的时间后，
		// 使用了该context的运行流程就会对context提供的超时时间进行处理，并在指定的时间内取消请求
		// 需要将设置了超时的 c.Request.Context 方法传递进去，在验证时可以调短默认超时时间来进行调试
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
/*
需要注意的是，如果在进行多应用/服务的调用时，把父级的上下文信息（ctx）不断地传递下去，那么在统计超时控制的中间件中所设置的超时时间，
其实是针对整条链路的。如果需要单独调整某条链路的超时时间，那么只需调用context.WithTimeout等方法对父级 ctx 进行设置，然后取得子级 ctx，再进行新的传递即可
 */
