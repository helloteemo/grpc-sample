package log

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/helloteemo/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func InitLog() {
	logrus.AddHook(&TraceIdHook{})
}

type TraceIdHook struct {
}

func (m *TraceIdHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

const TraceIdKey = "__trace_id"

func (m *TraceIdHook) Fire(entry *logrus.Entry) (err error) {
	var tradeId string
	if entry.Context != nil {
		var ok bool
		tradeId, ok = entry.Context.Value(TraceIdKey).(string)
		if !ok {
			return
		}
	}

	if tradeId == "" {
		return
	}

	entry.Data["trace_id"] = tradeId

	return
}

// GinTraceIdMiddleware 生成trade_id
func GinTraceIdMiddleware(ginCtx *gin.Context) {
	reqCtx := ginCtx.Request.Context()
	reqCtx = ContextWithTraceId(reqCtx)
	ginCtx.Request = ginCtx.Request.WithContext(reqCtx)
}

func ContextWithTraceId(ctx context.Context) (newCtx context.Context) {
	traceId, ok := ctx.Value(TraceIdKey).(string)
	if ok && traceId != "" {
		return
	}

	traceId = utils.UUID()
	newCtx = context.WithValue(ctx, TraceIdKey, traceId)
	return
}

func BgContextWithCancel() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(context.Background())
	ctx = ContextWithTraceId(ctx)
	return
}

// BgContextWithTraceId 使用一个traceId来创建
func BgContextWithTraceId(traceId string) (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(context.Background())
	ctx = context.WithValue(ctx, TraceIdKey, traceId)
	return
}

func GetTraceId(ctx context.Context) (traceId string) {
	traceId, _ = ctx.Value(TraceIdKey).(string)
	return
}

func GPPCClientContext(ctx context.Context) (newCtx context.Context) {
	outgoingContext := metadata.NewOutgoingContext(ctx, metadata.Pairs(TraceIdKey, GetTraceId(ctx)))
	return outgoingContext
}
