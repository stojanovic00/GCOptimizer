import { Component, OnInit } from '@angular/core';
import { Gender, getGenderName } from 'src/app/model/core/gender';
import { SprotsOrgService } from 'src/app/services/sprots-org.service';
import { Contestant, ContestantTable } from '../../model/core/contestant';
import { HttpErrorResponse } from '@angular/common/http';
import { dateToUnixTimeStamp, formatDate, unixTimeStampToDate } from '../../utils/date-utils'
import { FormGroup, FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-contestants-view-register',
  templateUrl: './contestants-view-register.component.html',
  styleUrls: ['./contestants-view-register.component.css']
})
export class ContestantsViewRegisterComponent implements OnInit {

  //View
  table : ContestantTable = {
    displayedColumns : ["fullName", "email", "gender", "dateOfBirth"],
    dataSource : [], 
    selectedRow : null
  }
  getGenderName = getGenderName

  unixDateToString = (unixDate: number) : string =>{
    return formatDate(unixTimeStampToDate(unixDate));
  }


  //Register
  regDialogOpened : boolean = false

  registrationForm = new FormGroup({
    fullName: new FormControl<string>('', [Validators.required]),
    email: new FormControl<string>('', [Validators.required]),
    gender: new FormControl<Gender>(Gender.Male, [Validators.required]),
    dateOfBirth: new FormControl<Date>(new Date(), [Validators.required]),
  });

  get RegistrationForm() {
    return this.registrationForm.controls;
  } 

  genderOptions = [
    { label: 'Male', value: Gender.Male },
    { label: 'Female', value: Gender.Female },
  ]


  constructor(
    private readonly soService : SprotsOrgService
  ) { }

  ngOnInit(): void {
    this.loadTable();
  }

  loadTable = () => {
    this.soService.getContestants().subscribe({
      next: (response: Contestant[]) => {
        this.table.dataSource = response;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }


//View
selectRow = (row: Contestant) => {
  this.table.selectedRow = row;
}

//Register
openRegisterDialog = () => {
  this.regDialogOpened = true;
}

closeRegisterDialog = () => {
  this.regDialogOpened = false;
  //Reset form
  this.RegistrationForm.fullName.setValue("");
  this.RegistrationForm.email.setValue("");
  this.RegistrationForm.gender.setValue(Gender.Male)
  this.RegistrationForm.dateOfBirth.setValue(new Date())
}

register = () => {

  let contestant: Contestant = {
    delegationMember: {
      fullName: this.RegistrationForm.fullName.value!,
      email: this.RegistrationForm.email.value!,
      gender: this.RegistrationForm.gender.value!,
      position: {
        name: "contestant"
      }
    },
    dateOfBirth: dateToUnixTimeStamp(this.RegistrationForm.dateOfBirth.value!)
  }

  this.soService.registerContestant(contestant).subscribe({
    next: (response: string) => {
      this.loadTable()
      this.closeRegisterDialog()
    },
    error: (err: HttpErrorResponse) => {
      alert(err.error);
    }
  })

}
}