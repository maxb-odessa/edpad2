package file

import (
	"github.com/maxb-odessa/slog"

	pb "github.com/maxb-odessa/gamenode/pkg/gamenodepb"

	"edpad2/internal/router"
)

type handler struct {
	endpoint  router.Endpoint
	connector *router.Connector
}

func Connect(ep router.Endpoint) (router.Endpoint, *router.Connector) {
	h := new(handler)
	h.endpoint = ep
	h.connector = new(router.Connector)
	h.connector.FromRouterCh = make(chan *router.Message, 8) // wait for router messages on this chan
	h.connector.ToRouterCh = make(chan *router.Message, 8)   // send messages to the router into this chan
	h.connector.DoneCh = make(chan bool)                     // termination chan

	go h.run()

	// all done, return router connector
	return ep, h.connector
}

func (h *handler) run() {

	// configure supported processors
	processors := map[string]func(*pb.FileEvent){
		"ed journal": func(m *pb.FileEvent) { h.processJournalMsg(m) },
	}

	// start messages processing
	for {

		select {

		case <-h.connector.DoneCh:
			close(h.connector.ToRouterCh)
			close(h.connector.FromRouterCh)
			return

		case m := <-h.connector.FromRouterCh:
			slog.Debug(9, "file filter got msg! %+v", m)

			if m.Src == router.LocalFile {
				h.processJournalLine(m.Data.(string))
			}

			if m.Src == router.NetFile {
				msg := m.Data.(*pb.FileMsg)
				name := msg.GetName()
				if proc, ok := processors[name]; !ok {
					slog.Warn("filter file: unconfigured source file '%s'", name)
				} else {
					proc(msg.GetEvent())
				}
			}

		} //select...

	} //for...

}
