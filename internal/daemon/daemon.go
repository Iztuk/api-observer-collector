package daemon

import (
	filemonitor "api-observer-collector/internal/file-monitor"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
	"time"
)

const pidFile = "/tmp/api-observer-collector.pid"

func StartDaemon(bg bool) error {
	if bg {
		cmd := exec.Command(os.Args[0], "start", "--background=false")

		cmd.Stdout = nil
		cmd.Stderr = nil
		cmd.Stdin = nil

		cmd.SysProcAttr = &syscall.SysProcAttr{
			Setsid: true,
		}

		if err := cmd.Start(); err != nil {
			return fmt.Errorf("start background process: %w", err)
		}

		pid := cmd.Process.Pid

		if err := os.WriteFile(
			pidFile,
			[]byte(strconv.Itoa(pid)),
			0644,
		); err != nil {
			return fmt.Errorf("write pid file: %w", err)
		}

		fmt.Printf(
			"API Observer Collector running in background. PID %d\n",
			pid,
		)

		return nil
	} else {
		fmt.Printf(
			"API Observer Collector running in foreground. PID %d\n",
			os.Getpid(),
		)
	}

	err := filemonitor.MonitorFiles()
	if err != nil {
		return err
	}

	runDaemon()

	return nil
}

func StopDaemon() error {
	data, err := os.ReadFile(pidFile)
	if err != nil {
		return fmt.Errorf("read pid file: %w", err)
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return fmt.Errorf("invalid pid file: %w", err)
	}

	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("find process: %w", err)
	}

	if err := process.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("stop process: %w", err)
	}

	if err := os.Remove(pidFile); err != nil {
		return fmt.Errorf("remove pid file: %w", err)
	}

	fmt.Printf("Stopped API Observer Collector. PID %d\n", pid)

	return nil
}

func runDaemon() {
	file, _ := os.OpenFile("daemon.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	for {
		file.WriteString(fmt.Sprintf("Daemon ticker: %s\n", time.Now().Format(time.RFC3339)))
		time.Sleep(5 * time.Second)
	}
}
