package server

import (
	"github.com/twitter/scoot/daemon/protocol"
	"net"
	"os"
	"path"
	"syscall"
)

func Listen() (net.Listener, error) {
	socketPath, err := protocol.LocateSocket()
	if err != nil {
		return nil, err
	}
	err = os.MkdirAll(path.Dir(socketPath), 0700)
	if err != nil {
		return nil, err
	}
	l, err := net.Listen("unix", socketPath)
	if err == nil {
		return l, nil
	}

	l = replaceDeadServer(socketPath, err)
	if l != nil {
		return l, nil
	}
	return nil, err
}

// replaceDeadServer handles the common case where a daemon has died but the socket file still exists.
// If the address is already in use, that means the file exists but there may not be a live server attached to it.
// (If we get an error other than in use, then we're not sure what's going on. Maybe it's a file instead of a socket?
// Don't make the situation worse, just error out.)
// To see if it's live, we try connecting to it. If we get connection refused, we infer the server is dead.
// So we remove the socket file, and then try serving.
// Returns a valid listener or nil.
func replaceDeadServer(socketPath string, err error) net.Listener {
	if !isErrno(err, syscall.EADDRINUSE) {
		return nil
	}

	conn, connErr := net.Dial("unix", socketPath)
	if connErr == nil {
		// There is an active server, so bow out gracefully
		conn.Close()
		return nil
	}
	if !isErrno(connErr, syscall.ECONNREFUSED) {
		return nil
	}
	pathErr := os.Remove(socketPath)
	if pathErr != nil {
		return nil
	}
	l, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil
	}
	return l
}

// Checks if an err is an errno that came from a call to a net function.
// This involves unwrapping the errors as we expect to receive them from the net package.
func isErrno(err error, expected syscall.Errno) bool {
	opErr, ok := err.(*net.OpError)
	if !ok {
		return false
	}

	sysErr, ok := opErr.Err.(*os.SyscallError)
	if !ok {
		return false
	}
	errno, ok := sysErr.Err.(syscall.Errno)
	if !ok {
		return false
	}
	return errno == expected
}
