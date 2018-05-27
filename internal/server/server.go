package server

import (
	"log"
	"net/http"
	"time"

	"github.com/richardwilkes/encounter/board"
	"github.com/richardwilkes/encounter/combatant"
	"github.com/richardwilkes/encounter/internal/assets"
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
	}
	s.Server.WebServer.Handler = s

	c := s.board.NewCombatant("Billy Joe Bob")
	c.Enemy = false
	c.Initiative = 13
	c.HP.Full = 32
	c.HP.Damage = 9
	c.Notes = append(c.Notes, combatant.Note{
		Description: "Blinded",
		Timed:       true,
		Who:         c.ID,
		Round:       6,
	}, combatant.Note{
		Description: "Nauseated",
		Timed:       true,
		Who:         c.ID,
		Round:       3,
	})

	c = s.board.NewCombatant("Guard")
	c.Enemy = false
	c.Initiative = 11
	c.HP.Full = 32
	c.HP.Damage = 16
	c.Notes = append(c.Notes, combatant.Note{
		Description: "Haste",
		Timed:       true,
		Who:         c.ID,
		Round:       5,
	})

	c = s.board.NewCombatant("Orc")
	c.Initiative = 2
	c.HP.Full = 8

	c2 := s.board.DuplicateCombatant(c)
	c2.HP.Temporary = 2
	c2.HP.Damage = 4

	c = s.board.DuplicateCombatant(c)
	c.HP.Damage = 18
	c.Notes = append(c.Notes, combatant.Note{Description: "Missing left leg"})

	c = s.board.DuplicateCombatant(c)
	c.Notes = make([]combatant.Note, 0)
	c.HP.Damage--

	c = s.board.DuplicateCombatant(c)
	c.HP.Damage = c.HP.Full

	s.board.Current = s.board.Combatants[0]

	return s
}

// ServeHTTP is the top-level handler for http requests.
func (s *Server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch web.PathHeadThenShift(req) {
	case "":
		s.handleMain(w, req)
	case "index.html":
		http.Redirect(w, req, "/", http.StatusTemporaryRedirect)
	default:
		s.staticFS.ServeHTTP(w, req)
	}
}
