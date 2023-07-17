import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { Competition } from '../model/core/competition';
import { Observable } from 'rxjs';
import { DelegationMemberProposition } from '../model/core/delegation-member-proposition';
import { AgeCategory } from '../model/core/age-category';
import { IdResponse } from '../model/dto/id-response';

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
}
