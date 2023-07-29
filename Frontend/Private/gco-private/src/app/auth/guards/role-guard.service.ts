import { Injectable } from '@angular/core';
import { Router, ActivatedRouteSnapshot, CanActivate } from '@angular/router';
import { JwtHelperService } from '@auth0/angular-jwt';
import { AuthService } from '../services/auth.service';

@Injectable({
  providedIn: 'root'
})
export class RoleGuardService implements CanActivate {

  constructor(
    private readonly auth: AuthService,
    private readonly router: Router,
    private readonly jwtHelper: JwtHelperService
  ) { }

  canActivate(route: ActivatedRouteSnapshot): boolean {
    const expectedRoles :string[] = route.data['expectedRoles'];
    const token = localStorage.getItem('jwt');
    if (!token) {
      this.router.navigate(['login']);
      return false;
    }

    const tokenPayload = this.jwtHelper.decodeToken(token!);
    const role = tokenPayload['role'];

    const hasRole = expectedRoles.includes(role)

    if (
      !this.auth.isAuthenticated() || !hasRole
    ) {
      this.router.navigate(['login']);
      return false;
    }
    return true;

  }
}
