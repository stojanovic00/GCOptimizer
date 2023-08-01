import { IdMessage } from "../dto/id-message";
import { Apparatus } from "./apparatus";

export interface Score{
    id?: string;
    dScore: number;
    eScore: number;
    totalScore: number;
    apparatus: Apparatus
    competitionId: string;
    contestant: IdMessage;
    submitted: boolean;
}
