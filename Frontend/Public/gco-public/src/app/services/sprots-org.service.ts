import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { RegistrationDto } from '../model/dto/registration-dto';
import { Judge } from '../model/core/judge'
import { SportsOrg } from '../model/core/sports-org';
import { Observable } from 'rxjs';
import { Contestant } from '../model/core/contestant';

@Injectable({
  providedIn: 'root'
})
export class SprotsOrgService {

  path: string = environment.applicationPath + "/sports-organisation";


  constructor(
    private readonly http: HttpClient
    ) { }

  register = (dto: RegistrationDto): Observable<string> => {
    return this.http.post<string>(this.path, dto);
  }

  getLoggedIn = (): Observable<SportsOrg> => {
    return this.http.get<SportsOrg>(this.path);
  }

  getJudges = (): Observable<Judge[]> => {
    return this.http.get<Judge[]>(this.path + "/judge");
  }

  registerJudge = (judge: Judge): Observable<string> => {
    return this.http.post<string>(this.path + "/judge", judge);
  }

  getContestants = (): Observable<Contestant[]> => {
    return this.http.get<Contestant[]>(this.path + "/contestant");
  }

  registerContestant = (contestant: Contestant): Observable<string> => {
    return this.http.post<string>(this.path + "/contestant", contestant);
  }
}
