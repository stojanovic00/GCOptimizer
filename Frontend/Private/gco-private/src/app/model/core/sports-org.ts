import { Address } from "./address";

export interface SportsOrg{
    name: string;
    email: string;
    phoneNumber: string;
    contactPersonFullName: string;
    address: Address;
    competitionOrganisingPrivilege: boolean;
}