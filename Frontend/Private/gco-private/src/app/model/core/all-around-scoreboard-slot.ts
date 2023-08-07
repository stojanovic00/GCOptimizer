import { Apparatus } from './apparatus';
import { ContestantScoring } from './contestant-scoring'
import { Score } from './score'


export interface AllAroundScoreboardSlot{
    id: string;
    place: number;
    contestant: ContestantScoring;
    scores: Score[];
    totalScore: number;
    apparatusScore: Map<Apparatus, Score> //Only for frontend
}

export function mapAllAroundScores(scores: Score[]):Map<Apparatus, Score>{
    let scoreMap: Map<Apparatus, Score> = new Map();
    for (const score of scores) {
      scoreMap.set(score.apparatus ?? Apparatus.Floor, score);
    }

    return scoreMap;
}


