import { Apparatus } from "../core/apparatus";
import { JudgingPanelType } from "../core/judging-panel-type";
import { ScoreCalcMethod } from "../core/score-calc-method";
import { JudgeBasicInfo } from "./judge-basic-info";

export interface JudgeJudgingInfo{
    judge: JudgeBasicInfo;
    competitionId: string;
    apparatus: Apparatus;
    judgingPanelType: JudgingPanelType;
    calculationMethod: ScoreCalcMethod;
}
