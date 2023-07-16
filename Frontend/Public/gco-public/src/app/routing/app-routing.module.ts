import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginPageComponent } from 'src/app/pages/login/login-page.component';
import { RegisterPageComponent } from '../pages/register-page/register-page.component';
import { SportsOrgPageComponent } from '../pages/sports-org-page/sports-org-page.component';
import { SportsOrgRoutingModule } from './sports-org-routing.module';

const routes: Routes = [
  {
    path: 'login',
    component: LoginPageComponent
  },
  {
    path: 'register',
    component: RegisterPageComponent
  },
  {
    path: 'sports-org',
    component: SportsOrgPageComponent,
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
