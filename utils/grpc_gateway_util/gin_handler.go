package grpc_gateway_util

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

var successRespBytes1 = []byte(`{"ret":1,"data":`)
var successRespBytes2 = []byte(`}`)

type customResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// Write 只写如缓存中,不写如流中
func (w *customResponseWriter) Write(b []byte) (int, error) {
	return w.body.Write(b)
}

// WriteString 只写如缓存中,不写如流中
func (w *customResponseWriter) WriteString(s string) (int, error) {
	return w.body.WriteString(s)
}

// Flush 把缓存写入到流中
func (w *customResponseWriter) Flush() {
	_, _ = w.ResponseWriter.Write(w.body.Bytes())
	w.ResponseWriter.Flush()
	return
}

// GinHandler 用作body流的复写操作.会自动把响应的数据放入到body中
// 如果存在 errHandlerFlagBit 的话会交给 ErrHandler 去写流
func GinHandler(c *gin.Context) {
	blw := &customResponseWriter{body: bytes.NewBuffer([]byte{}), ResponseWriter: c.Writer}
	c.Writer = blw

	c.Next()

	defer c.Writer.Flush()

	// resp 被err重写
	if c.Request.Header.Get(errHandlerFlagBit) == errHandlerFlagBitVal {
		return
	}

	buffer := bytes.NewBuffer([]byte{})
	buffer.Grow(len(successRespBytes1) + len(successRespBytes2) + blw.body.Len())
	buffer.Write(successRespBytes1)
	buffer.Write(blw.body.Bytes())
	buffer.Write(successRespBytes2)

	blw.body = buffer

	c.Writer.Header().Set("Content-Type", `application/json`)
}
