package server

import (
	"bytes"
	"net/http"

	"github.com/richardwilkes/encounter/board"
	"github.com/richardwilkes/encounter/internal/assets"
	"github.com/richardwilkes/toolbox/cmdline"
	"github.com/richardwilkes/toolbox/errs"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xio/fs/embedded/htmltmpl"
)

type Document struct {
	Title     string
	Copyright string
	Board     *board.Board
}

func (s *Server) handleMain(w http.ResponseWriter, req *http.Request) {
	tmpl, err := htmltmpl.Load(nil, assets.DynamicFS, "/", nil)
	if err != nil {
		jot.Error(errs.Wrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var buffer bytes.Buffer
	if err := tmpl.ExecuteTemplate(&buffer, "/index.html", &Document{
		Title:     cmdline.AppName,
		Copyright: cmdline.Copyright(),
		Board:     &s.board,
	}); err != nil {
		jot.Error(errs.Wrap(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if _, err := buffer.WriteTo(w); err != nil {
		jot.Warn(errs.Wrap(err))
	}
}
