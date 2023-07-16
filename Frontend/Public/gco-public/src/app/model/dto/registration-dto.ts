import { SportsOrg } from "../core/SportsOrg";
import { Account } from "../core/account";

export interface RegistrationDto{
    account : Account;
    sportsOrganisation : SportsOrg;
}