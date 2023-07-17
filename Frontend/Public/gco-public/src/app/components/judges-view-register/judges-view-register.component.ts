import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Judge, JudgeTable } from 'src/app/model/core/judge';
import { SprotsOrgService } from 'src/app/services/sprots-org.service';
import { LicenceType, getLicenceTypeName  } from '../../model/core/licence-type';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import {Gender, getGenderName} from '../../model/core/gender';

@Component({
  selector: 'app-judges-view-register',
  templateUrl: './judges-view-register.component.html',
  styleUrls: ['./judges-view-register.component.css']
})

export class JudgesViewRegisterComponent implements OnInit {
  //View
  table : JudgeTable = {
    displayedColumns : ["fullName", "email", "licenceType", "licenceName", "gender"],
    dataSource : [], 
    selectedRow : null
  }
  getLicenceTypeName = getLicenceTypeName
  getGenderName = getGenderName

  //Register
  regDialogOpened : boolean = false

  registrationForm = new FormGroup({
    fullName: new FormControl<string>('', [Validators.required]),
    email: new FormControl<string>('', [Validators.required]),
    gender: new FormControl<Gender>(Gender.Male, [Validators.required]),
    licenceType: new FormControl<LicenceType>(LicenceType.National, [Validators.required]),
    licenceName: new FormControl<string>('', [Validators.required]),
  });

  get RegistrationForm() {
    return this.registrationForm.controls;
  } 

  genderOptions = [
    { label: 'Male', value: Gender.Male },
    { label: 'Female', value: Gender.Female },
  ]


  licenceTypeOptions = [
    { label: 'National', value: LicenceType.National },
    { label: 'International', value: LicenceType.International },
  ]

  constructor(
    private readonly soService : SprotsOrgService
  ) { }

  ngOnInit(): void {
    this.loadTable();
  }

  //View
  selectRow = (row: Judge) => {
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
    this.RegistrationForm.licenceType.setValue(LicenceType.National)
    this.RegistrationForm.licenceName.setValue("");
  }

  register = () => {

    let judge : Judge = {
      delegationMember : {
        fullName : this.RegistrationForm.fullName.value!,
        email : this.RegistrationForm.email.value!,
        gender : this.RegistrationForm.gender.value!,
        position:{
          name : "judge"
        } 
      },
      licenceType : this.RegistrationForm.licenceType.value!,
      licenceName : this.RegistrationForm.licenceName.value!
    }

    this.soService.registerJudge(judge).subscribe({
      next: (response: string) => {
        this.loadTable()
        this.closeRegisterDialog()
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })

  }

  loadTable = () => {
    this.soService.getJudges().subscribe({
      next: (response: Judge[]) => {
        this.table.dataSource = response;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }
}
