package board

import (
	"fmt"
	"strings"

	"github.com/richardwilkes/encounter/board/data"
)

// Combatant holds information for a single entity in combat.
type Combatant struct {
	ID               int
	Name             string
	Enemy            bool
	Initiative       int
	InitiativeBase   int
	RandomTieBreaker int
	HPFull           int
	HPTemporary      int
	HPDamage         int
	AC               int
	TouchAC          int
	FlatAC           int
	FortitudeSave    int
	ReflexSave       int
	WillSave         int
	Attacks          string
	Entity           *data.Entity
	Notes            []Note
}

// NewCombatant creates a new combatant.
func NewCombatant(id int, name string) *Combatant {
	return &Combatant{
		ID:      id,
		Name:    name,
		Enemy:   true,
		AC:      10,
		TouchAC: 10,
		FlatAC:  10,
	}
}

// PossessiveName returns the possessive form of the combatant's name.
func (c *Combatant) PossessiveName() string {
	if strings.HasSuffix(strings.ToLower(c.Name), "s") {
		return c.Name + "'"
	}
	return c.Name + "'s"
}

// Clone creates a copy of the combatant, but with a new ID.
func (c *Combatant) Clone(newID int) *Combatant {
	clone := *c
	clone.ID = newID
	clone.Notes = make([]Note, len(c.Notes))
	copy(clone.Notes, c.Notes)
	return &clone
}

// Type returns the type of combatant, enemy or ally.
func (c *Combatant) Type() string {
	if c.Enemy {
		return "enemy"
	}
	return "ally"
}

// TypeDescription returns a human-readable type for the combatant.
func (c *Combatant) TypeDescription() string {
	if c.Enemy {
		return "Enemy"
	}
	return "Ally"
}

// Out returns true if the combatant is out of the combat.
func (c *Combatant) Out() bool {
	return c.CurrentHP() < 0
}

// Status returns a description of the combatant's current health.
func (c *Combatant) Status() string {
	if c.HPDamage == 0 {
		return "Healthy"
	}
	current := c.CurrentHP()
	if current > c.HPFull/2 {
		return "Hurt"
	}
	if current > 0 {
		return "Bloody"
	}
	if current == 0 {
		return "Staggered"
	}
	return "Dying"
}

// StatusTag returns the current status tag to use.
func (c *Combatant) StatusTag() string {
	if c.CurrentHP() <= c.HPFull/2 {
		return " danger"
	}
	return ""
}

// CurrentHP returns the current hit point total.
func (c *Combatant) CurrentHP() int {
	return c.HPFull + c.HPTemporary - c.HPDamage
}

// Heal a combatant.
func (c *Combatant) Heal(amount int) {
	if amount < 0 {
		return
	}
	if amount >= c.HPDamage {
		c.HPDamage = 0
		return
	}
	c.HPDamage -= amount
}

// Harm a combatant.
func (c *Combatant) Harm(amount int) {
	if amount < 0 {
		return
	}
	if c.HPTemporary >= amount {
		c.HPTemporary -= amount
		return
	}
	if c.HPTemporary > 0 {
		amount -= c.HPTemporary
		c.HPTemporary = 0
	}
	c.HPDamage += amount
}

// FormattedHP returns the HP formatted for display.
func (c *Combatant) FormattedHP() string {
	if c.HPTemporary == 0 {
		return fmt.Sprintf("%d", c.HPFull-c.HPDamage)
	}
	return fmt.Sprintf("%d+%d", c.HPFull-c.HPDamage, c.HPTemporary)
}
