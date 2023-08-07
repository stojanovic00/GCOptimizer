export enum CompetitionType{
    TeamFinals,
    AllAroundFinals,
    ApparatusFinals,
    Qualifications,
  }
  
export function getCompetitionTypeName(value: CompetitionType): string {
    return CompetitionType[value] || CompetitionType[0];
}