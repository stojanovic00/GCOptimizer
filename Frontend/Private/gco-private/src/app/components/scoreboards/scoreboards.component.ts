import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { mapAllAroundScores } from 'src/app/model/core/all-around-scoreboard-slot';
import { getApparatusName } from 'src/app/model/core/apparatus';
import { Apparatus } from 'src/app/model/core/apparatus';
import { ScoreboardBundle } from 'src/app/model/dto/scoreboard-bundle';
import { ScoringService } from 'src/app/services/scoring.service';

@Component({
  selector: 'app-scoreboards',
  templateUrl: './scoreboards.component.html',
  styleUrls: ['./scoreboards.component.css']
})
export class ScoreboardsComponent implements OnInit {

  competitionId: string = ""
  scoreBoardBundle : ScoreboardBundle | null = null;
  allAroundDisplayedColumns: string[] = ['place', 'contestant', 'organization'];
  teamDisplayedColumns: string[] = ['place', 'organization', 'team'];

  constructor(
    private readonly route : ActivatedRoute,
    private readonly scService : ScoringService,
  ){}

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      this.competitionId = params.get('id') || "";
      this.loadData();
    });
  }

  loadData = () => {
    this.scService.getScoreboards(this.competitionId).subscribe({
      next: (response: ScoreboardBundle) => {
          this.scoreBoardBundle = response;

          //Creating data structures for apparatus columns
          //All around
          this.allAroundDisplayedColumns = this.allAroundDisplayedColumns.concat(this.scoreBoardBundle.allAroundScoreboards[0].apparatuses.map(app => getApparatusName(app)));
          this.allAroundDisplayedColumns.push('total');
          this.scoreBoardBundle.allAroundScoreboards.forEach(scb =>{
            scb.slots.forEach(slot => {
              slot.apparatusScore = mapAllAroundScores(slot.scores);
            })
          })

          //Team
          this.teamDisplayedColumns = this.teamDisplayedColumns.concat(this.scoreBoardBundle.teamScoreboards[0].apparatuses.map(app => getApparatusName(app)));
          this.teamDisplayedColumns.push('total');


      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }

}
