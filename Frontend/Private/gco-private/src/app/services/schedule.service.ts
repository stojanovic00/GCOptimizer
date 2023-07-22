import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { SchedulingParameters } from '../model/core/scheduling-paramters';
import { ScheduleDto } from '../model/dto/schedule-dto';
import { Observable } from 'rxjs';


@Injectable({
  providedIn: 'root'
})
export class ScheduleService {
  path: string = environment.schedulingPath + "/schedule";

  constructor(
    private readonly http: HttpClient
    ) { }

  generateSchedule = (params: SchedulingParameters): Observable<ScheduleDto> => {
    return this.http.post<ScheduleDto>(this.path, params);
  }

}
