import { Competition } from "./competition";

export interface AgeCategory{
    id? : string;
    name : string;
    minAge : number;
    maxAge : number;
    competition? : Competition;
}

export function formatAgeCategory(cat : AgeCategory) :string{
    const minAge = cat.minAge ?? 0
    const maxAge = cat.maxAge ?? 0
    return cat.name + "(" + minAge + ", " +maxAge + ")";
}
export interface AgeCategoryTable {
    displayedColumns: string[];
    dataSource: AgeCategory[];
    selectedRow: AgeCategory | null;
}