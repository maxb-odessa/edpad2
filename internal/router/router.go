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

	for ep, con := range endpoints {
		slog.Debug(5, "routing endpoint %s", ep)
		wg.Add(1)
		e := ep
		go func() {
			defer wg.Done()
			direct(e, con.ToRouterCh)
		}()
	}

	wg.Wait()

}

func direct(ep Endpoint, rc <-chan *Message) {
	for {
		select {
		case data, ok := <-rc:
			// read chan closed - the endpoint is terminated
			if !ok {
				slog.Debug(5, "director: endpoint '%s' closed its read chan", ep)
				return
			}
			if dstEp, ok := endpoints[data.Dst]; ok {
				data.Src = ep
				slog.Debug(9, "director: writing to ep '%s': %+v", ep, data)
				dstEp.FromRouterCh <- data
			} else {
				// impossible case (?)
				slog.Err("endpoint %s is not registered (src = %s)", data.Dst, ep)
			}
		}
	} //for
}
