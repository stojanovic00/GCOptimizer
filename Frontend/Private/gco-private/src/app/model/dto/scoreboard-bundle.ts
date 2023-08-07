import { AllAroundScoreboard } from '../core/all-around-scoreboard'
import { TeamScoreboard } from '../core/team-scoreboard'

export interface ScoreboardBundle{
    allAroundScoreboards: AllAroundScoreboard[];
    teamScoreboards: TeamScoreboard[];
}
