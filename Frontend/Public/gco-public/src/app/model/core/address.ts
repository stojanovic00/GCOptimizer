export interface Address{
    country:string;
    city:string;
    street:string;
    streetNumber:string;
}


export function formatAddress(addr: Address){
    return addr.country! + " " + addr.city! + ", " + addr.street! + " " + addr.streetNumber!;
}