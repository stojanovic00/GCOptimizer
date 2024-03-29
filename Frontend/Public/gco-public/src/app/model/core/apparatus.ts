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

export function getApparatusName(value: Apparatus): string {
    return Apparatus[value] || Apparatus[0];
}

