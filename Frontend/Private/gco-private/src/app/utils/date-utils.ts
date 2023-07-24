//Golangs time.Time expects unix time stamp in seconds
export function unixTimeStampToDate(unixTimeStamp: number): Date {
    return new Date(unixTimeStamp * 1000);
}

export function dateToUnixTimeStamp(date: Date): number {
    return Math.floor(date.getTime() / 1000);
}

export function formatDate(date: Date): string {
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const year = date.getFullYear().toString();
    
    return `${day}.${month}.${year}.`;
  }

export function unixDateToString(unixDate: number){
    return formatDate(unixTimeStampToDate(unixDate));
}

export function timeStringToDate(time: string) : Date{
    const [hours, minutes] = time.split(":").map(Number);

    let date = new Date(2001, 0, 1);

    date.setHours(hours);
    date.setMinutes(minutes);
    return date;
}