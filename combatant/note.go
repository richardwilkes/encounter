package combatant

import "github.com/satori/go.uuid"

// Note holds a note about a combatant, which may be set to expire at a given
// point in time.
type Note struct {
	Description string    `json:"description"`
	Timed       bool      `json:"timed"`
	UntilEnd    bool      `json:"until_end"`
	Who         uuid.UUID `json:"who"`
	Round       int       `json:"round"`
}
