package pty

import (
	"golang.org/x/term"
	"io"
	"net"
	"os"
)

// NewPty pty is pseudo terminal
func NewPty(conn net.Conn) {
	// important
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer func() { _ = term.Restore(int(os.Stdin.Fd()), oldState) }()

	go func() {
		_, _ = io.Copy(conn, os.Stdin)
	}()
	_, _ = io.Copy(os.Stdout, conn)
}
