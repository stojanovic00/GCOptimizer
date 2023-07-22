import { ScheduleSlot } from "../core/schedule-slot";
import { ScheduleSessionView} from "../dto/schedule-session-view";
import { ScheduleDto } from "../dto/schedule-dto";
import { unixTimeStampToDate } from "src/app/utils/date-utils";
import { Apparatus, getApparatusName } from "../core/apparatus";
import { ApparatusContestantList} from '../dto/apparatus-contestant-list';

export function ScheduleDtoToScheduleView(dto: ScheduleDto):ScheduleSessionView[] {
    let slots : ScheduleSlot[] = dto.schedule.slots;
    let startingTimes = dto.startingTimes;

    let scheduleSessionViews : ScheduleSessionView[] = [];


    let slotsBySession: ScheduleSlot[][] = groupBySession(slots);

    //Iterate over sessions
    for(const sessionSlots of slotsBySession){
        const session : number = sessionSlots[0].session; 
        const startTimeDate : Date = unixTimeStampToDate(startingTimes[session-1]);  
        let startTime : string = startTimeDate.getHours().toString().padStart(2,'0') + ":" + startTimeDate.getMinutes().toString().padStart(2,'0');

        let slotsByStartingApparatus: ScheduleSlot[][] = groupByStartingApparatus(sessionSlots);



        let apparatusContestantLists : ApparatusContestantList[] = []
        for(const startingAppSlotList of slotsByStartingApparatus){
            const apparatusName = getApparatusName(startingAppSlotList[0].startingApparatus)
            apparatusContestantLists.push({
                apparatusName: apparatusName,
                slots: startingAppSlotList
            })
        }

        scheduleSessionViews.push({
            session: session,
            startTime: startTime,
            apparatusLists: apparatusContestantLists
        })
    }


    return scheduleSessionViews;
}



function groupByStartingApparatus(sessionSlots: ScheduleSlot[]) {
    let groupedSlots: { [apparatus: string]: ScheduleSlot[]; } = sessionSlots.reduce((acc, slot) => {
        const apparatusKey = Apparatus[slot.startingApparatus]; // Convert the enum to a string key
        if (!acc[apparatusKey]) {
            acc[apparatusKey] = [];
        }
        acc[apparatusKey].push(slot);
        return acc;
    }, {} as { [apparatus: string]: ScheduleSlot[]; });

    // Convert the groupedSlots object into an array of arrays
    let slotsByStartingApparatus: ScheduleSlot[][] = Object.values(groupedSlots);
    return slotsByStartingApparatus;
}

function groupBySession(slots: ScheduleSlot[]) {
    let groupedSlots: { [session: number]: ScheduleSlot[]; } = slots.reduce((acc, slot) => {
        if (!acc[slot.session]) {
            acc[slot.session] = [];
        }
        acc[slot.session].push(slot);
        return acc;
    }, {} as { [session: number]: ScheduleSlot[]; }); // Explicitly define the type here


    // Convert the groupedSlots object into an array of arrays
    let slotsBySession: ScheduleSlot[][] = Object.values(groupedSlots);
    return slotsBySession;
}
