package middleware

import (
	"bytes"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"sync"
)

var operationRecordService = service.ServiceGroupApp.SystemServiceGroup.OperationRecordService

var respPool sync.Pool
var bufferSize = 1024

func init() {
	respPool.New = func() interface{} {
		return make([]byte, bufferSize)
	}
}

func OperationRecord() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		//var body []byte
		//var userId int
		//if !c.IsGet() {
		//	var err error
		//	body, err = io.ReadAll(c.Request.Body())
		//	if err != nil {
		//		global.GVA_LOG.Error("read body from request error:", zap.Error(err))
		//	} else {
		//		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		//	}
		//} else {
		//	//query := c.Request.URL.RawQuery
		//	//TODO check
		//	query := c.QueryArgs().String()
		//	query, _ = url.QueryUnescape(query)
		//	split := strings.Split(query, "&")
		//	m := make(map[string]string)
		//	for _, v := range split {
		//		kv := strings.Split(v, "=")
		//		if len(kv) == 2 {
		//			m[kv[0]] = kv[1]
		//		}
		//	}
		//	body, _ = json.Marshal(&m)
		//}
		//claims, _ := utils.GetClaims(c)
		//if claims != nil && claims.BaseClaims.ID != 0 {
		//	userId = int(claims.BaseClaims.ID)
		//} else {
		//	id, err := strconv.Atoi(c.Request.Header.Get("x-user-id"))
		//	if err != nil {
		//		userId = 0
		//	}
		//	userId = id
		//}
		//record := system.SysOperationRecord{
		//	Ip:     c.ClientIP(),
		//	Method: c.Request.Method(),
		//	Path:   string(c.Request.Path()),
		//	Agent:  c.Request.UserAgent(),
		//	Body:   string(body),
		//	UserID: userId,
		//}
		//
		//// 上传文件时候 中间件日志进行裁断操作
		//if strings.Contains(string(c.GetHeader("Content-Type")), "multipart/form-data") {
		//	if len(record.Body) > bufferSize {
		//		record.Body = "[文件]"
		//	}
		//}
		//
		//writer := responseBodyWriter{
		//	ResponseWriter: c.Writer,
		//	body:           &bytes.Buffer{},
		//}
		//c.Writer = writer
		//now := time.Now()
		//
		//c.Next(ctx)
		//
		//latency := time.Since(now)
		//record.ErrorMessage = c.Errors.ByType(errors.ErrorTypePrivate).String()
		//record.Status = c.Writer.Status()
		//record.Latency = latency
		//record.Resp = writer.body.String()
		//
		//if strings.Contains(c.Writer.Header().Get("Pragma"), "public") ||
		//	strings.Contains(c.Writer.Header().Get("Expires"), "0") ||
		//	strings.Contains(c.Writer.Header().Get("Cache-Control"), "must-revalidate, post-check=0, pre-check=0") ||
		//	strings.Contains(c.Writer.Header().Get("Content-Type"), "application/force-download") ||
		//	strings.Contains(c.Writer.Header().Get("Content-Type"), "application/octet-stream") ||
		//	strings.Contains(c.Writer.Header().Get("Content-Type"), "application/vnd.ms-excel") ||
		//	strings.Contains(c.Writer.Header().Get("Content-Type"), "application/download") ||
		//	strings.Contains(c.Writer.Header().Get("Content-Disposition"), "attachment") ||
		//	strings.Contains(c.Writer.Header().Get("Content-Transfer-Encoding"), "binary") {
		//	if len(record.Resp) > bufferSize {
		//		// 截断
		//		record.Body = "超出记录长度"
		//	}
		//}
		//
		//if err := operationRecordService.CreateSysOperationRecord(record); err != nil {
		//	global.GVA_LOG.Error("create operation record error:", zap.Error(err))
		//}
	}
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
