package board

// Entity holds the static information for a combatant.
type Entity struct {
	Name  string `json:",omitempty"` // 0
	CR    string `json:",omitempty"` // 1
	XP    int    `json:",omitempty"` // 2
	Race  string `json:",omitempty"` // 3
	Class string `json:",omitempty"` // 4
	// MonsterSource         string  `json:",omitempty"` // 5
	Alignment string `json:",omitempty"` // 6
	Size      string `json:",omitempty"` // 7
	Type      string `json:",omitempty"` // 8
	SubType   string `json:",omitempty"` // 9
	Init      string `json:",omitempty"` // 10
	Senses    string `json:",omitempty"` // 11
	Aura      string `json:",omitempty"` // 12
	AC        string `json:",omitempty"` // 13
	ACMods    string `json:",omitempty"` // 14
	HP        int    `json:",omitempty"` // 15
	HD        string `json:",omitempty"` // 16
	HPMods    string `json:",omitempty"` // 17
	Saves     string `json:",omitempty"` // 18
	// Fort                  string `json:",omitempty"` // 19
	// Ref                   string `json:",omitempty"` // 20
	// Will                  string `json:",omitempty"` // 21
	// SaveMods              string `json:",omitempty"` // 22
	DefensiveAbilities string `json:",omitempty"` // 23
	DR                 string `json:",omitempty"` // 24
	Immune             string `json:",omitempty"` // 25
	Resist             string `json:",omitempty"` // 26
	SR                 string `json:",omitempty"` // 27
	Weaknesses         string `json:",omitempty"` // 28
	Speed              string `json:",omitempty"` // 29
	SpeedMod           string `json:",omitempty"` // 30
	Melee              string `json:",omitempty"` // 31
	Ranged             string `json:",omitempty"` // 32
	Space              string `json:",omitempty"` // 33
	Reach              string `json:",omitempty"` // 34
	SpecialAttacks     string `json:",omitempty"` // 35
	SpellLikeAbilities string `json:",omitempty"` // 36
	SpellsKnown        string `json:",omitempty"` // 37
	SpellsPrepared     string `json:",omitempty"` // 38
	SpellDomains       string `json:",omitempty"` // 39
	AbilityScores      string `json:",omitempty"` // 40
	// AbilityScoreMods      string `json:",omitempty"` // 41
	BaseAttack        int    `json:",omitempty"` // 42
	CMB               string `json:",omitempty"` // 43
	CMD               string `json:",omitempty"` // 44
	Feats             string `json:",omitempty"` // 45
	Skills            string `json:",omitempty"` // 46
	RacialMods        string `json:",omitempty"` // 47
	Languages         string `json:",omitempty"` // 48
	SQ                string `json:",omitempty"` // 49
	Environment       string `json:",omitempty"` // 50
	Organization      string `json:",omitempty"` // 51
	Treasure          string `json:",omitempty"` // 52
	DescriptionVisual string `json:",omitempty"` // 53
	Group             string `json:",omitempty"` // 54
	Source            string `json:",omitempty"` // 55
	IsTemplate        bool   `json:",omitempty"` // 56
	SpecialAbilities  string `json:",omitempty"` // 57
	Description       string `json:",omitempty"` // 58
	// FullHTMLText          string `json:",omitempty"` // 59
	Gender            string `json:",omitempty"` // 60
	Bloodline         string `json:",omitempty"` // 61
	ProhibitedSchools string `json:",omitempty"` // 62
	BeforeCombat      string `json:",omitempty"` // 63
	DuringCombat      string `json:",omitempty"` // 64
	Morale            string `json:",omitempty"` // 65
	Gear              string `json:",omitempty"` // 66
	OtherGear         string `json:",omitempty"` // 67
	Vulnerability     string `json:",omitempty"` // 68
	Note              string `json:",omitempty"` // 69
	CharacterFlag     bool   `json:",omitempty"` // 70
	CompanionFlag     bool   `json:",omitempty"` // 71
	// Fly                   bool `json:",omitempty"`   // 72
	// Climb                 bool `json:",omitempty"`   // 73
	// Burrow                bool `json:",omitempty"`   // 74
	// Swim                  bool `json:",omitempty"`   // 75
	// Land                  bool `json:",omitempty"`   // 76
	TemplatesApplied      string `json:",omitempty"` // 77
	OffenseNote           string `json:",omitempty"` // 78
	BaseStatistics        string `json:",omitempty"` // 79
	ExtractsPrepared      string `json:",omitempty"` // 80
	AgeCategory           string `json:",omitempty"` // 81
	DontUseRacialHD       bool   `json:",omitempty"` // 82
	VariantParent         string `json:",omitempty"` // 83
	Mystery               string `json:",omitempty"` // 84
	ClassArchetypes       string `json:",omitempty"` // 85
	Patron                string `json:",omitempty"` // 86
	CompanionFamiliarLink string `json:",omitempty"` // 87
	FocusedSchool         string `json:",omitempty"` // 88
	Traits                string `json:",omitempty"` // 89
	AlternateNameForm     string `json:",omitempty"` // 90
	StatisticsNote        string `json:",omitempty"` // 91
	LinkText              string `json:",omitempty"` // 92
	MonsterID             int    `json:",omitempty"` // 93
	UniqueMonster         bool   `json:",omitempty"` // 94
	MR                    int    `json:",omitempty"` // 95
	Mythic                bool   `json:",omitempty"` // 96
	MT                    bool   `json:",omitempty"` // 97
}
