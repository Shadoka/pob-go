package model

type PassiveTree struct {
	Tree string
}

type PlayerClass struct {
	Name         string
	Base_str     uint16
	Base_dex     uint16
	Base_int     uint16
	Ascendancies []Ascendancy
}

type Ascendancy struct {
	Id                string
	Name              string
	FlavourText       string
	FlavourTextColour string
	FlavourTextRect   Rectangle
}

type Rectangle struct {
	X      uint16
	Y      uint16
	Width  uint16
	Height uint16
}
