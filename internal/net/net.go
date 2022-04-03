package net

import (
	"context"
	"edpad2/internal/local/display"
	"edpad2/internal/router"
	"fmt"
	"io"
	"sync"

	pb "github.com/maxb-odessa/gamenode/pkg/gamenodepb"
	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"

	"google.golang.org/grpc"
)

type handler struct {
	endpoint  router.Endpoint
	connector *router.Connector
	dst       router.Endpoint
	sender    func(interface{}) error
	reader    func() (interface{}, error)
}

var once sync.Once
var grpcConn *grpc.ClientConn
var grpcReady *sync.Mutex

func Connect(ep router.Endpoint) (router.Endpoint, *router.Connector) {

	h := new(handler)
	h.endpoint = ep
	h.connector = new(router.Connector)
	h.connector.FromRouterCh = make(chan *router.Message, 8) // wait for router messages on this chan
	h.connector.ToRouterCh = make(chan *router.Message, 8)   // send messages to the router into this chan
	h.connector.DoneCh = make(chan bool)                     // termination chan

	// will connect to configured server or to default "127.0.0.1:12346"
	addr := sconf.StrDef("net", "connect", "127.0.0.1:12346")

	grpcReady = &sync.Mutex{}
	grpcReady.Lock()

	// wait for the connection to be establish and inform everybody faiting for it
	cf := func() {
		slog.Debug(1, "grpc: connecting to '%s'", addr)
		grpcConn, _ = grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
		slog.Debug(1, "grpc: connected to '%s'", addr)
		grpcReady.Unlock()
	}

	go once.Do(cf)

	go h.Run()

	// all done, return router connector
	return ep, h.connector
}

func (h *handler) Run() error {

	ctx, cancel := context.WithCancel(context.Background())

	// go and wait for DoneCh event, then cancel sender and reader
	go func() {
		<-h.connector.DoneCh
		cancel()
		close(h.connector.ToRouterCh)
		close(h.connector.FromRouterCh)
	}()

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:   display.VP_SYSTEM,
			Text:       fmt.Sprintf("\n%s: connecting...", h.endpoint),
			UpdateText: true,
			AppendText: true,
		},
	}

	slog.Debug(1, "endpoint '%s': waiting for server", h.endpoint)
	grpcReady.Lock()
	slog.Debug(1, "endpoint '%s': connected to server", h.endpoint)

	h.connector.ToRouterCh <- &router.Message{
		Dst: router.LocalDisplay,
		Data: &display.Text{
			ViewPort:   display.VP_SYSTEM,
			Text:       fmt.Sprintf("\n%s: connected!", h.endpoint),
			UpdateText: true,
			AppendText: true,
		},
	}

	grpcReady.Unlock()

	client := pb.NewGameNodeClient(grpcConn)

	switch h.endpoint {

	case router.NetFile:
		stream, err := client.File(ctx)
		if err != nil {
			return err
		}
		h.dst = router.FilterFile
		h.sender = func(m interface{}) error {
			d := m.(*pb.FileMsg)
			return stream.Send(d)
		}
		h.reader = func() (interface{}, error) {
			return stream.Recv()
		}

	case router.NetJoystick:
		stream, err := client.Joy(ctx)
		if err != nil {
			return err
		}
		h.dst = router.FilterJoystick
		h.sender = func(m interface{}) error {
			d := m.(*pb.JoyMsg)
			return stream.Send(d)
		}
		h.reader = func() (interface{}, error) {
			return stream.Recv()
		}

	case router.NetKeyboard:
		stream, err := client.Kbd(ctx)
		if err != nil {
			return err
		}
		h.dst = router.FilterKeyboard
		h.sender = func(m interface{}) error {
			d := m.(*pb.KbdMsg)
			return stream.Send(d)
		}
		h.reader = func() (interface{}, error) {
			return stream.Recv()
		}

	case router.NetSound:
		h.dst = router.FilterSound
		stream, err := client.Snd(ctx)
		if err != nil {
			return err
		}
		h.sender = func(m interface{}) error {
			d := m.(*pb.SndMsg)
			return stream.Send(d)
		}
		h.reader = func() (interface{}, error) {
			return stream.Recv()
		}

	}

	// run sender and reader
	go h.runReader()
	go h.runSender()

	return nil
}

func (h *handler) runReader() {

	for {

		msg, err := h.reader()

		// remote closed
		if err == io.EOF {
			return
		}

		// cancelled, other error
		if err != nil {
			//slog.Debug(5, "endpoint '%s' failed to receive: %s", h.endpoint, err)
			return
		}

		slog.Debug(9, "endpoint '%s': got msg: '%+v'", h.endpoint, msg)

		h.connector.ToRouterCh <- &router.Message{Dst: h.dst, Data: msg}

	} //for

}

func (h *handler) runSender() {

	for {

		select {
		case data, ok := <-h.connector.FromRouterCh:
			if !ok {
				return // must exit
			}
			msg := data.Data
			if err := h.sender(msg); err != nil {
				slog.Warn("endpoint '%s': failed to send: %s", h.endpoint, err)
				return
			}
		}

	} //for
}
