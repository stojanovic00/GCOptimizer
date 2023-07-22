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



@NgModule({
  declarations: [
    SportsOrgPageComponent,
    SportsOrgViewComponent,
    CompetitionViewComponent,
    ScheduleCreateComponent,
    ScheduleViewComponent,
    ApplicationViewComponent

  ],
  imports: [
    CommonModule,
    BrowserModule,
    RouterModule,
    MaterialModule
  ]
})
export class SportsOrgPageModule { }
