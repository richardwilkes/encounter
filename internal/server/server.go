// Copyright ©2018-2023 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/richardwilkes/encounter/board"
	"github.com/richardwilkes/encounter/board/data"
	"github.com/richardwilkes/encounter/internal/assets"
	"github.com/richardwilkes/rpgtools/dice"
	"github.com/richardwilkes/toolbox/log/jot"
	xfs "github.com/richardwilkes/toolbox/xio/fs"
	"github.com/richardwilkes/toolbox/xio/fs/paths"
	"github.com/richardwilkes/toolbox/xio/network/xhttp/web"
)

// Server holds the data necessary for the server.
type Server struct {
	web.Server
	staticFS  http.Handler
	funcMap   template.FuncMap
	boardFile string
	board     board.Board
}

// New creates a new server.
func New(address string) *Server {
	const connectionTimeout = 5 * time.Second
	s := &Server{
		Server: web.Server{
			Logger: &jot.Logger{},
			WebServer: &http.Server{
				Addr:         address,
				ReadTimeout:  connectionTimeout,
				WriteTimeout: connectionTimeout,
				IdleTimeout:  connectionTimeout,
				ErrorLog:     log.New(&jot.LoggerWriter{}, "", 0),
			},
			StartedChan: make(chan interface{}, 1),
		},
		staticFS: http.FileServer(http.FS(assets.StaticFS)),
		funcMap: template.FuncMap{
			"comma":     func(v int) string { return humanize.Comma(int64(v)) },
			"lowercase": strings.ToLower,
		},
		boardFile: filepath.Join(paths.AppDataDir(), "board.json"),
		board: board.Board{
			InitiativeDice: dice.New("1d20"),
			HPMethod:       board.AverageHPMethod,
		},
	}
	s.board.SetLibrarySelection(data.Entities[0])
	s.Server.WebServer.Handler = s
	s.Server.ShutdownCallback = s.HandleShutdown
	if xfs.FileExists(s.boardFile) {
		t := jot.Time("Loading previous board")
		if err := s.board.Load(s.boardFile); err != nil {
			jot.Error(err)
		}
		t.End()
	}
	return s
}

// ServeHTTP is the top-level handler for http requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch web.PathHeadThenShift(req) {
	case "":
		s.handleIndex(w, req)
	case "index.html":
		http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
	case "cmds":
		s.handleCmds(w, req)
	case "describe":
		s.handleDescribe(w, req)
	default:
		s.staticFS.ServeHTTP(w, req)
	}
}

func (s *Server) loadTemplates() (*template.Template, error) {
	return template.New("").Funcs(s.funcMap).ParseFS(assets.DynamicFS, "*.html")
}

// HandleShutdown is called during shutdown to save state.
func (s *Server) HandleShutdown() {
	t := jot.Time("Saving board")
	defer t.End()
	if err := os.MkdirAll(filepath.Dir(s.boardFile), 0755); err != nil {
		jot.Error(err)
		return
	}
	if err := s.board.Save(s.boardFile); err != nil {
		jot.Error(err)
	}
}
