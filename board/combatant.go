package board

import "github.com/satori/go.uuid"

// Combantant holds information for a single entity in combat.
type Combatant struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	Enemy          bool         `json:"enemy"`
	Strength       AbilityScore `json:"str"`
	Dexterity      AbilityScore `json:"dex"`
	Constitution   AbilityScore `json:"con"`
	Intelligence   AbilityScore `json:"int"`
	Wisdom         AbilityScore `json:"wis"`
	Charisma       AbilityScore `json:"cha"`
	Initiative     int          `json:"init"`
	InitativeBonus int          `json:"init_bonus"`
	HP             HitPoints    `json:"hp"`
	Notes          []Note       `json:"notes,omitempty"`
}

// New creates a new combatant.
func NewCombatant(name string) *Combatant {
	return &Combatant{
		ID:           uuid.Must(uuid.NewV4()).String(),
		Name:         name,
		Enemy:        true,
		Strength:     AbilityScore{Value: 10},
		Dexterity:    AbilityScore{Value: 10},
		Constitution: AbilityScore{Value: 10},
		Intelligence: AbilityScore{Value: 10},
		Wisdom:       AbilityScore{Value: 10},
		Charisma:     AbilityScore{Value: 10},
		HP:           HitPoints{Full: 6},
	}
}

// Clone creates a copy of the combatant, but with a new ID.
func (c *Combatant) Clone() *Combatant {
	clone := *c
	clone.ID = uuid.Must(uuid.NewV4()).String()
	clone.Notes = make([]Note, len(c.Notes))
	copy(clone.Notes, c.Notes)
	return &clone
}

func (c *Combatant) Type() string {
	if c.Enemy {
		return "enemy"
	}
	return "ally"
}

func (c *Combatant) TypeDescription() string {
	if c.Enemy {
		return "Enemy"
	}
	return "Ally"
}

func (c *Combatant) Out() bool {
	return c.HP.Current() < 0
}

// Status returns a description of the combatant's current health.
func (c *Combatant) Status() string {
	if c.HP.Damage == 0 {
		return "Healthy"
	}
	current := c.HP.Current()
	if current > c.HP.Full/2 {
		return "Hurt"
	}
	if current > 0 {
		return "Bloody"
	}
	if current == 0 {
		return "Staggered"
	}
	if current > -c.Constitution.Current() {
		return "Unconscious"
	}
	return "Dead"
}

func (c *Combatant) StatusTag() string {
	if c.HP.Current() <= c.HP.Full/2 {
		return " danger"
	}
	return ""
}
