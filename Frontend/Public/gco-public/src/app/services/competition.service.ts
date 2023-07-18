import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Competition } from '../model/core/competition';
import { Observable } from 'rxjs';
import { DelegationMemberProposition } from '../model/core/delegation-member-proposition';
import { AgeCategory } from '../model/core/age-category';
import { IdResponse } from '../model/dto/id-response';
import { CreateJudgeApplicationDto } from '../model/dto/create-judge-application-dto';
import { CreateContestantApplicationRequest } from '../model/dto/create-contestant-application-request';

@Injectable({
  providedIn: 'root'
})
export class CompetitionService {

  path: string = environment.applicationPath + "/competition";

  constructor(
    private readonly http: HttpClient
    ) { }


  create = (competition: Competition): Observable<IdResponse> => {
    return this.http.post<IdResponse>(this.path, competition);
  }

  addDelegationMemberProposition = (prop: DelegationMemberProposition, competitionId: string): Observable<string> => {
    let path = this.path + "/"+  competitionId + "/delegation-member-prop";
    return this.http.post<string>(path, prop);
  }

  addAgeCategory = (cat: AgeCategory, competitionId: string): Observable<string> => {
    let path = this.path + "/"+  competitionId + "/age-category";
    return this.http.post<string>(path, cat);
  }

  getAll = (): Observable<Competition[]> => {
    return this.http.get<Competition[]>(this.path);
  }

  getById = (competitionId: string): Observable<Competition> => {
    let path = this.path + "/"+  competitionId;
    return this.http.get<Competition>(path);
  }

  addJudgeApplication = (app: CreateJudgeApplicationDto, competitionId: string): Observable<string> => {
    let path = this.path + "/"+  competitionId + "/app/judge";
    return this.http.post<string>(path, app);
  }

  addContestantApplication = (app: CreateContestantApplicationRequest, competitionId: string): Observable<string> => {
    let path = this.path + "/"+  competitionId + "/app/contestant";
    return this.http.post<string>(path, app);
  }
}
