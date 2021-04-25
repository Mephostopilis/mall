package biz

import (
	"fmt"
	"sync"
	"time"

	"edu/pkg/net/http"

	"github.com/go-kratos/kratos/v2/log"
)

var timeFormat = "2006-01-02 15:04:05"
var retryCount = 3

var lock sync.Mutex

type JobCore struct {
	InvokeTarget   string
	Name           string
	JobId          int
	EntryId        int
	CronExpression string
	log            *log.Helper
}

// HttpJob 任务类型 http
type HttpJob struct {
	JobCore
}

//http 任务接口
func (h *HttpJob) Run() {

	startTime := time.Now()
	var count = 0
	/* 循环 */
LOOP:
	if count < retryCount {
		/* 跳过迭代 */
		str, err := http.Get(h.InvokeTarget)
		if err != nil {
			// 如果失败暂停一段时间重试
			fmt.Println(time.Now().Format(timeFormat), " [ERROR] mission failed! ", err)
			fmt.Printf(time.Now().Format(timeFormat)+" [INFO] Retry after the task fails %d seconds! %s \n", time.Duration(count)*time.Second, str)
			time.Sleep(time.Duration(count) * time.Second)
			goto LOOP
		}
		count = count + 1
	}
	// 结束时间
	endTime := time.Now()

	// 执行时间
	latencyTime := endTime.Sub(startTime)
	//TODO: 待完善部分
	//str := time.Now().Format(timeFormat) + " [INFO] JobCore " + string(h.EntryId) + "exec success , spend :" + latencyTime.String()
	//ws.SendAll(str)
	h.log.Infof(time.Now().Format(timeFormat), " [INFO] JobCore ", h, "exec success , spend :", latencyTime)
}

type ExecJob struct {
	JobCore
}

func (e *ExecJob) Run() {
	// startTime := time.Now()
	// var obj = jobList[e.InvokeTarget]
	// if obj == nil {
	// 	log.Warn(" ExecJob Run job nil", e)
	// 	return
	// }
	// CallExec(obj.(JobsExec))
	// // 结束时间
	// endTime := time.Now()

	// // 执行时间
	// latencyTime := endTime.Sub(startTime)
	//TODO: 待完善部分
	//str := time.Now().Format(timeFormat) + " [INFO] JobCore " + string(e.EntryId) + "exec success , spend :" + latencyTime.String()
	//ws.SendAll(str)
	// log.Info(" [INFO] JobCore ", e, "exec success , spend :", latencyTime)
}
