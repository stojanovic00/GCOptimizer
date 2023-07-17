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



@NgModule({
  declarations: [
    SportsOrgPageComponent,
    SportsOrgViewComponent,
    JudgesViewRegisterComponent,
    ContestantsViewRegisterComponent,
    CompetitionCreateComponent
  ],
  imports: [
    CommonModule,
    BrowserModule,
    RouterModule,
    MaterialModule
  ]
})
export class SportsOrgPageModule { }
