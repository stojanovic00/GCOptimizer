package domain

type CompetitionType int8

const (
	Qualifications CompetitionType = iota
	TeamFinals
	AllAroundFinals
	ApparatusFinals
)
