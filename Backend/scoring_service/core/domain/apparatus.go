package domain

import (
	"encoding/json"
	"fmt"
)

type Apparatus int8

const (
	Floor Apparatus = iota
	PommelHorse
	StillRings
	Vault
	ParallelBars
	HorizontalBar
	BalanceBeam
	UnevenBars
)

type Apparatuses []Apparatus

func (a *Apparatuses) scanApparatuses(src interface{}) error {
	switch src := src.(type) {
	case Apparatuses:
		jsonData, err := json.Marshal(src)
		if err != nil {
			return err
		}
		return json.Unmarshal(jsonData, a)
	default:
		return fmt.Errorf("unsupported type for Apparatuses: %T", src)
	}
}
