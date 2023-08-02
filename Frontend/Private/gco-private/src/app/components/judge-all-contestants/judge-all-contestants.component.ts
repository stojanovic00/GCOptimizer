import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Apparatus } from 'src/app/model/core/apparatus';
import { ContestantScoring } from 'src/app/model/core/contestant-scoring';
import { ContestantScoringDto, ContestantScoringDtoTable } from 'src/app/model/dto/contestant-scoring-dto';
import { JudgeJudgingInfo } from 'src/app/model/dto/judge-judging-info';
import { ScoringService } from 'src/app/services/scoring.service';

@Component({
  selector: 'app-judge-all-contestants',
  templateUrl: './judge-all-contestants.component.html',
  styleUrls: ['./judge-all-contestants.component.css']
})
export class JudgeAllContestantsComponent implements OnInit {
  constructor(
    private readonly route : ActivatedRoute,
    private readonly router : Router,
    private readonly scService : ScoringService,
  ){}

judgingInfo: JudgeJudgingInfo | null = null;
apparatus : Apparatus | null = null
currentContestant: ContestantScoring | null = null

contestantsTable : ContestantScoringDtoTable = {
    displayedColumns: ['compId', 'name', 'team', 'ageCat', 'organization'],
    dataSource: [],
    selectedRow: null
}
  ngOnInit(): void {
      this.loadData()
  }

  loadData = () =>{
    //Get judging info
    this.scService.getLoggedJudgeInfo().subscribe({
      next: (response: JudgeJudgingInfo) => {
        this.judgingInfo = response;
        this.loadContestants()
        this.loadCurrentContestant()
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }


  loadContestants = () =>{
    this.scService.getCurrentContestantsForApparatus(this.judgingInfo?.competitionId!, this.judgingInfo?.apparatus ?? 0).subscribe({
      next: (response: ContestantScoringDto[]) => {
        this.contestantsTable.dataSource = response
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }

  loadCurrentContestant = () =>{
    this.scService.getCurrentApparatusContestant(this.judgingInfo?.competitionId!, this.judgingInfo?.apparatus ?? 0).subscribe({
      next: (response: ContestantScoring) => {
        this.currentContestant = response
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }

}
