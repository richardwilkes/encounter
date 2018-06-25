package data

import (
	"regexp"
	"strings"
)

// Entity holds the static information for a combatant.
type Entity struct {
	Name               string // 0
	CR                 string // 1
	XP                 int    // 2
	Alignment          string // 6
	Size               string // 7
	Type               string // 8
	SubType            string // 9
	Init               string // 10
	Senses             string // 11
	Aura               string // 12
	AC                 string // 13
	ACMods             string // 14
	HD                 string // 16
	HPMods             string // 17
	Saves              string // 18
	DefensiveAbilities string // 23
	DR                 string // 24
	Immune             string // 25
	Resist             string // 26
	SR                 string // 27
	Weaknesses         string // 28
	Speed              string // 29
	SpeedMod           string // 30
	Melee              string // 31
	Ranged             string // 32
	Space              string // 33
	Reach              string // 34
	SpecialAttacks     string // 35
	SpellLikeAbilities string // 36
	SpellsKnown        string // 37
	SpellsPrepared     string // 38
	SpellDomains       string // 39
	AbilityScores      string // 40
	BaseAttack         int    // 42
	CMB                string // 43
	CMD                string // 44
	Feats              string // 45
	Skills             string // 46
	Languages          string // 48
	SQ                 string // 49
	Environment        string // 50
	Organization       string // 51
	Treasure           string // 52
	DescriptionVisual  string // 53
	Source             string // 55
	SpecialAbilities   string // 57
	Description        string // 58
	Gear               string // 66
	OtherGear          string // 67
	MonsterID          int    // 93
	MR                 int    // 95
	HasPCClass         bool

	// Unused fields and their column positions
	//
	// Race                  string // 3
	// Class                 string // 4
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
			if len(part) == 0 {
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
