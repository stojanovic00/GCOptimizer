import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SportsOrgPageComponent } from './sports-org-page.component';
import { RouterModule } from '@angular/router';
import { BrowserModule } from '@angular/platform-browser';
import { SportsOrgViewComponent } from 'src/app/components/sports-org-view/sports-org-view.component';
import { MaterialModule } from 'src/app/material/material/material.module';
import { JudgesViewRegisterComponent } from 'src/app/components/judges-view-register/judges-view-register.component';
import { ContestantsViewRegisterComponent } from 'src/app/components/contestants-view-register/contestants-view-register.component';
import { CompetitionCreateComponent } from 'src/app/components/competition-create/competition-create.component'
import { CompetitionViewComponent } from 'src/app/components/competition-view/competition-view.component';
import { ApplicationCreateContestantComponent } from 'src/app/components/application-create-contestant/application-create-contestant.component';
import { ApplicationCreateJudgeComponent } from 'src/app/components/application-create-judge/application-create-judge.component';
import { ApplicationViewComponent } from 'src/app/components/application-view/application-view.component';



@NgModule({
  declarations: [
    SportsOrgPageComponent,
    SportsOrgViewComponent,
    JudgesViewRegisterComponent,
    ContestantsViewRegisterComponent,
    CompetitionCreateComponent,
    CompetitionViewComponent,
    ApplicationViewComponent,
    ApplicationCreateJudgeComponent,
    ApplicationCreateContestantComponent

  ],
  imports: [
    CommonModule,
    BrowserModule,
    RouterModule,
    MaterialModule
  ]
})
export class SportsOrgPageModule { }
