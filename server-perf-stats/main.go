package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
)

func VmemInfo() {
	vmem, _ := mem.VirtualMemory()
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", vmem.Total, vmem.Free, vmem.UsedPercent)
}

func CpuInfo() {
	cpu_host, _ := cpu.Info()
	fmt.Printf("%v\n\n", cpu_host)
}

func DiskInfo() {
	disk_partitions_stats, _ := disk.Partitions(true)
	for _, partition_info := range disk_partitions_stats {
		fmt.Printf("%v\n\n", partition_info)
	}
}

func main() {

	VmemInfo()
	CpuInfo()
	// DiskInfo()
}
