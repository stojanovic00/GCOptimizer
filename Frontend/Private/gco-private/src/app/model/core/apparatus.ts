export enum Apparatus{
    Floor = 0,
    PommelHorse = 1,
    StillRings = 2,
    Vault = 3,
    ParallelBars = 4,
    HorizontalBar = 5,
    BalanceBeam = 6,
    UnevenBars = 7,
}

export interface ApparatusTable {
    displayedColumns: string[];
    dataSource: Apparatus[];
    selectedRow: Apparatus | null;
}

export function getApparatusName(value: Apparatus): string {
    return Apparatus[value] || Apparatus[0];
}

export function parseApparatus(value: string): Apparatus | undefined {
    // Convert the string to the enum value
    return (Apparatus as any)[value] as Apparatus | undefined;
  }