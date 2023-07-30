import { IdMessage } from "../dto/id-message";

export interface Score{
    id?: string;
    dScore: number;
    eScore: number;
    totalScore: number;
    competitionId: string;
    contestant: IdMessage
}
