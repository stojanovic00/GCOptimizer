import { Apparatus } from './apparatus'
import { ContestantInfo } from './contestant-info'

export interface ScheduleSlot{
     session : number;
     startingApparatus : Apparatus;
     contestantInfo : ContestantInfo;
  }
  