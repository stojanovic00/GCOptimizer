import { AgeCategory } from './age-category'
import { Apparatus } from './apparatus'
import { TeamScoreboardSlot } from './team-scoreboard-slot';

export interface TeamScoreboard{
    id: string;
    competitionId: string;
    ageCategory: string;
    apparatuses: Apparatus[];
    slots: TeamScoreboardSlot[];
}


