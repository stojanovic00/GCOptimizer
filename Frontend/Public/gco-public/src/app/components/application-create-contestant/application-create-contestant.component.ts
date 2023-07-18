import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { AgeCategory, AgeCategoryTable, formatAgeCategory } from 'src/app/model/core/age-category';
import { Competition } from 'src/app/model/core/competition';
import { ApparatusAnnouncement, ApparatusAnnouncementTable, formatApparatusAnnouncements } from 'src/app/model/core/apparatus-announcement';
import { ContestantTable, Contestant } from 'src/app/model/core/contestant';
import { CompetitionService } from 'src/app/services/competition.service';
import { SprotsOrgService } from 'src/app/services/sprots-org.service';
import { Apparatus } from 'src/app/model/core/apparatus';
import { Gender } from 'src/app/model/core/gender';
import { getApparatusName } from 'src/app/model/core/apparatus';
import { CreateContestantApplicationRequest} from 'src/app/model/dto/create-contestant-application-request'
import { ContestantApplicationPreview, ContestantApplicationPreviewTable} from 'src/app/model/core/contestant-application'

@Component({
  selector: 'app-application-create-contestant',
  templateUrl: './application-create-contestant.component.html',
  styleUrls: ['./application-create-contestant.component.css']
})
export class ApplicationCreateContestantComponent implements OnInit {

  //View
  tableContestant : ContestantTable = {
    displayedColumns : ["fullName"],
    dataSource : [], 
    selectedRow : null
  }

  tableAgeCategory : AgeCategoryTable = {
    displayedColumns : ["name", "minAge", "maxAge"],
    dataSource : [], 
    selectedRow : null
  }



maleApparatuses : ApparatusAnnouncement[] = [
  { apparatus : Apparatus.Floor, },
  { apparatus : Apparatus.PommelHorse, },
  { apparatus : Apparatus.StillRings, },
  { apparatus : Apparatus.Vault, },
  { apparatus : Apparatus.ParallelBars, },
  { apparatus : Apparatus.HorizontalBar, }
]

femaleApparatuses : ApparatusAnnouncement[] = [
  { apparatus : Apparatus.Floor, },
  { apparatus : Apparatus.UnevenBars, },
  { apparatus : Apparatus.BalanceBeam, },
  { apparatus : Apparatus.Vault, },
]

tableApparatus : ApparatusAnnouncementTable = {
    displayedColumns : ["name"],
    dataSource : [], 
    selectedRows : []
}
getApparatusName = getApparatusName



addedApplications : ContestantApplicationPreview[] = []
tableApplicationsContestant : ContestantApplicationPreviewTable = {
    displayedColumns : ["fullName", "teamNumber", "ageCategory", "apparatuses"],
    dataSource : [], 
    selectedRow : null
}

formatAgeCategory = formatAgeCategory 
formatApparatusAnnouncements = formatApparatusAnnouncements



  //Create
  competitionId : string = "";

  teamForm = new FormGroup({
    teamNumber: new FormControl<number>(1, [Validators.required]),
  });

  public get TeamNumber() {
    return this.teamForm.controls.teamNumber.value;
  }


  constructor(
    private readonly soService : SprotsOrgService,
    private readonly compService : CompetitionService,
    private readonly route : ActivatedRoute,
    private readonly router : Router
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      this.competitionId = params.get('id') || "";
      this.loadData(this.competitionId);
    });
  }

loadData = (compId: string) => {
    //Competition with age categories loaded  
    this.compService.getById(compId).subscribe({
      next: (responseComp: Competition) => {
        this.tableAgeCategory.dataSource = responseComp.ageCategories!;

        //Contestants
        this.soService.getContestants().subscribe({
          next: (response: Contestant[]) => {
            this.tableContestant.dataSource = response;

            //Apparatuses
            if(responseComp.gender === Gender.Female){
              this.tableApparatus.dataSource = this.femaleApparatuses;
            }
            else{
              this.tableApparatus.dataSource = this.maleApparatuses;
            }

          },
          error: (err: HttpErrorResponse) => {
            alert(err.error);
          }
        });
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }

  //View
  selectRow = (row: Contestant) => {
    this.tableContestant.selectedRow = row;
  }

  selectAgeCategoryRow = (row: AgeCategory) => {
    this.tableAgeCategory.selectedRow = row;
  }

  selectApparatusRow = (row: ApparatusAnnouncement) => {
    if (!this.isApparatusSelected(row)) {
      this.tableApparatus.selectedRows.push(row)
    }
    else {
      // Remove the row from the selectedApparatuses 
      this.tableApparatus.selectedRows = this.tableApparatus.selectedRows.filter((item) => item !== row);
    };
  }
  

  isApparatusSelected = (apparatus : ApparatusAnnouncement) : boolean => {
    return this.tableApparatus.selectedRows.includes(apparatus);
  }


addContestantApplication = () =>{
  let application : CreateContestantApplicationRequest = {
    competitionId : this.competitionId,
    contestantId : this.tableContestant.selectedRow?.delegationMember.id!,
    teamNumber : this.TeamNumber!,
    ageCategoryId : this.tableAgeCategory.selectedRow?.id!,
    apparatusAnnouncements : this.tableApparatus.selectedRows
  }

  let applicationPreview :ContestantApplicationPreview = {
    contestantFullName : this.tableContestant.selectedRow?.delegationMember.fullName!,
    teamNumber : this.TeamNumber!,
    ageCategory : this.tableAgeCategory.selectedRow!,
    apparatusAnnouncements: this.tableApparatus.selectedRows
  }

  this.compService.addContestantApplication(application, this.competitionId).subscribe({
    next: (response: string) => {
      //It has to be done like this because it will change only if pointer to array changes
      this.addedApplications.push(applicationPreview);
      this.tableApplicationsContestant.dataSource = this.addedApplications.slice();

      //Reset
      this.tableContestant.selectedRow = null;
      this.teamForm.controls.teamNumber.setValue(1)
      this.tableAgeCategory.selectedRow = null;
      this.tableApparatus.selectedRows = []

    },
    error: (err: HttpErrorResponse) => {
      alert(err.error);
    }
  });

}


finishCreation = () =>
{
    this.router.navigate(['sports-org/competition/view']);
}
}
