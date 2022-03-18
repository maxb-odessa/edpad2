package router

import (
	"sync"

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

	FilterJoystick
	FilterFile
	FilterKeyboard
	FilterSound
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
	case FilterJoystick:
		return "FilterJoystick"
	case FilterFile:
		return "FilterFile"
	case FilterKeyboard:
		return "FilterKeyboard"
	case FilterSound:
		return "FilterSound"

	}
	return "Unknown"
}

type Message struct {
	Src, Dst Endpoint
	Data     interface{}
}

type Connector struct {
	ToRouterCh   chan *Message
	FromRouterCh chan *Message
	DoneCh       chan bool
}

var endpoints map[Endpoint]*Connector

func Init() error {

	endpoints = make(map[Endpoint]*Connector)

	return nil
}

func Register(ep Endpoint, con *Connector) error {

	if con == nil {
		slog.Fatal("router: endpoint '%s' failed to init", ep)
	}

	if _, ok := endpoints[ep]; ok {
		slog.Fatal("endpoint %s is already registered!", ep)
	}

	slog.Debug(1, "router: registered endpoint %s", ep)
	endpoints[ep] = con

	return nil
}

func Unregister(ep Endpoint) {
	slog.Debug(1, "router: unregistering '%s'", ep)
	endpoints[ep].DoneCh <- true
}

// this must be called after all endpoints have registered!
func DoRouting() {

	// this main loop will spawn N directors (each connected to its endpoint read chan)
	// each director lives while it can read from endpoint chan
	var wg sync.WaitGroup

	for e, c := range endpoints {
		slog.Debug(5, "routing endpoint %s", e)
		wg.Add(1)
		ep := e
		con := c
		go func() {
			defer wg.Done()
			distribute(ep, con.ToRouterCh)
		}()
	}

	wg.Wait()

}

func distribute(ep Endpoint, rc <-chan *Message) {
	for {
		select {
		case data, ok := <-rc:
			// read chan closed - the endpoint is terminated
			if !ok {
				slog.Debug(5, "distribute: endpoint '%s' closed its read chan", ep)
				return
			}
			if dstEp, ok := endpoints[data.Dst]; ok {
				data.Src = ep
				slog.Debug(99, "distribute: writing: %+v", data)
				dstEp.FromRouterCh <- data
			} else {
				// impossible case (?)
				slog.Err("distribute: endpoint %s is not registered (src = %s)", data.Dst, ep)
			}
		}
	} //for
}
