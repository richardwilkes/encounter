package board

import (
	"strconv"
	"strings"

	"github.com/richardwilkes/rpgtools/dice"
)

// HPMethod is an enum of possible ways to determine HP for an entity.
type HPMethod byte

// Possible HPMethods.
const (
	RolledHPMethod HPMethod = iota
	MinimumHPMethod
	AverageHPMethod
	MaximumHPMethod
)

// Entity holds the static information for a combatant.
type Entity struct {
	Name  string // 0
	CR    string // 1
	XP    int    // 2
	Race  string // 3
	Class string // 4
	// MonsterSource         string  // 5
	Alignment string // 6
	Size      string // 7
	Type      string // 8
	SubType   string // 9
	Init      string // 10
	Senses    string // 11
	Aura      string // 12
	AC        string // 13
	ACMods    string // 14
	// HP        int          // 15
	HD     string // 16
	HPMods string // 17
	Saves  string // 18
	// Fort                  string // 19
	// Ref                   string // 20
	// Will                  string // 21
	// SaveMods              string // 22
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
	// AbilityScoreMods      string // 41
	BaseAttack        int    // 42
	CMB               string // 43
	CMD               string // 44
	Feats             string // 45
	Skills            string // 46
	RacialMods        string // 47
	Languages         string // 48
	SQ                string // 49
	Environment       string // 50
	Organization      string // 51
	Treasure          string // 52
	DescriptionVisual string // 53
	Group             string // 54
	Source            string // 55
	IsTemplate        bool   // 56
	SpecialAbilities  string // 57
	Description       string // 58
	// FullHTMLText          string // 59
	Gender                string // 60
	Bloodline             string // 61
	ProhibitedSchools     string // 62
	BeforeCombat          string // 63
	DuringCombat          string // 64
	Morale                string // 65
	Gear                  string // 66
	OtherGear             string // 67
	Vulnerability         string // 68
	Note                  string // 69
	CharacterFlag         bool   // 70
	CompanionFlag         bool   // 71
	Fly                   bool   // 72
	Climb                 bool   // 73
	Burrow                bool   // 74
	Swim                  bool   // 75
	Land                  bool   // 76
	TemplatesApplied      string // 77
	OffenseNote           string // 78
	BaseStatistics        string // 79
	ExtractsPrepared      string // 80
	AgeCategory           string // 81
	DontUseRacialHD       bool   // 82
	VariantParent         string // 83
	Mystery               string // 84
	ClassArchetypes       string // 85
	Patron                string // 86
	CompanionFamiliarLink string // 87
	FocusedSchool         string // 88
	Traits                string // 89
	AlternateNameForm     string // 90
	StatisticsNote        string // 91
	LinkText              string // 92
	MonsterID             int    // 93
	UniqueMonster         bool   // 94
	MR                    int    // 95
	Mythic                bool   // 96
	MT                    bool   // 97
	HasPCClass            bool
}

// SortingName returns the name to sort by.
func (e *Entity) SortingName() string {
	if e.LinkText != "" {
		return e.LinkText
	}
	return e.Name
}

// HP returns the HP using the specified method.
func (e *Entity) HP(method HPMethod) int {
	switch method {
	case MinimumHPMethod:
		return e.MinimumHP()
	case AverageHPMethod:
		return e.AverageHP()
	case MaximumHPMethod:
		return e.MaximumHP()
	default:
		return e.RolledHP()
	}
}

// MinimumHP returns the minimum HP.
func (e *Entity) MinimumHP() int {
	total := 0
	hd := extractDice(e.HD)
	for i, d := range hd {
		if e.HasPCClass && i == len(hd)-1 {
			d.Count--
			total += d.Sides
		}
		if d.Count > 0 {
			total += d.Minimum()
		} else {
			total += d.Modifier
		}
	}
	return total
}

// AverageHP returns the average HP.
func (e *Entity) AverageHP() int {
	total := 0
	hd := extractDice(e.HD)
	for i, d := range hd {
		if e.HasPCClass && i == len(hd)-1 {
			d.Count--
			total += d.Sides
		}
		if d.Count > 0 {
			total += d.Average()
		} else {
			total += d.Modifier
		}
	}
	return total
}

// MaximumHP returns the maximum HP.
func (e *Entity) MaximumHP() int {
	total := 0
	hd := extractDice(e.HD)
	for _, d := range hd {
		total += d.Maximum()
	}
	return total
}

// RolledHP returns the rolled HP.
func (e *Entity) RolledHP() int {
	total := 0
	hd := extractDice(e.HD)
	for _, d := range hd {
		total += d.Roll()
	}
	return total
}

func extractDice(in string) []*dice.Dice {
	in = strings.Replace(in, " plus ", "+", -1)
	d := dice.New(nil, in)
	if d.String() == in || "1"+d.String() == in {
		return []*dice.Dice{d}
	}
	result := make([]*dice.Dice, 0)
	for _, p := range strings.Split(in, "+") {
		d = dice.New(nil, p)
		if d.String() != p && "1"+d.String() != p {
			var buffer strings.Builder
			for _, c := range p {
				if c >= '0' && c <= '9' {
					buffer.WriteRune(c)
				} else {
					break
				}
			}
			v, err := strconv.Atoi(buffer.String())
			if err != nil {
				v = 0
			}
			if len(result) > 0 {
				result[len(result)-1].Modifier += v
			}
		} else {
			result = append(result, d)
		}
	}
	return result
}
