import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Schedule } from 'src/app/model/core/schedule';
import { ScheduleSessionView } from 'src/app/model/dto/schedule-session-view';
import { ScheduleDtoToScheduleView } from 'src/app/model/mappers/schedule-mapper';
import { CompetitionService } from 'src/app/services/competition.service';
import { ScheduleService } from 'src/app/services/schedule.service';

@Component({
  selector: 'app-schedule-view',
  templateUrl: './schedule-view.component.html',
  styleUrls: ['./schedule-view.component.css']
})
export class ScheduleViewComponent implements OnInit {

  sessionViews : ScheduleSessionView[] = []

  constructor(
    private readonly route : ActivatedRoute,
    private readonly router : Router,
    private readonly schService : ScheduleService,
    private readonly compService : CompetitionService,
  ){}

  competitionId : string = ""

  ngOnInit(): void {
      this.route.paramMap.subscribe((params) => {
      this.competitionId = params.get('id') || "";
      this.loadData(this.competitionId)
    });
  }

  loadData = (compId: string) =>{
      this.schService.getByCompetitionId(compId).subscribe({
      next: (response: Schedule) => {
          this.sessionViews = ScheduleDtoToScheduleView(response);
      },
      error: (err: HttpErrorResponse) => {
          this.sessionViews = [];
      }
    }); 
  }

  startCompetition = () =>{
      this.compService.startCompetition(this.competitionId).subscribe({
      next: (response: string) => {
        this.router.navigate(['sports-org/competition/'  + this.competitionId + '/judging-panel/unassigned']);
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error)
      }
    }); 
  }

}
