import { Competition } from './competition'
import { Judge } from './judge'

export interface JudgeApplication{
    id? : string;
    competition : Competition;
    judge : Judge;
  }
  
export interface JudgeApplicationTable {
    displayedColumns: string[];
    dataSource: JudgeApplication[];
    selectedRow: JudgeApplication | null;
}