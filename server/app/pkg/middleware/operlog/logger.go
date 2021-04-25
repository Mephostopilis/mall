package operlog

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	transportgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transporthttp "github.com/go-kratos/kratos/v2/transport/http"
)

// LoggerToFile 日志记录到文件
func LoggerToFile(logger log.Logger) middleware.Middleware {
	log := log.NewHelper("middleware/operlog", logger)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			tr, ok := transport.FromContext(ctx)
			if ok {
				log.Info("transport: %+v", tr)
			}
			h, ok := transporthttp.FromServerContext(ctx)
			if ok {
				log.Info("http: [%s] %s", h.Request.Method, h.Request.URL.Path)
				// 开始时间
				startTime := time.Now()

				// 结束时间
				endTime := time.Now()

				// 执行时间
				latencyTime := endTime.Sub(startTime)

				// 请求方式
				reqMethod := h.Request.Method

				// 请求路由
				reqUri := h.Request.RequestURI

				// 状态码
				statusCode := h.Request.Response.StatusCode

				// 请求IP
				clientIP := h.Request.RemoteAddr

				// 日志格式
				log.Info("%s %s %3d %13v %15s",
					reqMethod,
					reqUri,
					statusCode,
					latencyTime,
					clientIP,
				)

				if h.Request.Method != "GET" && h.Request.Method != "OPTIONS" {
					SetDBOperLog(ctx, h, h.Request.Method, clientIP, statusCode, reqUri, reqMethod, latencyTime, h.Request.UserAgent())
				}
			}
			g, ok := transportgrpc.FromServerContext(ctx)
			if ok {
				log.Info("grpc: %s", g.FullMethod)
			}
			return handler(ctx, req)
		}
	}
}

// 写入操作日志表
// 该方法后续即将弃用
func SetDBOperLog(c context.Context, svr transporthttp.ServerInfo, requestMethod, clientIP string, statusCode int, reqUri string, reqMethod string, latencyTime time.Duration, userAgent string) {
	// menu := model.Menu{}
	// menu.Path = reqUri
	// menu.Action = reqMethod
	// menuList, _ := global.AdminDao.GetMenus(&menu)
	// sysOperLog := model.SysOperLog{}
	// sysOperLog.OperIp = clientIP
	// sysOperLog.OperLocation = ip.GetLocation(clientIP)
	// sysOperLog.Status = xtools.IntToString(statusCode)
	// sysOperLog.OperName = jwtauth.GetUserName(c)
	// sysOperLog.RequestMethod = requestMethod
	// sysOperLog.OperUrl = reqUri
	// if reqUri == "/login" {
	// 	sysOperLog.BusinessType = "10"
	// 	sysOperLog.Title = "用户登录"
	// 	sysOperLog.OperName = "-"
	// } else if strings.Contains(reqUri, "/api/v1/logout") {
	// 	sysOperLog.BusinessType = "11"
	// } else if strings.Contains(reqUri, "/api/v1/getCaptcha") {
	// 	sysOperLog.BusinessType = "12"
	// 	sysOperLog.Title = "验证码"
	// } else {
	// 	if reqMethod == "POST" {
	// 		sysOperLog.BusinessType = "1"
	// 	} else if reqMethod == "PUT" {
	// 		sysOperLog.BusinessType = "2"
	// 	} else if reqMethod == "DELETE" {
	// 		sysOperLog.BusinessType = "3"
	// 	}
	// }
	// sysOperLog.Method = reqMethod
	// if len(menuList) > 0 {
	// 	sysOperLog.Title = menuList[0].Title
	// }
	// b := svr.Request.Form.Get("body")
	// sysOperLog.OperParam, _ = xtools.StructToJsonStr(b)
	// sysOperLog.CreateBy = jwtauth.GetUserName(c)
	// sysOperLog.OperTime = xtools.GetCurrentTime()
	// sysOperLog.LatencyTime = (latencyTime).String()
	// sysOperLog.UserAgent = userAgent
	// if c.Err() == nil {
	// 	sysOperLog.Status = "0"
	// } else {
	// 	sysOperLog.Status = "1"
	// }
	// _, _ = global.AdminDao.CreateSysOperLog(&sysOperLog)
}
