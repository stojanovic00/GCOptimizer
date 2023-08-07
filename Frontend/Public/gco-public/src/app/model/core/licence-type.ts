export enum LicenceType {
    National = 0,
    International = 1,
}

export function getLicenceTypeName(value: LicenceType): string {
    return LicenceType[value] || LicenceType[0];
}
