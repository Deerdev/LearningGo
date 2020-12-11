package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/logger"
)

/*
从功能上讲，它会记录每一次请求的请求方法、方法调用开始时间、方法调用结束时间、方法响应结果和方法响应结果状态码。
除此之外，它还会记录RequestId、TraceId、SpanId等附加属性，以达到日志链路追踪的效果


 */
type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 遇到一个问题，即无法直接获取方法返回的响应主体，这时需要巧妙利用Go interface的特性。实际上在写入流时，调用的是http.ResponseWriter
// 在 AccessLogWriter 的 Write 方法中实现了双写，因此可以直接通过 AccessLogWriter的body取到值
func (w AccessLogWriter) Write(p []byte) (int, error) {
	// 双写
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

/*
在AccessLog方法中，我们初始化了AccessLogWriter，将其赋予当前的Writer写入流（可理解为替换原有），并且通过指定方法得到所需的日志属性，最终写到日志中，其中涉及的信息如下：

· method：当前的调用方法。
· request：当前的请求参数。
· response：当前的请求结果响应主体。
· status_code：当前的响应结果状态码。
· begin_time/end_time：调用方法的开始时间、调用方法的结束时间。
*/
func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyWriter

		beginTime := time.Now().Unix()
		c.Next()
		endTime := time.Now().Unix()

		fields := logger.Fields{
			"request":  c.Request.PostForm.Encode(),
			"response": bodyWriter.body.String(),
		}
		s := "access log: method: %s, status_code: %d, " +
			"begin_time: %d, end_time: %d"
		global.Logger.WithFields(fields).Infof(c, s,
			c.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
