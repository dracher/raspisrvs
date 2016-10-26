package pistatus

import (
	"os/exec"
	"sync"

	"github.com/kataras/iris"
)

// PiStatus represent pistatus
type PiStatus struct {
	CPUTemp []byte
	GPUTemp []byte
	lock    *sync.Mutex
}

// NewPiStatus is init function
func NewPiStatus() *PiStatus {
	return &PiStatus{
		lock: &sync.Mutex{},
	}
}

func (p *PiStatus) update() {
	p.lock.Lock()
	defer p.lock.Unlock()

	// get cpu temp
	ct, err := exec.
		Command("cat", "/sys/class/thermal/thermal_zone0/temp").Output()
	if err != nil {
		iris.Logger.Println(err)
		ct = []byte("-1")
	}
	p.CPUTemp = ct

	// get gpu temp
	gt, err := exec.Command("vcgencmd", "measure_temp").Output()
	if err != nil {
		iris.Logger.Println(err)
		gt = []byte("-1")
	}
	p.GPUTemp = gt
}

// Get is
func (p *PiStatus) Get() map[string]string {
	p.update()
	ret := make(map[string]string)
	ret["cpu_t"] = string(p.CPUTemp)
	ret["gpu_t"] = string(p.GPUTemp)
	return ret
}
