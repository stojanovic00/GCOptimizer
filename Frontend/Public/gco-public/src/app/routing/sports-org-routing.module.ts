import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SportsOrgViewComponent } from '../components/sports-org-view/sports-org-view.component';

const routes: Routes = [
  {
    path: 'view',
    component: SportsOrgViewComponent
  },
  {
    path: '',
    redirectTo: 'view',
    pathMatch: 'full'
  },
  {
    path: '**',
    redirectTo: 'view',
    pathMatch: 'full'
  },
];

@NgModule({
  declarations: [],
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class SportsOrgRoutingModule { }
