import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Apparatus, ApparatusTable, getApparatusName } from 'src/app/model/core/apparatus';
import { ApparatusTypeFoScheduling } from 'src/app/model/core/apparatus-type-for-scheduling';
import { Competition } from 'src/app/model/core/competition';
import { Gender } from 'src/app/model/core/gender';
import { SchedulingParameters } from 'src/app/model/core/scheduling-paramters';
import { CompetitionService } from 'src/app/services/competition.service';
import { ScheduleService } from 'src/app/services/schedule.service';
import { dateToUnixTimeStamp, timeStringToDate } from 'src/app/utils/date-utils';
import {ScheduleDtoToScheduleView} from '../../model/mappers/schedule-mapper'
import { ScheduleSessionView } from 'src/app/model/dto/schedule-session-view';
import { Schedule } from 'src/app/model/core/schedule';

@Component({
  selector: 'app-schedule-create',
  templateUrl: './schedule-create.component.html',
  styleUrls: ['./schedule-create.component.css']
})
export class ScheduleCreateComponent implements OnInit {

 
  paramsForm = new FormGroup({
    startTime: new FormControl<string>('', [Validators.required]),
    endTime: new FormControl<string>('', [Validators.required]),
    warmupRoomAvailable: new FormControl<boolean>(false, [Validators.required]),
    generalWarmupTime: new FormControl<number>(0, [Validators.required]),
    warmupTime: new FormControl<number>(0, [Validators.required]),
    warmupsPerApparatus: new FormControl<number>(0, [Validators.required]),
    contestantNumPerApparatus: new FormControl<number>(0, [Validators.required]),
    executionTime: new FormControl<number>(0, [Validators.required]),
    apparatusRotationTime: new FormControl<number>(0, [Validators.required]),
    medalCeremonyAfterOneSessionTime: new FormControl<number>(0, [Validators.required]),
    finalMedalCeremonyTime: new FormControl<number>(0, [Validators.required]),
    // halfApparatusPerSessionMode: new FormControl<boolean>(false, [Validators.required]),
  });

  public get ParamsForm() {
    return this.paramsForm.controls;
  }

  compGender : Gender = Gender.Male
  availableApparatuses : Apparatus[] = []
  availableApparatusesTable : ApparatusTable = {
    displayedColumns : ["name"],
    dataSource : [],
    selectedRow : null 
  }
  chosenApparatuses: Apparatus[] = []
  chosenApparatusesTable : ApparatusTable = {
    displayedColumns : ["name"],
    dataSource : [],
    selectedRow : null 
  }

  getApparatusName = getApparatusName

  maleApparatuses : Apparatus[] = [
    Apparatus.Floor,
    Apparatus.PommelHorse,
    Apparatus.StillRings,
    Apparatus.Vault,
    Apparatus.ParallelBars,
    Apparatus.HorizontalBar
  ]
  
  femaleApparatuses: Apparatus[] = [
    Apparatus.Floor,
    Apparatus.UnevenBars,
    Apparatus.BalanceBeam,
    Apparatus.Vault
  ]

  sessionViews : ScheduleSessionView[] = []

  constructor(
    private readonly route : ActivatedRoute,
    private readonly router : Router,
    private readonly schService : ScheduleService,
    private readonly compService : CompetitionService
  ){}

  competitionId : string = "";


  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      this.competitionId = params.get('id') || "";
      this.loadData()
    });
  }

  loadData = () => {
             this.compService.getById(this.competitionId).subscribe({
              next: (response: Competition) => {
    
                //Apparatuses
                if(response.gender === Gender.Female){
                    this.availableApparatuses = this.femaleApparatuses;
                }
                else{
                  this.availableApparatuses = this.maleApparatuses;
                }

                this.availableApparatusesTable.dataSource = this.availableApparatuses.slice(); 
                this.compGender = response.gender;
              },
              error: (err: HttpErrorResponse) => {
                alert(err.error);
              }
            }); 
  }

  isApparatusSelected = (apparatus : Apparatus) : boolean => {
    return this.availableApparatusesTable.selectedRow === apparatus;
  }

//Implemented double click detection
private alreadyClicked = false;
selectRow(row: Apparatus) {
  this.availableApparatusesTable.selectedRow = row;

  if (!this.alreadyClicked) {
    // First click
    this.alreadyClicked = true;

    // Set a timer to reset the flag after a short delay (e.g., 300ms)
    setTimeout(() => {
      this.alreadyClicked = false;
    }, 300);
  }
  else {
    this.addToChosenApparatuses();
    // Reset the flag
    this.alreadyClicked = false;
  }
}

addToChosenApparatuses() {
   this.chosenApparatuses.push(this.availableApparatusesTable.selectedRow!);
   this.chosenApparatusesTable.dataSource = this.chosenApparatuses.slice();

   this.availableApparatuses = this.availableApparatuses.filter(app => app !== this.availableApparatusesTable.selectedRow!);
   this.availableApparatusesTable.dataSource = this.availableApparatuses.slice();
}

clearSelectedApparatuses = () =>{
  this.chosenApparatuses = [];
  this.chosenApparatusesTable.dataSource = this.chosenApparatuses.slice();

  if(this.compGender === Gender.Female){
      this.availableApparatuses = this.femaleApparatuses;
  }
  else{
      this.availableApparatuses = this.maleApparatuses;
  }


  this.availableApparatusesTable.dataSource = this.availableApparatuses.slice(); 
}

  generateSchedule = () => {

      let apparatusOrder : ApparatusTypeFoScheduling[] = []
      this.chosenApparatusesTable.dataSource = this.maleApparatuses;
      apparatusOrder = this.chosenApparatusesTable.dataSource.map(app => {
          return {
            type : app
          }
      })

    let params : SchedulingParameters = {
        competitionId : this.competitionId,
        startTime : dateToUnixTimeStamp(timeStringToDate(this.ParamsForm.startTime.value!)),
        endTime : dateToUnixTimeStamp(timeStringToDate(this.ParamsForm.endTime.value!)),
        warmupRoomAvailable : this.ParamsForm.warmupRoomAvailable.value!,
        generalWarmupTime : this.ParamsForm.generalWarmupTime.value!,
        warmupTime : this.ParamsForm.warmupTime.value!,
        warmupsPerApparatus : this.ParamsForm.warmupsPerApparatus.value!,
        contestantNumPerApparatus : this.ParamsForm.contestantNumPerApparatus.value!,
        executionTime : this.ParamsForm.executionTime.value!,
        apparatusRotationTime : this.ParamsForm.apparatusRotationTime.value!,
        medalCeremonyAfterOneSessionTime : this.ParamsForm.medalCeremonyAfterOneSessionTime.value!,
        finalMedalCeremonyTime : this.ParamsForm.finalMedalCeremonyTime.value!,
        halfApparatusPerSessionMode : false,
        apparatusOrder : apparatusOrder
    }


      this.schService.generateSchedule(params).subscribe({
      next: (response: Schedule) => {
          this.sessionViews = ScheduleDtoToScheduleView(response);
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    }); 
    alert("Schedule is generating...")
  }

saveSchedule = () => {
    this.router.navigate(['sports-org/competition/'  + this.competitionId + '/schedule/view']);
}
  
}

