import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './routing/app-routing.module';
import { AppComponent } from './app.component';
import { LoginPageModule } from './pages/login/login-page.module';
import { HTTP_INTERCEPTORS, HttpClientModule } from '@angular/common/http';
import { JWT_OPTIONS, JwtHelperService } from '@auth0/angular-jwt';
import { AuthInterceptor } from './auth/auth.interceptor';
import { SportsOrgPageModule } from './pages/sports-org-page/sports-org-page.module';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { ChoosePanelComponent } from './components/choose-panel/choose-panel.component';
import { FormPanelComponent } from './components/form-panel/form-panel.component';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    LoginPageModule,
    SportsOrgPageModule,


    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule
  ],
  providers: [
    { provide: JWT_OPTIONS, useValue: JWT_OPTIONS },
    JwtHelperService,
    {
      provide: HTTP_INTERCEPTORS,
      useClass: AuthInterceptor,
      multi: true,
    },
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
