import { Component, OnInit } from '@angular/core';
import { LoginService } from 'src/app/auth/services/login.service';

@Component({
  selector: 'app-sports-org-page',
  templateUrl: './sports-org-page.component.html',
  styleUrls: ['./sports-org-page.component.css']
})
export class SportsOrgPageComponent implements OnInit {

  constructor(private loginService: LoginService) { }

  ngOnInit(): void {
  }

  logOut(): void {
    this.loginService.logout();
  }
}
