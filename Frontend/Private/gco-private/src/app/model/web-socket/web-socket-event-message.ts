import { ScoringEvent } from './scoring-event'
import { Apparatus } from '../core/apparatus'

export interface WebSocketEventMessage{
	event: ScoringEvent
	competitionId: string
	apparatus:     Apparatus
	ContestantId: string
}