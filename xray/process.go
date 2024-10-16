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

	exitErr   error
	logWriter *LogWriter

	version string
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

func (p *process) GetErr() error {
	return p.exitErr
}

func (p *process) GetResult() string {
	if len(p.logWriter.lastLine) == 0 && p.exitErr != nil {
		return p.exitErr.Error()
	}
	return p.logWriter.lastLine
}

func (p *process) GetVersion() string {
	return p.version
}
