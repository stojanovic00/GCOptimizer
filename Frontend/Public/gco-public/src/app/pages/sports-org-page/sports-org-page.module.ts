import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { SportsOrgPageComponent } from './sports-org-page.component';
import { RouterModule } from '@angular/router';
import { BrowserModule } from '@angular/platform-browser';
import { SportsOrgViewComponent } from 'src/app/components/sports-org-view/sports-org-view.component';
import { MaterialModule } from 'src/app/material/material/material.module';



@NgModule({
  declarations: [
    SportsOrgPageComponent,
    SportsOrgViewComponent
  ],
  imports: [
    CommonModule,
    BrowserModule,
    RouterModule,
    MaterialModule
  ]
})
export class SportsOrgPageModule { }
