package file

import (
	"edpad2/internal/router"

	"github.com/maxb-odessa/slog"
)

type handler struct {
	// router data
	endpoint  router.Endpoint
	connector *router.Connector

	// priv data
	dev *dev
}

func (h *handler) init() error {
	h.dev = newDev("local file")
	return nil
}

func Connect(ep router.Endpoint) (router.Endpoint, *router.Connector) {
	h := new(handler)
	h.endpoint = ep
	h.connector = new(router.Connector)
	h.connector.FromRouterCh = make(chan *router.Message, 8) // wait for router messages on this chan
	h.connector.ToRouterCh = make(chan *router.Message, 8)   // send messages to the router into this chan
	h.connector.DoneCh = make(chan bool)                     // termination chan

	if err := h.init(); err != nil {
		slog.Err("endpoint '%s': failed: %s", ep, err)
		return ep, nil
	}

	h.dev.run()
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
			return

		case l, ok := <-h.dev.linesCh:
			if ok {
				slog.Debug(9, "got journal line: %s\n", l)
				h.connector.ToRouterCh <- &router.Message{
					Dst:  router.FilterFile,
					Data: l,
				}
			}

		} //select...

	} //for...

}
