import { Component, OnInit } from '@angular/core';
import { LoginService } from 'src/app/auth/services/login.service';

@Component({
  selector: 'app-judge-page',
  templateUrl: './judge-page.component.html',
  styleUrls: ['./judge-page.component.css']
})


export class JudgePageComponent implements OnInit {

  constructor(private loginService: LoginService) { }

  ngOnInit(): void {
  }

  logOut(): void {
    this.loginService.logout();
  }
}