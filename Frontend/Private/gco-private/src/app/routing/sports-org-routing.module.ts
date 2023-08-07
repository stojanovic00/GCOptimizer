import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { SportsOrgViewComponent } from '../components/sports-org-view/sports-org-view.component';
import { CompetitionViewComponent } from '../components/competition-view/competition-view.component';
import { ScheduleCreateComponent } from '../components/schedule-create/schedule-create.component';
import { ScheduleViewComponent } from '../components/schedule-view/schedule-view.component';
import { ApplicationViewComponent } from '../components/application-view/application-view.component';
import { ChoosePanelComponent } from '../components/choose-panel/choose-panel.component';
import { FormPanelComponent } from '../components/form-panel/form-panel.component';
import { CompetitionMonitoringComponent } from '../components/competition-monitoring/competition-monitoring.component';
import { ScoreboardsComponent } from '../components/scoreboards/scoreboards.component';

const routes: Routes = [
  {
    path: 'view',
    component: SportsOrgViewComponent
  },
  {
    path: 'competition/view',
    component: CompetitionViewComponent
  },
  {
    path: 'competition/:id/application/view',
    component: ApplicationViewComponent
  },
  {
    path: 'competition/:id/schedule/make',
    component: ScheduleCreateComponent
  },
  {
    path: 'competition/:id/schedule/view',
    component: ScheduleViewComponent
  },
  {
    path: 'competition/:id/judging-panel/unassigned',
    component: ChoosePanelComponent
  },
  {
    path: 'competition/:id/judging-panel/form/:apparatus',
    component: FormPanelComponent
  },
  {
    path: 'competition/:id/monitoring',
    component: CompetitionMonitoringComponent
  },
  {
    path: 'competition/:id/scoreboards',
    component: ScoreboardsComponent
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
