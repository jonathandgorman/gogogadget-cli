package periscope

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SystemInfo struct {
	Hostname     string
	Platform     string
	OS           string
	Kernel       string
	CPU          string
	CPUCores     int32
	RAMTotal     uint64
	RAMAvailable uint64
	DiskTotal    uint64
	Uptime       string
}

func formatUptime(seconds uint64) string {
	duration := time.Duration(seconds) * time.Second
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	secondsRemaining := int(duration.Seconds()) % 60

	uptimeString := ""
	if days > 0 {
		uptimeString += fmt.Sprintf("%d days, ", days)
	}
	if hours > 0 {
		uptimeString += fmt.Sprintf("%d hours, ", hours)
	}
	if minutes > 0 {
		uptimeString += fmt.Sprintf("%d minutes, ", minutes)
	}
	uptimeString += fmt.Sprintf("%d seconds", secondsRemaining)
	return uptimeString
}

func Run() {

	fmt.Println("\n   ___        ___        ___         _          _           ___         _                      _ \n  / __|___   / __|___   / __|__ _ __| |__ _ ___| |_        | _ \\___ _ _(_)___ __ ___ _ __  ___| |\n | (_ / _ \\ | (_ / _ \\ | (_ / _` / _` / _` / -_)  _|_ _ _  |  _/ -_) '_| (_-</ _/ _ \\ '_ \\/ -_)_|\n  \\___\\___/  \\___\\___/  \\___\\__,_\\__,_\\__, \\___|\\__(_|_|_) |_| \\___|_| |_/__/\\__\\___/ .__/\\___(_)\n                                      |___/                                         |_|          ")

	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Usage("/")
	uptime := formatUptime(hostStat.Uptime)

	info := SystemInfo{
		Hostname:     hostStat.Hostname,
		Platform:     hostStat.Platform,
		OS:           hostStat.OS,
		Kernel:       hostStat.KernelVersion,
		CPU:          cpuStat[0].ModelName,
		CPUCores:     cpuStat[0].Cores,
		RAMTotal:     vmStat.Total / 1024 / 1024,
		RAMAvailable: vmStat.Available / 1024 / 1024,
		DiskTotal:    diskStat.Total / 1024 / 1024,
		Uptime:       uptime,
	}

	fmt.Println("System Information:")
	fmt.Println("-------------------")
	fmt.Printf("%-14s: %s\n", "Hostname", info.Hostname)
	fmt.Printf("%-14s: %s\n", "Platform", info.Platform)
	fmt.Printf("%-14s: %s\n", "OS", info.OS)
	fmt.Printf("%-14s: %s\n", "Kernel", info.Kernel)
	fmt.Printf("%-14s: %s\n", "CPU", info.CPU)
	fmt.Printf("%-14s: %d Cores\n", "CPU Cores", info.CPUCores)
	fmt.Printf("%-14s: %d MB / %d MB\n", "RAM", info.RAMAvailable, info.RAMTotal)
	fmt.Printf("%-14s: %d MB\n", "Disk Total", info.DiskTotal)
	fmt.Printf("%-14s: %s\n", "Uptime", info.Uptime)
}
