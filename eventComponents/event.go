package eventComponents

import "net"

type Event struct {
	Message  string
	Client   net.Conn
	ClientID int
}
