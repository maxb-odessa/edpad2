package router

import (
	"fmt"

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

type Route struct {
	Src, Dst Endpoint
	Data     interface{}
}

type connector struct {
	readCh  <-chan Route
	writeCh chan<- Route
}

var endpoints map[Endpoint]connector

func Init() error {

	endpoints = make(map[Endpoint]connector)

	return nil
}

func Register(ep Endpoint, rc <-chan Route, wc chan<- Route) error {

	if _, ok := endpoints[ep]; !ok {
		return fmt.Errorf("endpoint already registered")
	}

	endpoints[ep] = connector{readCh: rc, writeCh: wc}

	return nil
}

// this must be called after all endpoints have registered!
func DoRouting() {
	for ep, con := range endpoints {
		slog.Debug(1, "routing endpoint %d", ep)
		go direct(ep, con.readCh)
	}
}

func direct(ep Endpoint, rc <-chan Route) {
	for {
		select {
		case data, ok := <-rc:
			if !ok {
				return
			}
			if dstEp, ok := endpoints[data.Dst]; ok {
				data.Src = ep
				dstEp.writeCh <- data
			} else {
				slog.Err("endpoint %d is not registered (src = %d)", data.Dst, ep)
			}
		}
	} //for
}
