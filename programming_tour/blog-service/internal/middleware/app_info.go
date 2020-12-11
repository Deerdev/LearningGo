package middleware

import "github.com/gin-gonic/gin"

// 我们经常需要在进程内上下文设置一些内部信息，既可以是应用名称和应用版本号这类基本信息，也可以是业务属性信息。
// 例如，想要根据不同的租户号获取不同的数据库实例对象，这时就需要在一个统一的地方进行处理
func AppInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app_name", "blog-service")
		c.Set("app_version", "1.0.0")
		c.Next()
	}
}
