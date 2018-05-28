package server

import (
	"log"
	"net/http"
	"time"

	"github.com/richardwilkes/encounter/board"
	"github.com/richardwilkes/encounter/internal/assets"
	"github.com/richardwilkes/rpgtools/dice"
	"github.com/richardwilkes/toolbox/log/jot"
	"github.com/richardwilkes/toolbox/xio/network/xhttp/web"
)

const connectionTimeout = 5 * time.Second

type ctxPathType int

var ctxPath ctxPathType = 1

// Server holds the data necessary for the server.
type Server struct {
	web.Server
	staticFS http.Handler
	board    board.Board
}

// New creates a new server.
func New(address string) *Server {
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
		},
		staticFS: http.FileServer(assets.StaticFS),
		board:    board.Board{InitiativeDice: dice.New(nil, "1d20")},
	}
	s.Server.WebServer.Handler = s
	return s
}

// ServeHTTP is the top-level handler for http requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch web.PathHeadThenShift(req) {
	case "":
		s.handleMain(w, req)
	case "index.html":
		http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
	case "cmds":
		s.handleCmds(w, req)
	default:
		s.staticFS.ServeHTTP(w, req)
	}
}
