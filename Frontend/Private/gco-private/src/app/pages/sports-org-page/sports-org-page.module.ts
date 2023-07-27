import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SportsOrgPageComponent } from './sports-org-page.component';
import { SportsOrgViewComponent } from 'src/app/components/sports-org-view/sports-org-view.component';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { MaterialModule } from 'src/app/material/material.module';
import { CompetitionViewComponent } from 'src/app/components/competition-view/competition-view.component';
import { ScheduleCreateComponent } from 'src/app/components/schedule-create/schedule-create.component';
import { ScheduleViewComponent } from 'src/app/components/schedule-view/schedule-view.component';
import { ApplicationViewComponent } from 'src/app/components/application-view/application-view.component';
import { ChoosePanelComponent } from 'src/app/components/choose-panel/choose-panel.component';
import { FormPanelComponent } from 'src/app/components/form-panel/form-panel.component';



@NgModule({
  declarations: [
    SportsOrgPageComponent,
    SportsOrgViewComponent,
    CompetitionViewComponent,
    ScheduleCreateComponent,
    ScheduleViewComponent,
    ApplicationViewComponent,
    ChoosePanelComponent,
    FormPanelComponent


  ],
  imports: [
    CommonModule,
    BrowserModule,
    RouterModule,
    MaterialModule
  ]
})
export class SportsOrgPageModule { }
