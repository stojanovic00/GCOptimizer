import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Apparatus } from 'src/app/model/core/apparatus';
import { CurrentSessionInfo } from 'src/app/model/dto/current-session-info';
import { ScoringEvent } from 'src/app/model/web-socket/scoring-event';
import { WebSocketEventMessage } from 'src/app/model/web-socket/web-socket-event-message';
import { ScoringService } from 'src/app/services/scoring.service';
import { WebSocketService } from 'src/app/services/web-socket.service';

@Component({
  selector: 'app-competition-monitoring',
  templateUrl: './competition-monitoring.component.html',
  styleUrls: ['./competition-monitoring.component.css']
})
export class CompetitionMonitoringComponent implements OnInit {

  competitionId : string = "";
  currentSessionInfo: CurrentSessionInfo | null = null;
  
  private  socket : WebSocketService | null= null;
  constructor(
    private readonly route : ActivatedRoute,
    private readonly router : Router,
    private readonly scService : ScoringService,
  ){}

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      this.competitionId = params.get('id') || "";
      this.loadData();
    });
  }

  loadData = () => {
    this.getCurrentSessionInfo();
    this.openWebSocket();
  }


  openWebSocket = () => {
    this.socket = new WebSocketService(Apparatus.CompetitionAdmin, this.competitionId!)
    this.socket.getEventListener().subscribe(event => {
      if (event.type == "message") {
        switch(event.data.event){
          case ScoringEvent.RetrievedScore:
            this.getCurrentSessionInfo(); 
            break;
        }
      }
    });
  }


  sendEvent = (event: ScoringEvent) => {
    let socketMessage: WebSocketEventMessage = {
      event: event,
      competitionId: this.competitionId,
      apparatus: Apparatus.CompetitionAdmin,
      ContestantId: "",
    }
    this.socket?.send(socketMessage)
  }

  getCurrentSessionInfo = () =>{
    this.scService.getCurrentSessionInfo(this.competitionId).subscribe({
      next: (response: CurrentSessionInfo) => {
          this.currentSessionInfo = response;
          if(!this.currentSessionInfo.currentRotation){
            this.currentSessionInfo.currentRotation = 0;
          }

          if(this.currentSessionInfo.competitionFinished){
            this.finishCompetition();
          }

      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });


  }


  finishRotation = () =>{
    this.scService.finishRotation(this.competitionId).subscribe({
      next: (response: string) => {
        this.getCurrentSessionInfo();
        this.sendEvent(ScoringEvent.FinishedRotationOrSession);
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }

  finishSession = () =>{
    this.scService.finishSession(this.competitionId).subscribe({
      next: (response: string) => {
    this.scService.getCurrentSessionInfo(this.competitionId).subscribe({
      next: (response: CurrentSessionInfo) => {
          this.currentSessionInfo = response;
          if(!this.currentSessionInfo.currentRotation){
            this.currentSessionInfo.currentRotation = 0;
          }

          if(this.currentSessionInfo.competitionFinished){
            this.finishCompetition();
          }
          else
          {
            this.sendEvent(ScoringEvent.FinishedRotationOrSession);
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

  finishCompetition = () =>{
    this.scService.finishCompetition(this.competitionId).subscribe({
      next: (response: string) => {
        this.sendEvent(ScoringEvent.FinishedCompetition);
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }

  goToScoreboards = () =>{
      this.router.navigate(['sports-org/competition/' + this.competitionId + '/scoreboards']); 
  }
}
