import { ApparatusContestantList } from "./apparatus-contestant-list";

export interface ScheduleSessionView{
    session: number;
    startTime: string;
    apparatusLists: ApparatusContestantList[];
}