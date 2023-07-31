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
import { WebSocketEventMessage } from '../../model/web-socket/web-socket-event-message';
import { ScoringEvent } from 'src/app/model/web-socket/scoring-event';

@Component({
  selector: 'app-judge-current-contestant',
  templateUrl: './judge-current-contestant.component.html',
  styleUrls: ['./judge-current-contestant.component.css']
})
export class JudgeCurrentContestantComponent implements OnInit {

rotationFinished = false;
judgingInfo: JudgeJudgingInfo | null = null;
currentContestant: ContestantScoring | null = null;

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

tempScoreSubmitted : boolean = false;

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
        //Gets info needed for web socket connection (apparatus and competitionId)
        this.judgingInfo = response;
        //OPENING WEB SOCKET!!!
        this.openWebSocket()
        this.loadCurrentContestant()
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }

  openWebSocket = () => {
    this.socket = new WebSocketService(this.judgingInfo?.apparatus ?? 0, this.judgingInfo?.competitionId!)
    this.socket.getEventListener().subscribe(event => {
      if (event.type == "message") {
        switch(event.data.event){
          case ScoringEvent.RetrievedContestantsTempScores:
            this.parseTempScoresResponse(event.data.response) 

            this.sendEvent(ScoringEvent.RetrievedContestantsTempScores);
            break;
          case ScoringEvent.RetrievedCanCalculate:
              this.canCalculateScore = event.data.response;
            break;
          case ScoringEvent.RetrievedScore:
            this.score = event.data.response;
            if(this.score?.submitted ?? false){
              this.contestantScored = true ;
              this.sendEvent(ScoringEvent.ScoredContestant);
            }
            break;
          case ScoringEvent.RetrievedNextCurrentApparatusContestant:
              if(!event.data.response.competingId){
                this.rotationFinished = true;
                break;
              }
              //Restarting all data for next contestant
              this.rotationFinished = false;
              this.currentContestant = event.data.response;
              this.tempScoreSubmitted = false;
              this.score = null;
              this.contestantScored = false;
              this.sendEvent(ScoringEvent.RetrievedNextCurrentApparatusContestant);
            break;
        }
      }
    });
  }

  sendEvent = (event: ScoringEvent) => {
    let socketMessage: WebSocketEventMessage = {
      event: event,
      competitionId: this.judgingInfo?.competitionId!,
      apparatus: this.judgingInfo?.apparatus!,
      ContestantId: this.currentContestant?.id!,
    }
    this.socket?.send(socketMessage)
  }

  loadCurrentContestant = () =>{
    this.scService.getCurrentApparatusContestant(this.judgingInfo?.competitionId!, this.judgingInfo?.apparatus ?? 0).subscribe({
      next: (response: ContestantScoring) => {
        if(!response.competingId){
          this.rotationFinished = true;
          return;
        }
        this.rotationFinished = false;
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
        this.parseTempScoresResponse(response)
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }



  parseTempScoresResponse = (tempScores: TempScore[]) =>{
          if(tempScores === null){
            this.dScoreTable.dataSource = []
            this.eScoreTable.dataSource = []
            return
          }

          if(tempScores.map(tmpScore => tmpScore.judge.id).includes(this.judgingInfo?.judge.id!)){
            this.tempScoreSubmitted = true;
          }
          
          this.dScoreTable.dataSource = tempScores.filter(tmpScore => tmpScore.type === undefined || tmpScore.type === ScoreType.D)
          this.eScoreTable.dataSource = tempScores
              .filter(tmpScore => tmpScore.type === ScoreType.E)
              .sort((a,b) => a.value - b.value)
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
      this.tempScoreSubmitted = true;

      this.sendEvent(ScoringEvent.TempScoreSubmitted);
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
      this.tempScoreSubmitted = true;

      this.sendEvent(ScoringEvent.TempScoreSubmitted);
    },
    error: (err: HttpErrorResponse) => {
      alert(err.error);
    }
  }); 
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



calculateScore = () =>{
    let scoreRequest: ScoreRequest = {
        apparatus: this.judgingInfo?.apparatus ?? 0,
        contestantId: this.currentContestant?.id!
    }
    this.scService.calculateScore(this.judgingInfo?.competitionId!, scoreRequest).subscribe({
      next: (response: Score) => {
        this.score = response;
        
        this.sendEvent(ScoringEvent.CalculatedScore)
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })

}

submitScore = () =>{
    this.scService.submitScore(this.judgingInfo?.competitionId!, this.score!).subscribe({
      next: (response: string) => {
        this.sendEvent(ScoringEvent.SubmittedScore);
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
          this.contestantScored = this.score.submitted;
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
