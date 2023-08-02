import { Injectable } from '@angular/core';
import { environment } from '../../environments/environment'
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { CreateJudgePanelsForApparatusRequest } from '../model/dto/create-judge-panels-for-apparatus-request';
import {  CreateJudgePanelsForApparatusResponse  } from '../model/dto/create-judge-panels-for-apparatus-response';
import { Apparatus } from '../model/core/apparatus';
import { JudgeBasicInfo } from '../model/dto/judge-basic-info';
import { ScoreCalcMethod } from '../model/core/score-calc-method';

@Injectable({
  providedIn: 'root'
})
export class JudgePanelService {

  path: string = environment.scoringPath + "/judging-panel";

  constructor(
    private readonly http: HttpClient
    ) { }


  getAppsWithoutPanel = (compId : string): Observable<Apparatus[]> => {
    let path = this.path + "/competition/" + compId + "/unassigned";
    return this.http.get<Apparatus[]>(path);
  }

  createJudgePanelsForApparatus = (dto : CreateJudgePanelsForApparatusRequest): Observable<CreateJudgePanelsForApparatusResponse> => {
    return this.http.post<CreateJudgePanelsForApparatusResponse>(this.path, dto);
  }
  getAssignedJudges = (compId : string): Observable<JudgeBasicInfo[]> => {
    let path = this.path + "/judge/competition/" + compId;
    return this.http.get<JudgeBasicInfo[]>(path);
  }

  assignJudgeToPanel = (judge : JudgeBasicInfo, panelId: string): Observable<string> => {
    let path = this.path + "/" + panelId + "/judge"
    return this.http.post<string>(path, judge);
  }

  assignScoreCalcMethodToPanel = (method : ScoreCalcMethod, panelId: string): Observable<string> => {
    let path = this.path + "/" + panelId + "/score-calc-method"
    return this.http.post<string>(path, method);
  }
}