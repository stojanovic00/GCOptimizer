import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
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
    private readonly schService : ScheduleService,
  ){}

  ngOnInit(): void {
      this.route.paramMap.subscribe((params) => {
      let compId = params.get('id') || "";
      this.loadData(compId)
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

}
