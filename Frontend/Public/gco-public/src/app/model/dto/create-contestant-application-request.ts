import { ApparatusAnnouncement } from "../core/apparatus-announcement";

export interface CreateContestantApplicationRequest{
    competitionId : string;
    teamNumber : number;
    contestantId : string;
    ageCategoryId : string;
    apparatusAnnouncements : ApparatusAnnouncement[];
  }
  