package main

import (
	
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
)

func mb(b uint64) uint64 {
	return b / 1024 / 1024
}

func gb(b uint64) uint64 {
	return b / 1024 / 1024 / 1024
}

func main() {
	log.SetFlags(log.Ldate | log.Ltime)

	// Host info (logged once)
	hostInfo, err := host.Info()
	if err == nil {
		log.Println("===== SYSTEM INFO =====")
		log.Println("Hostname:", hostInfo.Hostname)
		log.Println("OS:", hostInfo.OS)
		log.Println("Platform:", hostInfo.Platform, hostInfo.PlatformVersion)
		log.Println("Uptime:", hostInfo.Uptime, "seconds")
		log.Println("=======================\n")
	}

	// Current process
	proc, _ := process.NewProcess(int32(os.Getpid()))

	for {
		// CPU
		cpuPercent, _ := cpu.Percent(time.Second, false)

		// Memory
		vm, _ := mem.VirtualMemory()

		// Disk (root volume on macOS)
		diskUsage, _ := disk.Usage("/")

		// Network
		netIO, _ := net.IOCounters(false)

		// Process stats
		procCPU, _ := proc.CPUPercent()
		procMem, _ := proc.MemoryInfo()

		log.Println("----- SYSTEM METRICS -----")

		// CPU
		log.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])

		// RAM
		log.Printf(
			"RAM: %d MB used / %d MB total (%.2f%%)\n",
			mb(vm.Used),
			mb(vm.Total),
			vm.UsedPercent,
		)

		// Disk
		log.Printf(
			"Disk: %d GB used / %d GB total (%.2f%%)\n",
			gb(diskUsage.Used),
			gb(diskUsage.Total),
			diskUsage.UsedPercent,
		)

		// Network
		if len(netIO) > 0 {
			log.Printf(
				"Network: Sent=%d MB | Received=%d MB\n",
				mb(netIO[0].BytesSent),
				mb(netIO[0].BytesRecv),
			)
		}

		// Process
		log.Printf(
			"Process: CPU=%.2f%% | RAM=%d MB\n",
			procCPU,
			mb(procMem.RSS),
		)

		log.Println("--------------------------\n")

		time.Sleep(3 * time.Second)
	}
}
