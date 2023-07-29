import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment'
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Apparatus } from '../model/core/apparatus';
import { JudgeJudgingInfo } from '../model/dto/judge-judging-info';
import { ContestantScoringDto } from '../model/dto/contestant-scoring-dto';
import { ContestantScoring } from '../model/core/contestant-scoring';

@Injectable({
  providedIn: 'root'
})
export class ScoringService {

  path: string = environment.scoringPath

  constructor(
    private readonly http: HttpClient
    ) { }



  getLoggedJudgeInfo = (): Observable<JudgeJudgingInfo> => {
    let path = this.path + "/judge";
    return this.http.get<JudgeJudgingInfo>(path);
  }

  getCurrentContestantsForApparatus = (compId: string, apparatus: Apparatus): Observable<ContestantScoringDto[]> => {
    let path = this.path + "/competition/" + compId + "/contestant?apparatus=" + apparatus;
    return this.http.get<ContestantScoringDto[]>(path);
  }

  getCurrentApparatusContestant = (compId: string, apparatus: Apparatus): Observable<ContestantScoring> => {
    let path = this.path + "/competition/" + compId + "/contestant/current?apparatus=" + apparatus;
    return this.http.get<ContestantScoring>(path);
  }
}
