import { DelegationMember } from "./delegation-member";
import { SportsOrg } from "./sports-org"
import { LicenceType } from "./licence-type"

export interface Judge{
    delegationMember : DelegationMember;
    licenceType : LicenceType;
    licenceName : string;
}

export interface JudgeTable {
    displayedColumns: string[];
    dataSource: Judge[];
    selectedRow: Judge | null;
}