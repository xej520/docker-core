package main

import (
	"os/exec"
	"syscall"
	"os"
	"github.com/CodisLabs/codis/pkg/utils/log"
)

// uts namespace提供的是 主机名和域名的隔离
func main() {
	cmd := exec.Command("sh")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS|
			syscall.CLONE_NEWIPC|
			syscall.CLONE_NEWPID|
			syscall.CLONE_NEWNS|
			syscall.CLONE_NEWNET,
	}
	cmd.SysProcAttr.Credential = &syscall.Credential{
		Uid:uint32(1),
		Gid:uint32(1),
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Error("------>\t%v", err)
	}
	os.Exit(-1)
}


