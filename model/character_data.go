package model

const COBALT_JEWEL = "JewelInt"
const CRIMSON_JEWEL = "JewelStr"
const VIRIDIAN_JEWEL = "JewelDex"
const PRISMATIC_JEWEL = "JewelPrismatic"
const TIMELESS_JEWEL = "JewelTimeless"

const MARAKETH_FACTION = "Maraketh"

type Jewel struct {
	Type         string
	Subgraph     JewelSubgraph
	Radius       uint32
	RadiusVisual string
}

type JewelSubgraph struct {
	// NOTE: "groups" has always a dynamic key: expansion_<ID_OF_JEWEL>
	// the value of the expansion_<ID_OF_JEWEL> is JewelGroupsExpansion
	Groups map[string]JewelGroupsExpansion
	Nodes  map[string]JewelNode
}

type JewelNode struct {
	Skill         string
	Name          string
	Icon          string
	IsNotable     bool
	IsKeystone    bool
	IsJewelSocket bool
	IsMastery     bool
	Stats         []string
	ReminderText  []string
	FlavourText   []string
	// group relates to the dynamic expansion name, see JewelSubgraph
	Group          string
	Orbit          uint8
	OrbitIndex     uint8
	Out            []string
	In             []string
	ExpansionJewel SocketedJewelLink
}

type SocketedJewelLink struct {
	Size  uint8
	Index uint8
	// TODO: to what relate proxy and parent?
	Proxy  string
	Parent string
}

type JewelGroupsExpansion struct {
	Proxy  string
	Nodes  []string
	X      float64
	Y      float64
	Orbits []uint8
}

type CharacterData struct {
	Hashes          []uint32
	Hashes_ex       []uint32
	Mastery_effects []string
	Items           []Item
}

type Item struct {
	Verified     bool
	W            uint
	H            uint
	Icon         string
	League       string
	Id           string
	Name         string
	TypeLine     string
	BaseType     string
	Identified   bool
	Ilvl         int8
	Corrupted    bool
	DescrText    string
	FrameType    uint8
	X            uint32
	Y            uint32
	InventoryId  string
	FlavourText  []string
	ImplicitMods []string
	ExplicitMods []string
	EnchantMods  []string
	Properties   []ItemProperty
	Requirements []ItemRequirement
}

type ItemProperty struct {
	Name        string
	Values      []interface{}
	DisplayMode uint32
	Type        uint32
}

type ItemRequirement struct {
	Name        string
	Values      []interface{}
	DisplayMode uint32
	Type        uint32
}
