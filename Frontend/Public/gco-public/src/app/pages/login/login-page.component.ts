import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { LoginService } from 'src/app/auth/services/login.service';
import { LoginDto } from 'src/app/auth/model/login-dto';

@Component({
  selector: 'app-login-page',
  templateUrl: './login-page.component.html',
  styleUrls: ['./login-page.component.css']
})
export class LoginPageComponent implements OnInit {

  loginForm = new FormGroup({
    email: new FormControl<string>('', [Validators.required]),
    password: new FormControl<string>('', [Validators.required]),
  });

  get Loginform() {
    return this.loginForm.controls;
  } 

  constructor(
    private readonly router : Router,
    private readonly loginService : LoginService
  ) { }

  ngOnInit(): void {
  }


  login = () => {
    let loginDto : LoginDto = {
      email: this.Loginform.email.value ?? "",
      password: this.Loginform.email.value ?? "",
    };

    this.loginService.login(loginDto);
  }

  goToRegister = () => {
  this.router.navigate(["register"])
  }
}
