import {Apparatus, getApparatusName} from './apparatus'

export interface ApparatusAnnouncement{
    id? : string; 
    apparatus : Apparatus;
}

export function formatApparatusAnnouncements(announcements : ApparatusAnnouncement[]){
    const apparatusNames = announcements.map((announcement) => getApparatusName(announcement.apparatus)).sort();
    return apparatusNames.join(', ');
}
  
export interface ApparatusAnnouncementTable {
    displayedColumns: string[];
    dataSource: ApparatusAnnouncement[];
    selectedRows: ApparatusAnnouncement[];
}