package board

import (
	"strconv"
	"strings"
)

// Note holds a note about a combatant, which may be set to expire at a given
// point in time.
type Note struct {
	Description string `json:"description,omitempty"`
	Timed       bool   `json:"timed,omitempty"`
	UntilEnd    bool   `json:"until_end,omitempty"`
	Who         int    `json:"who,omitempty"`
	Round       int    `json:"round,omitempty"`
}

// Detail returns the detail to use in a tooltip.
func (n *Note) Detail(b *Board) string {
	if !n.Timed {
		return ""
	}
	var buffer strings.Builder
	buffer.WriteString("Expires at the ")
	if n.UntilEnd {
		buffer.WriteString("end")
	} else {
		buffer.WriteString("start")
	}
	buffer.WriteString(" of ")
	c := b.Lookup(n.Who)
	if c == nil {
		buffer.WriteString("<unknown>'s")
	} else {
		buffer.WriteString(c.PossessiveName())
	}
	buffer.WriteString(" turn on round ")
	buffer.WriteString(strconv.Itoa(n.Round))
	return buffer.String()
}
