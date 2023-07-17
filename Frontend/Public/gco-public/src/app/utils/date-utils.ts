//Golangs time.Time expects unix time stamp in seconds
export function unixTimeStampToDate(unixTimeStamp: number): Date {
    return new Date(unixTimeStamp * 1000);
}

export function dateToUnixTimeStamp(date: Date): number {
    return date.getTime() / 1000;
}

export function formatDate(date: Date): string {
    const day = date.getDate().toString().padStart(2, '0');
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const year = date.getFullYear().toString();
    
    return `${day}.${month}.${year}`;
  }