export interface ScoringEvent{
    type : EventType; 
    competitionId: string;
    apparatus: Apparatus;
}

export enum EventType{
  SCORE_SAVED = 0,
  .
  .
  .
}

export enum Apparatus{
  Floor = 0,
  .
  .
  .
}
