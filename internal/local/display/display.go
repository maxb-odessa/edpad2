package display

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"

	"edpad2/internal/router"
)

const (
	VP_ROUTE int = iota
	VP_INFO
	VP_SYSTEM
	VP_SIGNALS
	VP_PLANETS
	VP_SRV
)

// we expect this from file filter
type Text struct {
	ViewPort       int
	Text           string
	UpdateText     bool
	AppendText     bool
	Subtitle       string
	UpdateSubtitle bool
	SetVisible     bool
}

type viewPort struct {
	title  string
	tvName string
	swName string
	sw     *gtk.ScrolledWindow
	tv     *gtk.TextView
	buff   *gtk.TextBuffer
	mark   *gtk.TextMark
}

type handler struct {
	endpoint    router.Endpoint
	connector   *router.Connector
	resourceDir string
	gtkBuilder  *gtk.Builder
	viewPorts   map[int]*viewPort
	gtkStack    *gtk.Stack
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

	// run!
	go h.run()

	// all done, return router connector
	return ep, h.connector
}

func (h *handler) init() (err error) {

	runtime.LockOSThread()

	h.resourceDir = os.ExpandEnv(sconf.StrDef("local display", "resource dir", "$HOME/.local/share/edpad2"))
	slog.Debug(9, "endpoint '%s': resource dir is '%s'", h.endpoint, h.resourceDir)

	// read ui into mem
	bres, err := ioutil.ReadFile(h.resourceDir + "/edpad2.ui")
	if err != nil {
		return
	}

	gtk.Init(nil)

	h.gtkBuilder, err = gtk.BuilderNew()
	if err != nil {
		return
	}

	if err = h.gtkBuilder.AddFromString(string(bres)); err != nil {
		return
	}

	obj, err := h.gtkBuilder.GetObject("top")
	if err != nil {
		return
	}

	top, ok := obj.(*gtk.Window)
	if !ok {
		err = fmt.Errorf("GTK3 error: obj 'top' is not a window")
		return
	}

	top.Connect("destroy", func() {
		gtk.MainQuit()
		os.Exit(0) // TODO graceful exit
	})

	css, _ := gtk.CssProviderNew()
	css.LoadFromPath(h.resourceDir + "/edpad2.css")
	if err != nil {
		return err
	}

	screen, _ := gdk.ScreenGetDefault()
	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_USER)

	// init and run keypad processor
	if err = h.keypad(); err != nil {
		return
	}

	st, err := h.gtkBuilder.GetObject("stack")
	if err != nil {
		return err
	}
	h.gtkStack = st.(*gtk.Stack)

	// init and setup text view ports

	//h.viewPorts = make(map[int]*viewPort)
	h.viewPorts = map[int]*viewPort{
		VP_ROUTE:   &viewPort{title: "Route", tvName: "route_tv", swName: ""},
		VP_INFO:    &viewPort{title: "Info", tvName: "info_tv", swName: ""},
		VP_SYSTEM:  &viewPort{title: "System", tvName: "system_tv", swName: "system_sw"},
		VP_PLANETS: &viewPort{title: "Planets", tvName: "planets_tv", swName: "planets_sw"},
		VP_SIGNALS: &viewPort{title: "Signals", tvName: "signals_tv", swName: "signals_sw"},
		VP_SRV:     &viewPort{title: "SRV", tvName: "srv_tv", swName: "srv_sw"},
	}

	for _, vp := range h.viewPorts {

		if obj, err = h.gtkBuilder.GetObject(vp.tvName); err != nil {
			return err
		} else {
			vp.tv = obj.(*gtk.TextView)
			if vp.buff, err = vp.tv.GetBuffer(); err != nil {
				return err
			}
			vp.mark = vp.buff.CreateMark("mark", vp.buff.GetEndIter(), false)
			// this viewport doesn't have scrolled window
			if vp.swName == "" {
				continue
			}
			if sw, err := h.gtkBuilder.GetObject(vp.swName); err != nil {
				return err
			} else {
				vp.sw = sw.(*gtk.ScrolledWindow)
			}
		}

	} //for

	// inits done, show the window
	top.ShowAll()

	if ok := sconf.BoolDef("local display", "maximize window", false); ok {
		top.Maximize()
	}

	go gtk.Main()

	return
}

func (h *handler) run() {

	for {

		select {

		case <-h.connector.DoneCh:

			close(h.connector.ToRouterCh)
			close(h.connector.FromRouterCh)
			return

		case d, ok := <-h.connector.FromRouterCh:

			if !ok {
				return // must exit
			}

			// we support file filter only atm
			//if d.Src != router.FilterFile {
			//	continue
			//}

			// display data from file filter
			glib.IdleAdd(func() bool { return h.printText(d.Data.(*Text)) })

		} //select

	} //for
}

func (h *handler) printText(t *Text) (ret bool) {

	ret = false

	vp, ok := h.viewPorts[t.ViewPort]
	if !ok {
		slog.Warn("endpoint '%s': undefind viewport %d", t.ViewPort)
		return
	}

	// update title
	if vp.sw != nil && t.UpdateSubtitle {
		h.gtkStack.ChildSetProperty(vp.sw, "title", vp.title+"\n"+t.Subtitle)
	}

	// print the text
	if t.UpdateText {
		// clear viewport if needed
		if !t.AppendText {
			vp.buff.SetText("")
		}
		vp.buff.InsertMarkup(vp.buff.GetEndIter(), t.Text)
		vp.tv.ScrollToMark(vp.mark, 0.0, false, 0.0, 0.0)
	}

	// make it visible
	if t.SetVisible {
		vp.sw.Show()
		h.gtkStack.SetVisibleChildName(vp.title)
	}

	/* TODO insert images?
	   img, _ := gtk.ImageNewFromFile(cfg.GtkResourcesDir + "/img.png")
	   ch, _ := vp.buff.CreateChildAnchor(vp.buff.GetEndIter())
	   vp.view.AddChildAtAnchor(img, ch)
	   img.Show()
	*/

	return
}
