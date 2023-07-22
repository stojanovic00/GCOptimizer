import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient } from '@angular/common/http';
import { SportsOrg } from '../model/core/sports-org';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SprotsOrgService {

  path: string = environment.applicationPath + "/sports-organisation";


  constructor(
    private readonly http: HttpClient
    ) { }


  getLoggedIn = (): Observable<SportsOrg> => {
    return this.http.get<SportsOrg>(this.path);
  }

}
