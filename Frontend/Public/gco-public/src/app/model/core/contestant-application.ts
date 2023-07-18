import { AgeCategory } from './age-category';
import { ApparatusAnnouncement } from './apparatus-announcement';
import {Competition} from './competition'
import {Contestant} from './contestant'

export interface ContestantApplication{
    id? : string;
    teamNumber : number;
    competition? : Competition;
    contestant : Contestant;
    ageCategory : AgeCategory;
    apparatusAnnouncements : ApparatusAnnouncement[];
}
  
export interface ContestantApplicationTable {
    displayedColumns: string[];
    dataSource: ContestantApplication[];
    selectedRow: ContestantApplication | null;
}

export interface ContestantApplicationPreview{
    id? : string;
    contestantFullName : string;
    teamNumber : number;
    ageCategory : AgeCategory;
    apparatusAnnouncements : ApparatusAnnouncement[];
  }

export interface ContestantApplicationPreviewTable {
    displayedColumns: string[];
    dataSource: ContestantApplicationPreview[];
    selectedRow: ContestantApplicationPreview | null;
}