package domain

import (
	"github.com/google/uuid"
)

type Panel struct {
	ID                       uuid.UUID
	Type                     PanelType
	Apparatus                Apparatus
	CompetitionID            uuid.UUID
	Competition              Competition
	ScoreCalculationMethodID uuid.UUID
	ScoreCalculationMethod   ScoreCalculationMethod
	Judges                   []Judge `gorm:"many2many:panel_judges;"`
}
