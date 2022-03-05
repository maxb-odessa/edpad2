package router

import (
	"github.com/maxb-odessa/slog"
)

type Endpoint int

const (
	Unknown Endpoint = iota

	NetJoystick
	NetFile
	NetKeyboard
	NetSound

	LocalSound
	LocalDisplay

	FilterFile
)

func (ep Endpoint) String() string {
	switch ep {
	case NetJoystick:
		return "NetJoystick"
	case NetFile:
		return "NetFile"
	case NetKeyboard:
		return "NetKeyboard"
	case NetSound:
		return "NetSound"
	case LocalSound:
		return "LocalSound"
	case LocalDisplay:
		return "LocalDisplay"
	case FilterFile:
		return "FilterFile"
	}
	return "Unknown"
}

type Message struct {
	Src, Dst Endpoint
	Data     interface{}
}

type Connector struct {
	ReadCh  <-chan Message
	WriteCh chan<- Message
}

var endpoints map[Endpoint]*Connector

func Init() error {

	endpoints = make(map[Endpoint]*Connector)

	return nil
}

func Register(ep Endpoint, con *Connector) error {

	if _, ok := endpoints[ep]; ok {
		slog.Fatal("endpoint %s is already registered!", ep)
	}

	slog.Debug(1, "router: registered endpoint %s", ep)
	endpoints[ep] = con

	return nil
}

// TODO
func Unregister(ep Endpoint) {
	//close(endpoints[ep].ReadCh)
	//close(endpoints[ep].WriteCh)
}

// this must be called after all endpoints have registered!
func DoRouting() {
	for ep, con := range endpoints {
		slog.Debug(5, "routing endpoint %s", ep)
		go direct(ep, con.ReadCh)
	}
}

func direct(ep Endpoint, rc <-chan Message) {
	for {
		select {
		case data, ok := <-rc:
			if !ok {
				return
			}
			if dstEp, ok := endpoints[data.Dst]; ok {
				data.Src = ep
				dstEp.WriteCh <- data
			} else {
				slog.Err("endpoint %s is not registered (src = %s)", data.Dst, ep)
			}
		}
	} //for
}
