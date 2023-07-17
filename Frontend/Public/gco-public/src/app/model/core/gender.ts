export enum Gender{
    Male = 0,
    Female = 1,
}

export function getGenderName(value: Gender): string {
    return Gender[value] || Gender[0];
}

