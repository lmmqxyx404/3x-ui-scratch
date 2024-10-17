package xray

import (
	"os/exec"
	"time"
)

type Process struct {
	*process
}

type process struct {
	cmd       *exec.Cmd
	startTime time.Time
}

func (p *process) IsRunning() bool {
	if p.cmd == nil || p.cmd.Process == nil {
		return false
	}
	if p.cmd.ProcessState == nil {
		return true
	}
	return false
}

func (p *Process) GetUptime() uint64 {
	return uint64(time.Since(p.startTime).Seconds())
}
