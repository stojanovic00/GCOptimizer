import { Role }from "../core/role";

export interface Account{
    email:string;
    password:string;
    role: Role
}