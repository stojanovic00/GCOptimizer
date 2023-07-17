import { SportsOrg } from "../core/sports-org";
import { Account } from "../core/account";

export interface RegistrationDto{
    account : Account;
    sportsOrganisation : SportsOrg;
}