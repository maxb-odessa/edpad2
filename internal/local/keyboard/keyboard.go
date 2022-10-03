package keyboard

import (
	"edpad2/internal/router"

	"github.com/maxb-odessa/slog"
)

type Key struct {
	Code    uint32
	Pressed bool
}

type handler struct {
	// router data
	endpoint  router.Endpoint
	connector *router.Connector

	// priv data
	dev *dev
}

func Connect(ep router.Endpoint) (router.Endpoint, *router.Connector) {
	h := new(handler)
	h.endpoint = ep
	h.connector = new(router.Connector)
	h.connector.FromRouterCh = make(chan *router.Message, 8) // wait for router messages on this chan
	h.connector.ToRouterCh = make(chan *router.Message, 8)   // send messages to the router into this chan
	h.connector.DoneCh = make(chan bool)                     // termination chan

	h.dev = newDev("local keyboard")

	if err := h.dev.run(); err != nil {
		slog.Err("failed to start LocalKeyboard: %s", err)
		return ep, nil
	}

	go h.run()

	// all done, return router connector
	return ep, h.connector
}

func (h *handler) run() {

	// start messages processing
	for {

		select {

		case <-h.connector.DoneCh:
			close(h.connector.ToRouterCh)
			close(h.connector.FromRouterCh)
			slog.Debug(9, "closing keyboard handler")
			h.dev.finish()
			return

		case m := <-h.connector.FromRouterCh:
			slog.Debug(9, "local keyboard got msg! %+v", m)
			h.dev.writeKey(m.Data.(*Key))

		} //select...

	} //for...

}
