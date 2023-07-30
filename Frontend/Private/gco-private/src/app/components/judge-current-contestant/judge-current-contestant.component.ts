import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { Apparatus } from 'src/app/model/core/apparatus';
import { ContestantScoring } from 'src/app/model/core/contestant-scoring';
import { JudgingPanelType } from 'src/app/model/core/judging-panel-type';
import { Score } from 'src/app/model/core/score';
import { ScoreType } from 'src/app/model/core/score-type';
import { TempScore, TempScoreTable } from 'src/app/model/core/temp-score';
import { JudgeJudgingInfo } from 'src/app/model/dto/judge-judging-info';
import { ScoreRequest } from 'src/app/model/dto/score-request';
import { ScoringService } from 'src/app/services/scoring.service';
import { WebSocketService } from 'src/app/services/web-socket.service';
import { WebSocketMessage } from '../../model/web-socket/web-socket-message';
import { ScoringEvent } from 'src/app/model/web-socket/scoring-event';

@Component({
  selector: 'app-judge-current-contestant',
  templateUrl: './judge-current-contestant.component.html',
  styleUrls: ['./judge-current-contestant.component.css']
})
export class JudgeCurrentContestantComponent implements OnInit {

judgingInfo: JudgeJudgingInfo | null = null;
currentContestant: ContestantScoring | null = null;
tempScores: TempScore[] = []

get IsDJudge():boolean{
    if(this.judgingInfo?.judgingPanelType === JudgingPanelType.DPanel){
      return true
    }
    return false;
}

get scoreType():ScoreType{
  if(this.IsDJudge){
    return ScoreType.D
  }
  else
  {
    return ScoreType.E
  }
}


//SCORING

scoreSubmitted : boolean = false;

//E SCORE
  eScoreForm = new FormGroup({
    eScore: new FormControl<number>(0, [Validators.required]),
  });

  public get EScore() {
    return this.eScoreForm.controls.eScore.value;
  }

//D SCORE
  dScoreForm = new FormGroup({
    dScore: new FormControl<number>(0, [Validators.required]),
  });

  public get DScore() {
    return this.dScoreForm.controls.dScore.value;
  }

// TABLES

eScoreTable :TempScoreTable = {
  displayedColumns :["judge", "value"],
  dataSource: [],
  selectedRow: null
}

dScoreTable :TempScoreTable = {
  displayedColumns :["judge", "value"],
  dataSource: [],
  selectedRow: null
}

  canCalculateScore : boolean = false;
  score :Score | null = null;
  contestantScored : boolean = false;

  private  socket : WebSocketService | null= null;
  constructor(
    private readonly route : ActivatedRoute,
    private readonly router : Router,
    private readonly scService : ScoringService,
  ){}



  public ngOnInit() {
    this.loadData()
}

public ngOnDestroy() {
  this.socket!.close();
}

  loadData = () =>{
    //Get judging info
    this.scService.getLoggedJudgeInfo().subscribe({
      next: (response: JudgeJudgingInfo) => {
        this.judgingInfo = response;
        //WARNING



          this.socket = new WebSocketService(this.judgingInfo?.apparatus ?? 0, this.judgingInfo?.competitionId) 
          this.socket.getEventListener().subscribe(event => {
          if(event.type == "message") {
            console.log(event)
          }
      });




        //WARNING
        this.loadCurrentContestant()
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
        this.getSCore();
        this.loadTempScores();
        this.checkCanCalculate();
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }

  loadTempScores = () =>{
    let scoreRequest: ScoreRequest = {
        apparatus: this.judgingInfo?.apparatus ?? 0,
        contestantId: this.currentContestant?.id!
    }
    this.scService.getContestantsTempScores(this.judgingInfo?.competitionId!, scoreRequest).subscribe({
      next: (response: TempScore[]) => {
        if(response === null){
          this.dScoreTable.dataSource = []
          this.eScoreTable.dataSource = []
          return
        }

        this.tempScores = response

        if(this.tempScores.map(tmpScore => tmpScore.judge.id).includes(this.judgingInfo?.judge.id!)){
          this.scoreSubmitted = true;
        }
        
        this.dScoreTable.dataSource = this.tempScores.filter(tmpScore => tmpScore.type === undefined || tmpScore.type === ScoreType.D)
        this.eScoreTable.dataSource = this.tempScores
            .filter(tmpScore => tmpScore.type === ScoreType.E)
            .sort((a,b) => a.value - b.value)


      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }

  checkCanCalculate = () => {
    let scoreRequest: ScoreRequest = {
        apparatus: this.judgingInfo?.apparatus ?? 0,
        contestantId: this.currentContestant?.id!
    }
    this.scService.canCalculateScore(this.judgingInfo?.competitionId!, scoreRequest).subscribe({
      next: (response: boolean) => {
        this.canCalculateScore = response;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
 }


  submitEScore = () => {
    let tempScore : TempScore = {
      apparatus: this.judgingInfo?.apparatus ?? 0,
      competitionId: this.judgingInfo?.competitionId!,
      contestant: this.currentContestant!,
      judge: this.judgingInfo?.judge!,
      type: this.scoreType!,
      value: this.EScore!,
    }
    this.scService.submitTempScore(this.judgingInfo?.competitionId!, tempScore).subscribe({
    next: (response: string) => {
      this.scoreSubmitted = true;
      let socketMessage : WebSocketMessage = {
          event: ScoringEvent.TempScoreSubmitted,
          competitionId: this.judgingInfo?.competitionId!,
          apparatus: this.judgingInfo?.apparatus!
      }
      this.socket?.send(socketMessage) 
    },
    error: (err: HttpErrorResponse) => {
      alert(err.error);
    }
  }); 
  }


  submitDScore = () => {
    let tempScore : TempScore = {
      apparatus: this.judgingInfo?.apparatus ?? 0,
      competitionId: this.judgingInfo?.competitionId!,
      contestant: this.currentContestant!,
      judge: this.judgingInfo?.judge!,
      type: this.scoreType!,
      value: this.DScore!,
    }
    this.scService.submitTempScore(this.judgingInfo?.competitionId!, tempScore).subscribe({
    next: (response: string) => {
      this.scoreSubmitted = true;
    },
    error: (err: HttpErrorResponse) => {
      alert(err.error);
    }
  }); 
  }

calculateScore = () =>{
    let scoreRequest: ScoreRequest = {
        apparatus: this.judgingInfo?.apparatus ?? 0,
        contestantId: this.currentContestant?.id!
    }
    this.scService.calculateScore(this.judgingInfo?.competitionId!, scoreRequest).subscribe({
      next: (response: Score) => {
        this.score = response;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })

}

submitScore = () =>{
    this.scService.submitScore(this.judgingInfo?.competitionId!, this.score!).subscribe({
      next: (response: string) => {
          this.contestantScored = true;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })

}

getSCore = () =>{
    let scoreRequest: ScoreRequest = {
        apparatus: this.judgingInfo?.apparatus ?? 0,
        contestantId: this.currentContestant?.id!
    }
    this.scService.getScore(this.judgingInfo?.competitionId!, scoreRequest).subscribe({
      next: (response: Score) => {
        if(response !== null){
          this.score = response;
          this.contestantScored = true;
        }
      },
      error: (err: HttpErrorResponse) => {
      }
    })
}

isFirstOrLastNScore(scoreIndex: number, dataSourceLength: number): boolean {
  return scoreIndex < this.judgingInfo?.calculationMethod.scoreDeductionNum! || scoreIndex >= dataSourceLength - this.judgingInfo?.calculationMethod.scoreDeductionNum!;
}



}
