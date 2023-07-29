import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { BrowserModule } from '@angular/platform-browser';
import { RouterModule } from '@angular/router';
import { MaterialModule } from '../../material/material.module';
import { JudgePageComponent } from './judge-page.component';
import { JudgeAllContestantsComponent } from 'src/app/components/judge-all-contestants/judge-all-contestants.component';
import { JudgeCurrentContestantComponent } from 'src/app/components/judge-current-contestant/judge-current-contestant.component';
import { JudgeInfoComponent } from 'src/app/components/judge-info/judge-info.component';



@NgModule({
  declarations: [
    JudgePageComponent,
    JudgeInfoComponent,
    JudgeAllContestantsComponent,
    JudgeCurrentContestantComponent

  ],
  imports: [
    CommonModule,
    BrowserModule,
    RouterModule,
    MaterialModule
  ]
})
export class JudgePageModule { }
