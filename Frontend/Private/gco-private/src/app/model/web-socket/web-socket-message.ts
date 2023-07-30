import { ScoringEvent } from '../web-socket/scoring-event'
import { Apparatus } from '../core/apparatus'

export interface WebSocketMessage{
	event: ScoringEvent
	apparatus:     Apparatus
	competitionId: string
}