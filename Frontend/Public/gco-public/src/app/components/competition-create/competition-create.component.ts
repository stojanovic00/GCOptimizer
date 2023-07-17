import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Competition } from 'src/app/model/core/competition';
import { CompetitionType } from 'src/app/model/core/competition-type';
import { Gender } from 'src/app/model/core/gender';
import { CompetitionService } from 'src/app/services/competition.service';
import {dateToUnixTimeStamp} from '../../utils/date-utils'
import { DelegationMemberProposition } from 'src/app/model/core/delegation-member-proposition';
import { AgeCategory } from 'src/app/model/core/age-category';
import { IdResponse } from 'src/app/model/dto/id-response';
import { Router } from '@angular/router';

@Component({
  selector: 'app-competition-create',
  templateUrl: './competition-create.component.html',
  styleUrls: ['./competition-create.component.css']
})
export class CompetitionCreateComponent implements OnInit {

  //Competition
  createdCompetitionId : string = "";

  competitionForm = new FormGroup({
    name: new FormControl<string>('', [Validators.required]),
    startDate: new FormControl<Date>(new Date(), [Validators.required]),
    endDate: new FormControl<Date>(new Date(), [Validators.required]),
    gender: new FormControl<Gender>(Gender.Male, [Validators.required]),
    type: new FormControl<CompetitionType>(CompetitionType.Qualifications, [Validators.required]),
    tiebreak: new FormControl<boolean>(false, [Validators.required]),
    country: new FormControl<string>('', [Validators.required]),
    city: new FormControl<string>('', [Validators.required]),
    street: new FormControl<string>('', [Validators.required]),
    streetNumber: new FormControl<string>('', [Validators.required]),
    baseContestantNumber: new FormControl<number>(0, [Validators.required]),
    bonusContestantNumber: new FormControl<number>(0, [Validators.required]),
    multiCategoryTeam: new FormControl<boolean>(false, [Validators.required]),
  });

  genderOptions = [
    { label: 'Male', value: Gender.Male },
    { label: 'Female', value: Gender.Female },
  ]

  competitionTypeOptions = [
    { label: 'Qualifications', value: CompetitionType.Qualifications},
    { label: 'Team finals', value: CompetitionType.TeamFinals },
    { label: 'Apparatus finals', value: CompetitionType.ApparatusFinals },
    { label: 'All around finals', value: CompetitionType.AllAroundFinals }
  ]

  get CompetitionForm() {
    return this.competitionForm.controls;
  } 


  //Delegation member proposition
  delMemPropForm = new FormGroup({
    positionName: new FormControl<string>('', [Validators.required]),
    minNumber: new FormControl<number>(0, [Validators.required]),
    maxNumber: new FormControl<number>(0, [Validators.required]),
  });

  positionOptions = [
    { label: 'Judge', value: "judge"},
    { label: 'Contestant', value: "contestant"}
  ]

  get DelMemPropForm() {
    return this.delMemPropForm.controls;
  } 

  addedPropositions : DelegationMemberProposition[] = []


  //Age categoty
  ageCatForm = new FormGroup({
    name: new FormControl<string>('', [Validators.required]),
    minAge: new FormControl<number>(0, [Validators.required]),
    maxAge: new FormControl<number>(0, [Validators.required]),
  });

  get AgeCatForm() {
    return this.ageCatForm.controls;
  } 

  addedAgeCategories : AgeCategory[] = []

  constructor(
    private readonly compService: CompetitionService,
    private readonly router: Router
  ) { }

  ngOnInit(): void {
  }

createCompetition = () =>{
  let competition: Competition = {
      name: this.CompetitionForm.name.value!,
      startDate: dateToUnixTimeStamp(this.CompetitionForm.startDate.value!),
      endDate: dateToUnixTimeStamp(this.CompetitionForm.endDate.value!),
      gender: this.CompetitionForm.gender.value!,
      type: this.CompetitionForm.type.value!,
      tiebreak: this.CompetitionForm.tiebreak.value!,
      address :{
        country: this.CompetitionForm.country.value!,
        city: this.CompetitionForm.city.value!,
        street: this.CompetitionForm.street.value!,
        streetNumber: this.CompetitionForm.streetNumber.value!,
      },
      teamComposition :{
        baseContestantNumber: this.CompetitionForm.baseContestantNumber.value!,
        bonusContestantNumber: this.CompetitionForm.bonusContestantNumber.value!,
        multiCategoryTeam: this.CompetitionForm.multiCategoryTeam.value!,
      }
   }
  this.compService.create(competition).subscribe({
    next: (response: IdResponse) => {
      this.createdCompetitionId = response.id;
    },
    error: (err: HttpErrorResponse) => {
      alert(err.error);
    }
  })
}
addDelMemProp = () =>{ 
  let proposition :DelegationMemberProposition ={
    minNumber : this.DelMemPropForm.minNumber.value!,
    maxNumber : this.DelMemPropForm.maxNumber.value!,
    position : {
      name: this.DelMemPropForm.positionName.value!
    }
  } 
  
  this.compService.addDelegationMemberProposition(proposition, this.createdCompetitionId).subscribe({
    next: (response: string) => {
      this.addedPropositions.push(proposition);
    },
    error: (err: HttpErrorResponse) => {
      alert(err.error);
    }
  })

}

addAgeCat = () =>{
  let ageCat :AgeCategory ={
    minAge : this.AgeCatForm.minAge.value!,
    maxAge : this.AgeCatForm.maxAge.value!,
    name: this.AgeCatForm.name.value!
  } 
  
  this.compService.addAgeCategory(ageCat, this.createdCompetitionId).subscribe({
    next: (response: string) => {
      this.addedAgeCategories.push(ageCat);
    },
    error: (err: HttpErrorResponse) => {
      alert(err.error);
    }
  })
}

finishCompCreation = () =>{
    this.router.navigate(['sports-org/view']);
}

}
