package board

import (
	"fmt"
	"strings"

	"github.com/richardwilkes/rpgtools/dice"
	"github.com/richardwilkes/toolbox/collection"
)

type Board struct {
	Round      int
	Current    int
	Combatants []*Combatant
}

func (b *Board) NewCombatant(nameHint string) *Combatant {
	c := NewCombatant(b.SuggestName(nameHint))
	b.Combatants = append(b.Combatants, c)
	return c
}

func (b *Board) DuplicateCombatant(who *Combatant) *Combatant {
	c := who.Clone()
	c.Name = b.SuggestName(c.Name)
	b.Combatants = append(b.Combatants, c)
	return c
}

func (b *Board) IsCurrent(c *Combatant) bool {
	return b.Round > 0 && b.Current == c.ID
}

func (b *Board) CurrentTag(c *Combatant) string {
	if b.IsCurrent(c) {
		return " current"
	}
	return ""
}

func (b *Board) Lookup(id int) *Combatant {
	for _, c := range b.Combatants {
		if c.ID == id {
			return c
		}
	}
	return nil
}

func (b *Board) SuggestName(nameHint string) string {
	names := collection.NewStringSet()
	for _, c := range b.Combatants {
		names.Add(strings.ToLower(c.Name))
	}
	if !names.Contains(strings.ToLower(nameHint)) {
		return nameHint
	}
	nameHint = strings.TrimRight(nameHint, "0123456789 ")
	if !strings.HasSuffix(nameHint, "#") {
		nameHint = nameHint + " "
	}
	counter := 2
	for {
		name := fmt.Sprintf("%s%d", nameHint, counter)
		if !names.Contains(strings.ToLower(name)) {
			return name
		}
		counter++
	}
}

func (b *Board) RollInitiativeDice() int {
	return dice.Roll(nil, "d20")
}
