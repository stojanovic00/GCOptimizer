import { Injectable } from '@angular/core';
import { environment } from 'src/environments/environment';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { RegistrationDto } from '../model/dto/registration-dto';

@Injectable({
  providedIn: 'root'
})
export class SprotsOrgService {

  path: string = environment.applicationPath + "/sports-organisation";


  constructor(private readonly http: HttpClient) { }

  register= (dto: RegistrationDto) => {
    return this.http.post(this.path, dto);
  }
}
