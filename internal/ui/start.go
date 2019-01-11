package ui

import (
	"fmt"

	"github.com/richardwilkes/cef/cef"
	"github.com/richardwilkes/encounter/internal/server"
	"github.com/richardwilkes/toolbox/atexit"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/webapp"
	"github.com/richardwilkes/webapp/stdmenu"
)

func Start(args *cef.MainArgs, s *server.Server) {
	webapp.WillFinishStartupCallback = func() {
		wnd, err := webapp.NewWindow(webapp.StdWindowMask, webapp.MainDisplay().UsableBounds, fmt.Sprintf("%s v%s", cmdline.AppName, cmdline.AppVersion), fmt.Sprintf("http://127.0.0.1:%d", s.Port()))
		jot.FatalIfErr(err)
		if bar, global, first := webapp.MenuBarForWindow(wnd); !global || first {
			stdmenu.FillMenuBar(bar, aboutHandler, prefsHandler, true)
		}
		wnd.ToFront()
	}
	// Start only returns on error
	jot.FatalIfErr(webapp.Start(args, nil, nil))
	atexit.Exit(0)
}

func aboutHandler() {
	fmt.Println("About menu item selected")
}

func prefsHandler() {
	fmt.Println("Preferences menu item selected")
}
