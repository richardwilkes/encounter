package board

import (
	"strconv"
	"strings"
)

// Note holds a note about a combatant, which may be set to expire at a given
// point in time.
type Note struct {
	Description string `json:"description"`
	Timed       bool   `json:"timed"`
	UntilEnd    bool   `json:"until_end"`
	Who         string `json:"who"`
	Round       int    `json:"round"`
}

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
		buffer.WriteString(c.Name)
		buffer.WriteString("'")
		if !strings.HasSuffix(c.Name, "s") {
			buffer.WriteString("s")
		}
	}
	buffer.WriteString(" turn on round ")
	buffer.WriteString(strconv.Itoa(n.Round))
	return buffer.String()
}
