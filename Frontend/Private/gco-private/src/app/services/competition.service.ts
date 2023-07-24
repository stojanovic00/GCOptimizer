import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Competition } from '../model/core/competition';
import { Observable } from 'rxjs';
import { ContestantApplication } from '../model/core/contestant-application';
import { JudgeApplication } from '../model/core/judge-application';

@Injectable({
  providedIn: 'root'
})
export class CompetitionService {

  path: string = environment.applicationPath + "/competition";

  constructor(
    private readonly http: HttpClient
    ) { }



  getAll = (): Observable<Competition[]> => {
    return this.http.get<Competition[]>(this.path);
  }

  getById = (competitionId: string): Observable<Competition> => {
    let path = this.path + "/"+  competitionId;
    return this.http.get<Competition>(path);
  }


  getContestantApplications = (competitionId: string): Observable<ContestantApplication[]> => {
    let path = this.path + "/"+  competitionId + "/app/contestant";
    return this.http.get<ContestantApplication[]>(path);
  }

  getJudgeApplications = (competitionId: string): Observable<JudgeApplication[]> => {
    let path = this.path + "/"+  competitionId + "/app/judge";
    return this.http.get<JudgeApplication[]>(path);
  }
}
