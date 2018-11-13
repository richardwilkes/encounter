package ui

import (
	"fmt"

	"github.com/richardwilkes/encounter/internal/server"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/webapp"
)

func Start(s *server.Server) {
	webapp.WillFinishStartupCallback = func() { createWindow(s) }
	webapp.Start()
}

func createWindow(s *server.Server) {
	wnd, err := webapp.NewWindow(webapp.StdWindowMask, webapp.MainDisplay().UsableBounds, fmt.Sprintf("http://127.0.0.1:%d", s.Port()))
	jot.FatalIfErr(err)
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
