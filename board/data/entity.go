// Copyright ©2018-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package data

import (
	"regexp"
	"strconv"
	"strings"
)

// Entity holds the static information for a combatant.
type Entity struct {
	ID                 int    `json:"id"`                             // 93 (replaced by generated ID)
	Name               string `json:"name"`                           // 0
	CR                 string `json:"cr,omitempty"`                   // 1
	XP                 int    `json:"xp,omitempty"`                   // 2
	Class              string `json:"class,omitempty"`                // 4
	Alignment          string `json:"alignment,omitempty"`            // 6
	Size               string `json:"size,omitempty"`                 // 7
	Type               string `json:"type,omitempty"`                 // 8
	SubType            string `json:"subtype,omitempty"`              // 9
	Initiative         string `json:"initiative,omitempty"`           // 10
	Senses             string `json:"senses,omitempty"`               // 11
	Aura               string `json:"aura,omitempty"`                 // 12
	AC                 string `json:"ac,omitempty"`                   // 13
	ACMods             string `json:"ac_mods,omitempty"`              // 14
	HD                 string `json:"hd,omitempty"`                   // 16
	HPMods             string `json:"hp_mods,omitempty"`              // 17
	Saves              string `json:"saves,omitempty"`                // 18
	DefensiveAbilities string `json:"defensive_abilities,omitempty"`  // 23
	DR                 string `json:"dr,omitempty"`                   // 24
	Immunities         string `json:"immunities,omitempty"`           // 25
	Resistances        string `json:"resistances,omitempty"`          // 26
	SR                 string `json:"sr,omitempty"`                   // 27
	Weaknesses         string `json:"weaknesses,omitempty"`           // 28
	Speed              string `json:"speed,omitempty"`                // 29
	SpeedMod           string `json:"speed_mod,omitempty"`            // 30
	Melee              string `json:"melee,omitempty"`                // 31
	Ranged             string `json:"ranged,omitempty"`               // 32
	Space              string `json:"space,omitempty"`                // 33
	Reach              string `json:"reach,omitempty"`                // 34
	SpecialAttacks     string `json:"special_attacks,omitempty"`      // 35
	SpellLikeAbilities string `json:"spell_like_abilities,omitempty"` // 36
	SpellsKnown        string `json:"spells_known,omitempty"`         // 37
	SpellsPrepared     string `json:"spells_prepared,omitempty"`      // 38
	SpellDomains       string `json:"spell_domains,omitempty"`        // 39
	AbilityScores      string `json:"ability_scores,omitempty"`       // 40
	BaseAttack         int    `json:"base_attack,omitempty"`          // 42
	CMB                string `json:"cmb,omitempty"`                  // 43
	CMD                string `json:"cmd,omitempty"`                  // 44
	Feats              string `json:"feats,omitempty"`                // 45
	Skills             string `json:"skills,omitempty"`               // 46
	Languages          string `json:"languages,omitempty"`            // 48
	SpecialQualities   string `json:"special_qualities,omitempty"`    // 49
	Environment        string `json:"environment,omitempty"`          // 50
	Organization       string `json:"organization,omitempty"`         // 51
	Treasure           string `json:"treasure,omitempty"`             // 52
	DescriptionVisual  string `json:"description_visual,omitempty"`   // 53
	Source             string `json:"source,omitempty"`               // 55
	SpecialAbilities   string `json:"special_abilities,omitempty"`    // 57
	Description        string `json:"description,omitempty"`          // 58
	Gear               string `json:"gear,omitempty"`                 // 66
	OtherGear          string `json:"other_gear,omitempty"`           // 67
	MR                 int    `json:"mr,omitempty"`                   // 95
	HasPCClass         bool   `json:"has_pc_class,omitempty"`

	// Unused fields and their column positions
	//
	// Race                  string // 3
	// MonsterSource         string // 5
	// HP                    int    // 15
	// Fort                  string // 19
	// Ref                   string // 20
	// Will                  string // 21
	// SaveMods              string // 22
	// AbilityScoreMods      string // 41
	// RacialMods            string // 47
	// Group                 string // 54
	// IsTemplate            bool   // 56
	// FullHTMLText          string // 59
	// Gender                string // 60
	// Bloodline             string // 61
	// ProhibitedSchools     string // 62
	// BeforeCombat          string // 63
	// DuringCombat          string // 64
	// Morale                string // 65
	// Vulnerability         string // 68
	// Note                  string // 69
	// CharacterFlag         bool   // 70
	// CompanionFlag         bool   // 71
	// Fly                   bool   // 72
	// Climb                 bool   // 73
	// Burrow                bool   // 74
	// Swim                  bool   // 75
	// Land                  bool   // 76
	// TemplatesApplied      string // 77
	// OffenseNote           string // 78
	// BaseStatistics        string // 79
	// ExtractsPrepared      string // 80
	// AgeCategory           string // 81
	// DontUseRacialHD       bool   // 82
	// VariantParent         string // 83
	// Mystery               string // 84
	// ClassArchetypes       string // 85
	// Patron                string // 86
	// CompanionFamiliarLink string // 87
	// FocusedSchool         string // 88
	// Traits                string // 89
	// AlternateNameForm     string // 90
	// StatisticsNote        string // 91
	// LinkText              string // 92
	// UniqueMonster         bool   // 94
	// Mythic                bool   // 96
	// MT                    bool   // 97
}

// LabeledField holds a label & field pair.
type LabeledField struct {
	Label string
	Field string
}

var saRegex = regexp.MustCompile(`^\s*(.* \([a-zA-Z]{2}\))\s+(.*)$`)

// ExtractSpecialAbilities extracts the special abilities from the entity.
func (e *Entity) ExtractSpecialAbilities() []*LabeledField {
	var result []*LabeledField
	var field *LabeledField
	for _, part := range strings.Split(e.SpecialAbilities, ".") {
		matches := saRegex.FindAllStringSubmatch(part, -1)
		if len(matches) > 0 {
			field = &LabeledField{
				Label: matches[0][1],
				Field: strings.TrimSpace(matches[0][2]) + ".",
			}
			result = append(result, field)
		} else {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			if field == nil {
				field = &LabeledField{}
				result = append(result, field)
			}
			if field.Field != "" {
				field.Field += " "
			}
			field.Field += part + "."
		}
	}
	return result
}

// ExtractInitiativeBase returns the base initiative.
func (e *Entity) ExtractInitiativeBase() int {
	var buffer strings.Builder
	leading := true
	for _, ch := range e.Initiative {
		if leading {
			if ch == ' ' {
				continue
			}
			if ch == '-' || ch >= '0' && ch <= '9' {
				buffer.WriteRune(ch)
				leading = false
			}
		} else {
			if ch >= '0' && ch <= '9' {
				buffer.WriteRune(ch)
			} else {
				break
			}
		}
	}
	init, err := strconv.Atoi(buffer.String())
	if err != nil {
		return 0
	}
	return init
}
