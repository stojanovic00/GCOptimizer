import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { environment } from 'src/environments/environment';
import { Jwt } from '../model/jwt';
import { LoginDto } from '../model/login-dto';

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor(
    private httpClient: HttpClient,
    private readonly router: Router,
    private readonly jwtHelper: JwtHelperService
  ) { }

  private redirectToMainPage = () => {
    var roleLandingPages = new Map<string, string>([
      ['SPORTS_ORG', 'sports-org/view'],
      ['E_JUDGE', 'judge/info'],
      ['D_JUDGE', 'judge/info'],
    ]);

    const token = localStorage.getItem('jwt');
    const tokenPayload = this.jwtHelper.decodeToken(token!);
    const role = tokenPayload['role'];

    this.router.navigate([roleLandingPages.get(role)]);
  }

  public login = (loginCredentials: LoginDto): void => {
    this.httpClient.post<Jwt>(environment.authPath + '/login', loginCredentials)
      .subscribe({
        next: (response) => {
          localStorage.setItem('jwt', response.token);
          this.redirectToMainPage();
        },
        //TODO: handle errors
        error: err => {
          alert(err.error);
        }
      }
      );
  }

  public logout = () => {
    localStorage.removeItem('jwt');
    this.router.navigate(['login']);
  }
}
