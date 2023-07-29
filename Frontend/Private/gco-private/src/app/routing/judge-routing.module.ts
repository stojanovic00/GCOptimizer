import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { JudgeInfoComponent } from '../components/judge-info/judge-info.component';
import { JudgeAllContestantsComponent } from '../components/judge-all-contestants/judge-all-contestants.component';
import { JudgeCurrentContestantComponent } from '../components/judge-current-contestant/judge-current-contestant.component';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: 'info',
    component: JudgeInfoComponent
  },
  {
    path: 'all-contestants',
    component: JudgeAllContestantsComponent
  },
  {
    path: 'current-contestant',
    component: JudgeCurrentContestantComponent
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
