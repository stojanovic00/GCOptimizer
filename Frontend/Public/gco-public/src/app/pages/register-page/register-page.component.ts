import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { RegistrationDto } from 'src/app/model/dto/registration-dto';
import { SprotsOrgService } from 'src/app/services/sprots-org.service';
import { HttpErrorResponse } from '@angular/common/http';

@Component({
  selector: 'app-register-page',
  templateUrl: './register-page.component.html',
  styleUrls: ['./register-page.component.css']
})
export class RegisterPageComponent implements OnInit {


  registrationForm = new FormGroup({
    email: new FormControl<string>('', [Validators.required]),
    password: new FormControl<string>('', [Validators.required]),
   //Role will always be sports organisation
    name: new FormControl<string>('', [Validators.required]),
    phoneNumber: new FormControl<string>('', [Validators.required]),
    contactPersonFullName: new FormControl<string>('', [Validators.required]),
    country: new FormControl<string>('', [Validators.required]),
    city: new FormControl<string>('', [Validators.required]),
    street: new FormControl<string>('', [Validators.required]),
    streetNumber: new FormControl<string>('', [Validators.required])
  });

  get regform() {
    return this.registrationForm.controls;
  } 


  constructor(
      private readonly router: Router,
      private readonly sportsOrgService: SprotsOrgService,
      ) { }

  ngOnInit(): void {
  }


  register = () =>{
    let dto: RegistrationDto = {
      account: {
        email: this.regform.email.value ?? "",
        password: this.regform.password.value ?? "",
        role: {
          name: "sports_org"
        }
      },
      sportsOrganisation: {
        name: this.regform.name.value ?? "",
        phoneNumber: this.regform.phoneNumber.value ?? "",
        contactPersonFullName: this.regform.contactPersonFullName.value ?? "",
        email: this.regform.email.value ?? "",
        competitionOrganisingPrivilege : false,
        address: {
          country: this.regform.country.value ?? "",
          city: this.regform.city.value ?? "",
          street: this.regform.street.value ?? "",
          streetNumber: this.regform.streetNumber.value ?? "",
        }
    }
   } 

    this.sportsOrgService.register(dto).subscribe({
      next: (response) => {
        this.router.navigate(["login"])
      },
      error:
        (err: HttpErrorResponse) => {
          alert(err.error)
        }
    })

  }
backToLogin = () => {
  this.router.navigate(["login"])
}

}
