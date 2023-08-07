import { Judge } from '../core/judge';
import { LicenceType } from '../core/licence-type'
import { SportsOrg } from '../core/sports-org'


export interface JudgeBasicInfo{
    id: string;
    fullName: string;
    email: string;
    licenceType: LicenceType 
    licenceName: string;
    sportsOrganization: SportsOrg
}

export interface JudgeBasicInfoTable {
    displayedColumns: string[];
    dataSource: JudgeBasicInfo[];
    selectedRow: JudgeBasicInfo | null;
}

export function parseJudgeBasicInfo(judge: Judge): JudgeBasicInfo{
    return {
        id: judge.delegationMember.id!,
        fullName: judge.delegationMember.fullName!,
        email: judge.delegationMember.email!,
        licenceType: judge.licenceType!,
        licenceName: judge.licenceName!,
        sportsOrganization: judge.delegationMember.sportsOrganisation!
    }
}