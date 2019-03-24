package main

//go:generate go run board/data/dataconverter/main.go

import (
	"runtime"

	"github.com/richardwilkes/encounter/internal/server"
	"github.com/richardwilkes/encounter/internal/ui"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/log/jotrotate"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/driver"
)

func main() {
	cmdline.AppName = "Encounter"
	cmdline.AppCmdName = "encounter"
	cmdline.AppVersion = "0.1"
	cmdline.AppIdentifier = "com.trollworks.encounter"
	cmdline.CopyrightYears = "2018-2019"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.License = "Mozilla Public License Version 2.0"

	args, err := webapp.Initialize(driver.ForPlatform())
	jot.FatalIfErr(err)

	address := "127.0.0.1:0"
	serverOnly := false

	cl := cmdline.New(true)
	cl.NewStringOption(&address).SetSingle('a').SetName("address").SetUsage(`Network address and port to listen on. Specify just the port (":8001") to listen on all connected networks. Specify a port of 0 to pick a random port.`)
	cl.NewBoolOption(&serverOnly).SetSingle('s').SetName("server-only").SetUsage(`Do not put up a user interface, just run the server.`)
	jotrotate.ParseAndSetup(cl)

	runtime.LockOSThread() // Ensure the main thread is used when starting the UI
	s := server.New(address)
	go func() {
		if err := s.Run(); err != nil {
			jot.Error(errs.NewfWithCause(err, "%s shutdown unexpectedly", s.Protocol()))
		}
		atexit.Exit(0)
	}()
	<-s.StartedChan
	if serverOnly {
		select {} // Wait forever
	} else {
		ui.Start(args, s)
	}
	atexit.Exit(0)
}
