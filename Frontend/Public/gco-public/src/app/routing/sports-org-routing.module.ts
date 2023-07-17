import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SportsOrgViewComponent } from '../components/sports-org-view/sports-org-view.component';
import { JudgesViewRegisterComponent } from '../components/judges-view-register/judges-view-register.component';
import { ContestantsViewRegisterComponent } from '../components/contestants-view-register/contestants-view-register.component';
import { CompetitionCreateComponent } from '../components/competition-create/competition-create.component';
import { CompetitionViewComponent } from '../components/competition-view/competition-view.component';

const routes: Routes = [
  {
    path: 'view',
    component: SportsOrgViewComponent
  },
  {
    path: 'judge',
    component: JudgesViewRegisterComponent
  },
  {
    path: 'contestant',
    component: ContestantsViewRegisterComponent
  },
  {
    path: 'competition',
    component: CompetitionCreateComponent
  },
  {
    path: 'competition/view',
    component: CompetitionViewComponent
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
