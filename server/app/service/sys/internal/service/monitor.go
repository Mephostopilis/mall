package service

import (
	"context"
	"fmt"
	"runtime"
	"strconv"

	pb "edu/api/sys/v1"
	"edu/pkg/net/ip"
	"edu/pkg/tools"

	"github.com/golang/protobuf/ptypes"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// @Summary 系统信息
// @Description 获取JSON
// @Tags 系统信息
// @Success 200 {object} pb.ApiReply "{"code": 200, "data": [...]}"
// @Router /api/v1/settings/serverInfo [get]
func (s *AdminService) GetServerInfo(ctx context.Context, req *pb.GetServerInfoRequest) (reply *pb.ApiReply, err error) {
	// mem
	mem, _ := mem.VirtualMemory()
	memUsedMB := int(mem.Used) / GB
	memTotalMB := int(mem.Total) / GB
	memFreeMB := int(mem.Free) / GB
	memUsedPercent := int(mem.UsedPercent)

	// cpu
	cpuCounts, _ := cpu.Counts(false)

	// disk
	dis, _ := disk.Usage("/")
	diskTotalGB := int(dis.Total) / GB
	diskFreeGB := int(dis.Free) / GB

	list := make([]*anypb.Any, 0)
	d := &pb.GetServerInfoReply{
		Os: &pb.GetServerInfoReply_OS{
			GoOs:         runtime.GOOS,
			Arch:         runtime.GOARCH,
			Mem:          int32(runtime.MemProfileRate),
			Compiler:     runtime.Compiler,
			Version:      runtime.Version(),
			NumGoroutine: int32(runtime.NumGoroutine()),
			Ip:           ip.GetLocaHonst(),
			ProjectDir:   tools.GetCurrentPath(),
		},
		Mem: &pb.GetServerInfoReply_MEM{
			Total: int32(memTotalMB),
			Used:  int32(memUsedMB),
			Free:  int32(memFreeMB),
			Usage: int32(memUsedPercent),
		},
		Cpu: &pb.GetServerInfoReply_Cpu{
			CpuInfo: make([]*pb.GetServerInfoReply_Cpu_Info, 0),
			Percent: make([]float64, 0),
			CpuNum:  int32(cpuCounts),
		},
		Disk: &pb.GetServerInfoReply_Disk{
			Total: int32(diskTotalGB),
			Free:  int32(diskFreeGB),
		},
	}
	cpuInfoStat, _ := cpu.Info()
	for _, cpuInfo := range cpuInfoStat {
		d.Cpu.CpuInfo = append(d.Cpu.CpuInfo, &pb.GetServerInfoReply_Cpu_Info{
			ModelName: cpuInfo.ModelName,
			Cores:     int32(cpuInfo.Cores),
		})
	}
	cpuPercent, _ := cpu.Percent(0, false)
	for _, percent := range cpuPercent {
		p := fmt.Sprintf("%.2f", percent)
		pl, _ := strconv.ParseFloat(p, 64)
		d.Cpu.Percent = append(d.Cpu.Percent, pl)
	}

	any, err1 := ptypes.MarshalAny(d)
	if err1 != nil {
		err = err1
		return
	}
	list = append(list, any)
	reply = &pb.ApiReply{
		Code:    0,
		Message: "OK",
		Count:   int32(1),
		Data:    list,
	}
	return
}
