
package kcp

import (
	"context"
	"fmt"
	"net"
)

// Available returns whether or not SO_REUSEPORT or equivalent behaviour is
// available in the OS.
func Available() bool {
	return true
}

// var listenConfig = net.ListenConfig{
// 	Control: Control,
// }

func newListenerConfig(device string) net.ListenConfig {
	return net.ListenConfig{
		Control: GetBindToDeviceControl(device),
	}
}

// Listen listens at the given network and address. see net.Listen
// Returns a net.Listener created from a file discriptor for a socket
// with SO_REUSEPORT and SO_REUSEADDR option set.
func ReUseListen(device, network, address string) (net.Listener, error) {
	listenConfig := newListenerConfig(device)
	return listenConfig.Listen(context.Background(), network, address)
}

// ListenPacket listens at the given network and address. see net.ListenPacket
// Returns a net.Listener created from a file discriptor for a socket
// with SO_REUSEPORT and SO_REUSEADDR option set.
func ListenPacket(device, network, address string) (net.PacketConn, error) {
	listenConfig := newListenerConfig(device)
	return listenConfig.ListenPacket(context.Background(), network, address)
}

// Dial dials the given network and address. see net.Dialer.Dial
// Returns a net.Conn created from a file descriptor for a socket
// with SO_REUSEPORT and SO_REUSEADDR option set.
func ReUseDial(device, network, laddr, raddr string) (net.Conn, error) {
	nla, err := ResolveAddr(network, laddr)
	if err != nil {
		return nil, fmt.Errorf("failed to resolve local addr: %w", err)
	}
	d := net.Dialer{
		Control:   GetBindToDeviceControl(device),
		LocalAddr: nla,
	}
	return d.Dial(network, raddr)
}
