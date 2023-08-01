import { Apparatus } from './apparatus'
import { AllAroundScoreboardSlot} from './all-around-scoreboard-slot'

export interface AllAroundScoreboard{
    id: string;
    competitionId: string;
    ageCategory: string;
    tieBrake: boolean;
    apparatuses: Apparatus[];
    slots: AllAroundScoreboardSlot[];
}


