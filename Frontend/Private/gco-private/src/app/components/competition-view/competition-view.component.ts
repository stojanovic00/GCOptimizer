import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { formatAddress } from 'src/app/model/core/address';
import { Competition, CompetitionTable } from 'src/app/model/core/competition';
import { getCompetitionTypeName } from 'src/app/model/core/competition-type';
import { getGenderName } from 'src/app/model/core/gender';
import { ScoreboardBundle } from 'src/app/model/dto/scoreboard-bundle';
import { CompetitionService } from 'src/app/services/competition.service';
import { ScoringService } from 'src/app/services/scoring.service';
import { unixDateToString } from 'src/app/utils/date-utils';

@Component({
  selector: 'app-competition-view',
  templateUrl: './competition-view.component.html',
  styleUrls: ['./competition-view.component.css']
})
export class CompetitionViewComponent implements OnInit {

  //Table
  table : CompetitionTable = {
    displayedColumns : ["name", "gender", "startDate", "endDate", "address", "type", "organizer"],
    dataSource : [], 
    selectedRow : null
  }

  getGenderName = getGenderName
  unixDateToString = unixDateToString
  formatAddress = formatAddress
  getCompetitionTypeName = getCompetitionTypeName

  //Details
  selectedCompetition : Competition | null = null
  selectedCompetitionFinished : boolean = false;
  detailsDialogOpened : boolean = false

  constructor(
    private readonly router : Router,
    private readonly compService : CompetitionService,
    private readonly scService : ScoringService,
  ) { }

  ngOnInit(): void {
    this.loadTable();
  }

  loadTable = () => {
    this.compService.getAll().subscribe({
      next: (response: Competition[]) => {
        this.table.dataSource = response;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }

//Implemented double click detection
 private alreadyClicked = false;
 selectRow(row: Competition) {
   this.table.selectedRow = row;

   if (!this.alreadyClicked) {
     // First click
     this.alreadyClicked = true;

     // Set a timer to reset the flag after a short delay (e.g., 300ms)
     setTimeout(() => {
       this.alreadyClicked = false;
     }, 300);
   }
   else {
     this.showDetailsDialog();
     // Reset the flag
     this.alreadyClicked = false;
   }
 }

 showDetailsDialog() {
    this.compService.getById(this.table.selectedRow?.id!).subscribe({
      next: (response: Competition) => {
        this.selectedCompetition = response;

        //Check if competition finished
        this.scService.getScoreboards(this.selectedCompetition.id!).subscribe({
          next: (response: ScoreboardBundle) => {
            if(response.allAroundScoreboards && response.teamScoreboards){
              this.selectedCompetitionFinished = true;
            }
            else{
              this.selectedCompetitionFinished = false;
            }
            this.detailsDialogOpened = true;
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

 closeDetailsDialog = () =>{
   this.detailsDialogOpened = false;
 }


  goToMakeSchedule = () => {
    let compId = this.table.selectedRow?.id;
    this.router.navigate(['sports-org/competition/'  + compId + '/schedule/make']);
  }

  goToViewSchedule = () => {
    let compId = this.table.selectedRow?.id;
    this.router.navigate(['sports-org/competition/'  + compId + '/schedule/view']);
  }

  viewApplications = () => {
    let compId = this.table.selectedRow?.id;
    this.router.navigate(['sports-org/competition/' + compId + '/application/view']);
  }
  goToScoreBoards = () =>{
      this.router.navigate(['sports-org/competition/' + this.selectedCompetition?.id + '/scoreboards']); 
  }
}
