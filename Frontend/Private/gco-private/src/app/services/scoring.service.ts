import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment'
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { Apparatus } from '../model/core/apparatus';
import { JudgeJudgingInfo } from '../model/dto/judge-judging-info';
import { ContestantScoringDto } from '../model/dto/contestant-scoring-dto';
import { ContestantScoring } from '../model/core/contestant-scoring';
import { TempScore } from '../model/core/temp-score';
import { Score } from '../model/core/score';
import { ScoreRequest } from '../model/dto/score-request';

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

  submitTempScore = (compId: string, tempScore: TempScore): Observable<string> => {
    let path = this.path + "/competition/" + compId + "/temp-score";
    return this.http.post<string>(path, tempScore);
  }

  getContestantsTempScores = (compId: string, scoreRequest: ScoreRequest): Observable<TempScore[]> => {
    let path = this.path + "/competition/" + compId + "/temp-score?contestantId=" + scoreRequest.contestantId + "&apparatus=" + scoreRequest.apparatus;
    return this.http.get<TempScore[]>(path);
  }

  canCalculateScore = (compId: string, scoreRequest: ScoreRequest): Observable<boolean> => {
    let path = this.path + "/competition/" + compId + "/score/can-calculate?contestantId=" + scoreRequest.contestantId + "&apparatus=" + scoreRequest.apparatus;
    return this.http.get<boolean>(path);
  }

  calculateScore = (compId: string, scoreRequest: ScoreRequest): Observable<Score> => {
    let path = this.path + "/competition/" + compId + "/score/calculate?contestantId=" + scoreRequest.contestantId + "&apparatus=" + scoreRequest.apparatus;

    return this.http.get<Score>(path);
  }

  submitScore = (compId: string, score: Score): Observable<string> => {
    let path = this.path + "/competition/" + compId + "/score";
    return this.http.post<string>(path, score);
  }

  getScore = (compId: string, scoreRequest: ScoreRequest): Observable<Score> => {
    let path = this.path + "/competition/" + compId + "/score?contestantId=" + scoreRequest.contestantId + "&apparatus=" + scoreRequest.apparatus;


    return this.http.get<Score>(path);
  }
}
