package pty

import (
	"github.com/creack/pty"
	"golang.org/x/term"
	"io"
	"net"
	"os"
	"os/exec"
)

// NewTty tty is true terminal
func NewTty(conn net.Conn, cmd *exec.Cmd) {
	// important
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }()

	// start cmd
	ptmx, err := pty.Start(cmd)
	if err != nil {
		panic(err)
	}

	go func() { _, _ = io.Copy(ptmx, conn) }()
	_, _ = io.Copy(conn, ptmx)
}
