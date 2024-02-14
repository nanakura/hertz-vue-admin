package middleware

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//defer func() {
		//	if err := recover(); err != nil {
		//		// Check for a broken connection, as it is not really a
		//		// condition that warrants a panic stack trace.
		//		var brokenPipe bool
		//		if ne, ok := err.(*net.OpError); ok {
		//			if se, ok := ne.Err.(*os.SyscallError); ok {
		//				if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
		//					brokenPipe = true
		//				}
		//			}
		//		}
		//
		//		httpRequest, _ := httputil.DumpRequest(c.Request, false)
		//		if brokenPipe {
		//			global.GVA_LOG.Error(string(c.Request.Path()),
		//				zap.Any("error", err),
		//				zap.String("request", string(httpRequest)),
		//			)
		//			// If the connection is dead, we can't write a status to it.
		//			_ = c.Error(err.(error)) // nolint: errcheck
		//			c.Abort()
		//			return
		//		}
		//
		//		if stack {
		//			global.GVA_LOG.Error("[Recovery from panic]",
		//				zap.Any("error", err),
		//				zap.String("request", string(httpRequest)),
		//				zap.String("stack", string(debug.Stack())),
		//			)
		//		} else {
		//			global.GVA_LOG.Error("[Recovery from panic]",
		//				zap.Any("error", err),
		//				zap.String("request", string(httpRequest)),
		//			)
		//		}
		//		c.AbortWithStatus(http.StatusInternalServerError)
		//	}
		//}()
		//c.Next(ctx)
	}
}
