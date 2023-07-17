import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { RegistrationDto } from '../model/dto/registration-dto';
import { Judge } from '../model/core/judge'
import { SportsOrg } from '../model/core/sports-org';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SprotsOrgService {

  path: string = environment.applicationPath + "/sports-organisation";


  constructor(private readonly http: HttpClient) { }

  register = (dto: RegistrationDto) => {
    return this.http.post(this.path, dto);
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
}
