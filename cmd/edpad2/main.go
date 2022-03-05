package main

import (
	"edpad2/internal/router"
	"os"
	"os/signal"
	"syscall"

	"github.com/pborman/getopt/v2"

	"github.com/maxb-odessa/sconf"
	"github.com/maxb-odessa/slog"
)

func main() {

	// get cmdline args and parse them
	help := false
	debug := 0
	configFile := "etc/edpad2.conf"
	getopt.HelpColumn = 0
	getopt.FlagLong(&help, "help", 'h', "Show this help")
	getopt.FlagLong(&debug, "debug", 'd', "Set debug log level")
	getopt.FlagLong(&configFile, "config", 'c', "Path to config file")
	getopt.Parse()

	// help-only requested
	if help {
		getopt.Usage()
		return
	}

	// setup logger
	slog.Init("edpad2", debug, "2006-01-02 15:04:05")

	slog.Info("started")

	// read config file
	if err := sconf.Read(configFile); err != nil {
		slog.Fatal("config failed: %s", err)
	}

	// start main loop: serve network<->backends requests
	if err := router.Init(); err != nil {
		slog.Fatal("router, failed to init: %s", err)
	}

	// register everything in router
	// router.Register(display...)???????? HERE!
	// defer ???????

	// run router
	go router.DoRouting()

	// wait for termiation
	// wait for a system signal
	done := make(chan bool)
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		for sig := range sigChan {
			slog.Info("got signal %d\n", sig)
			done <- true
		}
	}()

	<-done
	slog.Info("stopped")

	// all done
	return
}
