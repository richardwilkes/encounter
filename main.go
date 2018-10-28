package main

//go:generate go run board/data/dataconverter/main.go
//go:generate mkembeddedfs --no-modtime --output internal/assets/dynamic_fs_gen.go --pkg assets --name dynamicFS --strip internal/assets/dynamic internal/assets/dynamic
//go:generate mkembeddedfs --no-modtime --output internal/assets/static_fs_gen.go --pkg assets --name staticFS --strip internal/assets/static internal/assets/static

import (
	"fmt"
	"runtime"

	"github.com/richardwilkes/encounter/internal/server"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/log/jotrotate"
	"github.com/richardwilkes/toolbox/xmath/geom"
	"github.com/richardwilkes/webapp"
)

func main() {
	cmdline.AppName = "Encounter"
	cmdline.AppCmdName = "encounter"
	cmdline.AppVersion = "0.1"
	cmdline.CopyrightYears = "2018"
	cmdline.CopyrightHolder = "Richard A. Wilkes"
	cmdline.License = "Mozilla Public License Version 2.0"

	address := "127.0.0.1:8001"

	cl := cmdline.New(true)
	cl.NewStringOption(&address).SetSingle('a').SetName("address").SetUsage(`Network address and port to listen on. Specify just the port (":8001") to listen on all connected networks`)
	jotrotate.ParseAndSetup(cl)

	runtime.LockOSThread()

	s := server.New(address)
	startedChan := make(chan bool, 1)
	go func() {
		if err := s.RunWithNotifyAtStart(startedChan); err != nil {
			jot.Error(errs.NewfWithCause(err, "%s shutdown unexpectedly", s.Protocol()))
		}
	}()
	<-startedChan
	webapp.WillFinishStartupCallback = func() {
		wnd := webapp.NewWindow(webapp.StdWindowMask, geom.Rect{
			Point: geom.Point{X: 20, Y: 20},
			Size:  geom.Size{Width: 1024, Height: 768},
		}, fmt.Sprintf("http://127.0.0.1:%d", s.Port()))
		wnd.SetTitle(fmt.Sprintf("%s v%s", cmdline.AppName, cmdline.AppVersion))
		bar := webapp.MenuBarForWindow(wnd)
		_, aboutItem, prefsItem := bar.InstallAppMenu()
		aboutItem.Handler = func() { fmt.Println("About menu item selected") }
		prefsItem.Handler = func() { fmt.Println("Preferences menu item selected") }
		bar.InstallEditMenu()
		bar.InstallWindowMenu()
		bar.InstallHelpMenu()
		wnd.ToFront()
	}
	webapp.Start()
}
