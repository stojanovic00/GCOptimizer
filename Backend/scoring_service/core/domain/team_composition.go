package domain

import "github.com/google/uuid"

type TeamComposition struct {
	ID                    uuid.UUID
	BaseContestantNumber  int
	BonusContestantNumber int
	MultiCategoryTeam     bool
}
