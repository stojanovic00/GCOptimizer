package dto

type CurrentSessionInfo struct {
	CurrentRotation     int32
	CurrentSession      int32
	RotationFinished    bool
	SessionFinished     bool
	CompetitionFinished bool
}
