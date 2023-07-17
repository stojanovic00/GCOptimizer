import { Competition } from "./competition";

export interface AgeCategory{
    id? : string;
    name : string;
    minAge : number;
    maxAge : number;
    competition? : Competition;
}