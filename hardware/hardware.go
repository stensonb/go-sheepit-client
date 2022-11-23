package hardware

import (
	"log"
	"path"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func PlatformName() string {
	return "linux"
}

func RendererPath() string {
	return path.Join("Blender", "blender.app", "Contents", "MacOS", "blender")
}

func CpuStat() *Cpu {

	//cpu := Cpu{Family: "6", Name: "Intel(R) Core(TM) i7-6700 CPU @ 3.40GHz", Model: "i7-6700", Cores: 8, Architecture: "64bit"}
	ans := Cpu{Family: "", Name: "", Model: "Unknown", TotalCores: 0, Architecture: "32bit"}

	is, err := cpu.Info()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v", is[0])

	coreCount, err := cpu.Counts(false)
	if err != nil {
		log.Fatal(err)
	}

	hostArch, err := host.KernelArch()
	if err != nil {
		log.Fatal(err)
	}

	ans.Family = is[0].Family
	ans.Name = is[0].VendorID
	ans.Model = is[0].ModelName
	ans.TotalCores = coreCount
	ans.Architecture = hostArch

	return &ans
}

func GpuStat() *Gpu {
	return &Gpu{}
}

func TotalMemory() uint64 {
	v, err := mem.VirtualMemory()
	if err != nil {
		log.Fatal(err)
	}

	return v.Total
}
