import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginPageComponent } from '../pages/login/login-page.component';
import { SportsOrgPageComponent } from '../pages/sports-org-page/sports-org-page.component';
import { SportsOrgRoutingModule } from './sports-org-routing.module';
import { RoleGuardService as RoleGuard } from '../auth/guards/role-guard.service';
import { IncognitoGuardService as IncognitoGuard } from '../auth/guards/incognito-guard.service';

const routes: Routes = [
  {
    path: 'login',
    canActivate: [IncognitoGuard],
    component: LoginPageComponent
  },
  {
    path: 'sports-org',
    component: SportsOrgPageComponent,
    canActivate: [RoleGuard],
    data: { expectedRole: 'SPORTS_ORG' },
    loadChildren: () =>  
      import('./sports-org-routing.module').then(
        (m) => SportsOrgRoutingModule
      ),
  },
  {
    path: '',
    redirectTo: 'login',
    pathMatch: 'full'
  },
  {
    path: '**',
    redirectTo: 'login',
    pathMatch: 'full'
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
