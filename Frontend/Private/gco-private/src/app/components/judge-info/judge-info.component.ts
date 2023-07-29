import { Component, OnInit } from '@angular/core';
import { JudgeJudgingInfo } from '../../model/dto/judge-judging-info';
import { ScoringService } from '../../services/scoring.service';
import { HttpErrorResponse } from '@angular/common/http';
import { JudgingPanelType } from 'src/app/model/core/judging-panel-type';
import { getLicenceTypeName } from 'src/app/model/core/licence-type';

@Component({
  selector: 'app-judge-info',
  templateUrl: './judge-info.component.html',
  styleUrls: ['./judge-info.component.css']
})
export class JudgeInfoComponent implements OnInit {


  judgingInfo: JudgeJudgingInfo | null = null;
  getLicenceTypeName = getLicenceTypeName
  public get JudgeType() :string {
    if(this.judgingInfo?.judgingPanelType == JudgingPanelType.DPanel){
      return "D";
    }
    else{
      return "E";
    }
    return "";
  }

  constructor(
    private readonly scService: ScoringService,
  ) { }

  ngOnInit(): void {
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
