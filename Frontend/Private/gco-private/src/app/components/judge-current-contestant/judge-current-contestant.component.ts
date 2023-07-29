import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Apparatus } from 'src/app/model/core/apparatus';
import { JudgeJudgingInfo } from 'src/app/model/dto/judge-judging-info';
import { ScoringService } from 'src/app/services/scoring.service';

@Component({
  selector: 'app-judge-current-contestant',
  templateUrl: './judge-current-contestant.component.html',
  styleUrls: ['./judge-current-contestant.component.css']
})
export class JudgeCurrentContestantComponent implements OnInit {


  constructor(
    private readonly route : ActivatedRoute,
    private readonly router : Router,
    private readonly scService : ScoringService,
  ){}

judgingInfo: JudgeJudgingInfo | null = null;
competitionId : string = ""
apparatus : Apparatus | null = null

  ngOnInit(): void {
      this.route.paramMap.subscribe((params) => {
      this.competitionId = params.get('id') || "";
      this.competitionId = params.get('id') || "";
      this.loadData()
    });
  }

  loadData = () =>{
    //Get judging info
    this.scService.getLoggedJudgeInfo().subscribe({
      next: (response: JudgeJudgingInfo) => {
        this.judgingInfo = response;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }
}
