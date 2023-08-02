import { ScoreType } from './score-type'
import { Apparatus } from './apparatus'
import { ContestantScoring } from './contestant-scoring'
import { JudgeBasicInfo } from '../dto/judge-basic-info'

export interface TempScore{
    id? : string;
    type : ScoreType;
    apparatus : Apparatus;
    value : number;
    contestant : ContestantScoring;
    competitionId: string;
    judge : JudgeBasicInfo;
}


export interface TempScoreTable {
    displayedColumns: string[];
    dataSource: TempScore[];
    selectedRow: TempScore | null;
}
