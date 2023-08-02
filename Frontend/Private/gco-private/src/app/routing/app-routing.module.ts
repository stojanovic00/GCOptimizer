import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginPageComponent } from '../pages/login/login-page.component';
import { SportsOrgPageComponent } from '../pages/sports-org-page/sports-org-page.component';
import { SportsOrgRoutingModule } from './sports-org-routing.module';
import { RoleGuardService as RoleGuard } from '../auth/guards/role-guard.service';
import { IncognitoGuardService as IncognitoGuard } from '../auth/guards/incognito-guard.service';
import { JudgeInfoComponent } from '../components/judge-info/judge-info.component';
import { JudgePageComponent } from '../pages/judge-page/judge-page.component';
import { JudgeRoutingModule } from './judge-routing.module';

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
    data: { expectedRoles: ['SPORTS_ORG'] },
    loadChildren: () =>  
      import('./sports-org-routing.module').then(
        (m) => SportsOrgRoutingModule
      ),
  },
  {
    path: 'judge',
    component: JudgePageComponent,
    canActivate: [RoleGuard],
    data: { expectedRoles: ['D_JUDGE', 'E_JUDGE'] },
    loadChildren: () =>  
      import('./judge-routing.module').then(
        (m) => JudgeRoutingModule
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
