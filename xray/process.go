package xray

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"syscall"
	"time"
	"x-ui-scratch/config"
)

type Process struct {
	*process
}

type process struct {
	cmd       *exec.Cmd
	startTime time.Time

	exitErr   error
	logWriter *LogWriter

	version       string
	onlineClients []string

	apiPort int
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

func (p *process) Stop() error {
	if !p.IsRunning() {
		return errors.New("xray is not running")
	}
	return p.cmd.Process.Signal(syscall.SIGTERM)
}

func GetBinaryPath() string {
	return config.GetBinFolderPath() + "/" + GetBinaryName()
}

func GetBinaryName() string {
	return fmt.Sprintf("xray-%s-%s", runtime.GOOS, runtime.GOARCH)
}

func (p *Process) SetOnlineClients(users []string) {
	p.onlineClients = users
}

func (p *Process) GetAPIPort() int {
	return p.apiPort
}
