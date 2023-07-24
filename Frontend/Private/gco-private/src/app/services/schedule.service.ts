import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { environment } from '../../environments/environment';
import { SchedulingParameters } from '../model/core/scheduling-paramters';
import { Observable } from 'rxjs';
import { Schedule } from '../model/core/schedule';


@Injectable({
  providedIn: 'root'
})
export class ScheduleService {
  path: string = environment.schedulingPath + "/schedule";

  constructor(
    private readonly http: HttpClient
    ) { }

  generateSchedule = (params: SchedulingParameters): Observable<Schedule> => {
    return this.http.post<Schedule>(this.path, params);
  }

  getByCompetitionId = (id: string): Observable<Schedule> => {
    let path = this.path + "/" + id;
    return this.http.get<Schedule>(path);
  }
}
