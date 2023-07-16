package domain

import (
	"time"
)

type Contestant struct {
	DelegationMember
	DateOfBirth time.Time
}
