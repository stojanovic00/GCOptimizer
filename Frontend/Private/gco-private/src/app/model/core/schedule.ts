import { ScheduleSlot} from './schedule-slot'

export interface Schedule{
    id : string;
    slots: ScheduleSlot[];
    startingTimes: number[];
}