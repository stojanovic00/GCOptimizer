package domain

import "github.com/google/uuid"

type ScoreCalculationMethod struct {
	ID uuid.UUID
	//How many scores from top and bottom will be deduced
	ScoreDeductionNum int32
}
