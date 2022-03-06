package net

import (
	"context"
	"edpad2/internal/router"
	"fmt"
	"io"

	pb "github.com/maxb-odessa/gamenode/pkg/gamenodepb"
	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"

	"google.golang.org/grpc"
)

type handler struct {
	endpoint  router.Endpoint
	connector *router.Connector
	sender    func(interface{}) error
	reader    func() (interface{}, error)
}

func Connect(ep router.Endpoint) *router.Connector {

	h := new(handler)
	h.endpoint = ep
	h.connector = new(router.Connector)
	h.connector.FromRouterCh = make(chan *router.Message) // wait for router messages on this chan
	h.connector.ToRouterCh = make(chan *router.Message)   // send messages to the router into this chan
	h.connector.DoneCh = make(chan bool)                  // termination chan

	if err := h.Init(); err != nil {
		slog.Err("endpoint '%s': init failed: %s", ep, err)
		return nil
	}

	// all done, return router connector
	return h.connector
}

func (h *handler) Init() error {

	// will connect to configured server or to default "127.0.0.1:12346"
	addr := sconf.StrDef("net", "connect", "127.0.0.1:12346")

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("fail to dial: %s", err)
	}

	slog.Debug(1, "endpoint '%s': connected to '%s'", h.endpoint, addr)
	client := pb.NewGameNodeClient(conn)

	ctx, cancel := context.WithCancel(context.Background())

	switch h.endpoint {

	case router.NetFile:
		stream, err := client.File(ctx)
		if err != nil {
			return err
		}
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
		h.sender = func(m interface{}) error {
			d := m.(*pb.KbdMsg)
			return stream.Send(d)
		}
		h.reader = func() (interface{}, error) {
			return stream.Recv()
		}

	case router.NetSound:
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

	// go and wait for DoneCh event, then cancel sender and reader
	go func() {
		<-h.connector.DoneCh
		cancel()
		close(h.connector.ToRouterCh)
		close(h.connector.FromRouterCh)
	}()

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

		h.connector.ToRouterCh <- &router.Message{Dst: router.FilterFile, Data: msg}

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
