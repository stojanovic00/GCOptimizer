package domain

import (
	"github.com/google/uuid"
)

type Panel struct {
	ID                       uuid.UUID
	Type                     JudgingPanelType
	Apparatus                Apparatus
	CompetitionID            uuid.UUID
	Competition              Competition
	ScoreCalculationMethodID *uuid.UUID `gorm:"default:null"`
	ScoreCalculationMethod   ScoreCalculationMethod
	Judges                   []Judge `gorm:"many2many:panel_judges;"`
}
