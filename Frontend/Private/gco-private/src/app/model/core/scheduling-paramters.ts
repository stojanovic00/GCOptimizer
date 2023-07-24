import { ApparatusTypeFoScheduling } from './apparatus-type-for-scheduling'

export interface SchedulingParameters{
    competitionId: string;
    startTime: number;
    endTime: number;
    warmupRoomAvailable: boolean;
    generalWarmupTime: number;
    warmupTime: number;
    warmupsPerApparatus: number;
    contestantNumPerApparatus: number;
    executionTime: number;
    apparatusRotationTime: number;
    medalCeremonyAfterOneSessionTime: number;
    finalMedalCeremonyTime: number;
    halfApparatusPerSessionMode: boolean;
    apparatusOrder: ApparatusTypeFoScheduling[];
}
