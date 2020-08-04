package sshop

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"sync"

	"github.com/ninedraft/sshop/pkg/proc"
	"github.com/pkg/sftp"
	"go.uber.org/multierr"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

type Agent struct {
	ssh  *ssh.Client
	sftp *sftp.Client
}

func NewWithAgent(user, domain string) (*Agent, error) {
	var socket = os.Getenv("SSH_AUTH_SOCK")
	var conn, errDialSSH = net.Dial("unix", socket)
	if errDialSSH != nil {
		return nil, fmt.Errorf("open SSH_AUTH_SOCK: %v", errDialSSH)
	}

	agentClient := agent.NewClient(conn)
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			// Use a callback rather than PublicKeys so we only consult the
			// agent once the remote server wants it.
			ssh.PublicKeysCallback(agentClient.Signers),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	var clSSH, errSSH = ssh.Dial("tcp", domain, config)
	if errSSH != nil {
		return nil, fmt.Errorf("ssh: %w", errSSH)
	}
	var clSFTP, errSFTP = sftp.NewClient(clSSH)
	if errSFTP != nil {
		return nil, fmt.Errorf("sftp: %w", errSFTP)
	}
	return &Agent{
		ssh:  clSSH,
		sftp: clSFTP,
	}, nil
}

func (agent *Agent) Close() error {
	_ = agent.sftp.Close()
	_ = agent.ssh.Close()
	return nil
}

type Metrics struct {
	Mem float64
	CPU float64
}

func (agent *Agent) Metrics() (Metrics, error) {
	var wg = &sync.WaitGroup{}
	var errs = make(chan error)
	wg.Add(2)
	go func() { wg.Wait(); close(errs) }()

	var metrics Metrics
	go func() {
		defer wg.Done()
		var err error
		metrics.CPU, err = agent.cpu()
		errs <- err
	}()
	go func() {
		defer wg.Done()
		var err error
		metrics.Mem, err = agent.mem()
		errs <- err
	}()

	var errMetrics error
	for err := range errs {
		multierr.AppendInto(&errMetrics, err)
	}
	return metrics, errMetrics
}

func (agent *Agent) mem() (float64, error) {
	var file, errOpen = agent.sftp.Open("/proc/meminfo")
	if errOpen != nil {
		return -1, errOpen
	}
	defer func() { _ = file.Close() }()
	// TODO: add buffer pool
	var buf = &bytes.Buffer{}
	io.Copy(buf, file)
	return proc.Mem(buf)
}

func (agent *Agent) cpu() (float64, error) {
	var file, errOpen = agent.sftp.Open("/proc/stat")
	if errOpen != nil {
		return -1, errOpen
	}
	defer func() { _ = file.Close() }()
	// TODO: add buffer pool
	var buf = &bytes.Buffer{}
	io.Copy(buf, file)
	return proc.CPU(buf)
}
