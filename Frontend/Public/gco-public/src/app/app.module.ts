import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppComponent } from './app.component';
import { AppRoutingModule } from './routing/app-routing.module';
import { LoginPageModule } from './pages/login/login-page.module';
import { RegisterPageModule } from './pages/register-page/register-page.module';
import { SportsOrgPageModule } from './pages/sports-org-page/sports-org-page.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule, HTTP_INTERCEPTORS } from '@angular/common/http';
@NgModule({
  declarations: [
    AppComponent,
  ],
  imports: [
    LoginPageModule,
    RegisterPageModule,
    SportsOrgPageModule,


    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule

  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
