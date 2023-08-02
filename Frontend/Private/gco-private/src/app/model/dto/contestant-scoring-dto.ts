import { ContestantScoring } from "../core/contestant-scoring";

export interface ContestantScoringDto{
    contestant: ContestantScoring;
    competes: boolean;
}

export interface ContestantScoringDtoTable {
    displayedColumns: string[];
    dataSource: ContestantScoringDto[];
    selectedRow: ContestantScoringDto | null;
}