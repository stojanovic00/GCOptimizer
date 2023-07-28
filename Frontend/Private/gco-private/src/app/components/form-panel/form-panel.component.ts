import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { FormGroup, FormControl, Validators } from '@angular/forms';
import { Router, ActivatedRoute } from '@angular/router';
import { parseApparatus } from 'src/app/model/core/apparatus';
import { JudgeApplication } from 'src/app/model/core/judge-application';
import { getLicenceTypeName } from 'src/app/model/core/licence-type';
import { ScoreCalcMethod } from 'src/app/model/core/score-calc-method';
import { CreateJudgePanelsForApparatusRequest } from 'src/app/model/dto/create-judge-panels-for-apparatus-request';
import { CreateJudgePanelsForApparatusResponse } from 'src/app/model/dto/create-judge-panels-for-apparatus-response';
import { JudgeBasicInfo, JudgeBasicInfoTable, parseJudgeBasicInfo } from 'src/app/model/dto/judge-basic-info';
import { CompetitionService } from 'src/app/services/competition.service';
import { JudgePanelService } from 'src/app/services/judge-panel.service';

@Component({
  selector: 'app-form-panel',
  templateUrl: './form-panel.component.html',
  styleUrls: ['./form-panel.component.css']
})
export class FormPanelComponent implements OnInit {

  apparatusName :string = "";
  competitionId :string = "";
  dPanelId :string = "";
  ePanelId :string = "";

  availableJudges: JudgeBasicInfo[] = []

  availableTable : JudgeBasicInfoTable = {
    displayedColumns : ["fullName", "email", "licenceType", "licenceName", "sportsOrg"],
    dataSource : [], 
    selectedRow : null
  }
  getLicenceTypeName = getLicenceTypeName

  assignedDJudges: JudgeBasicInfo[] = []

  assignedDTable : JudgeBasicInfoTable = {
    displayedColumns : ["fullName", "email", "licenceType", "licenceName", "sportsOrg"],
    dataSource : [], 
    selectedRow : null
  }

  assignedEJudges: JudgeBasicInfo[] = []

  assignedETable : JudgeBasicInfoTable = {
    displayedColumns : ["fullName", "email", "licenceType", "licenceName", "sportsOrg"],
    dataSource : [], 
    selectedRow : null
  }

  deductionForm = new FormGroup({
    deductionNumber: new FormControl<number>(0, [Validators.required]),
  });

  public get DeductionNumber() {
    return this.deductionForm.controls.deductionNumber.value;
  }

  constructor(
    private readonly router: Router,
    private readonly route: ActivatedRoute,
    private readonly jpService: JudgePanelService,
    private readonly compService: CompetitionService,
  ){}

  ngOnInit(): void {
    let competitionId = "";
    this.route.paramMap.subscribe((params) => {
      this.competitionId = params.get('id') || "";
      this.apparatusName = params.get('apparatus') || "";
      this.loadData(this.competitionId);
    });
  }


  loadData = (competitionId : string) => {

    let dto : CreateJudgePanelsForApparatusRequest ={
      apparatus : parseApparatus(this.apparatusName)!,
      competitionId: this.competitionId
    }

    // Create panels
    this.jpService.createJudgePanelsForApparatus(dto).subscribe({
      next: (response: CreateJudgePanelsForApparatusResponse) => {

        this.dPanelId = response.dPanelId;
        this.ePanelId = response.ePanelId;
        
        this.assignedDJudges = []
        this.assignedEJudges = []
        this.loadJudgesData()
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }


loadJudgesData = () => {
        //Get assigned judges
        this.jpService.getAssignedJudges(this.competitionId).subscribe({
          next :(assignedJudges: JudgeBasicInfo[]) => {
              if(!assignedJudges) assignedJudges = [] //Resolves bug when array is empty
              //Get applications
              this.compService.getJudgeApplications(this.competitionId).subscribe({
                next :(response: JudgeApplication[]) => {
                    //Calcualte available judges
                    let appliedJudges: JudgeBasicInfo[] = response.map(app => parseJudgeBasicInfo(app.judge))         
                    this.availableJudges = appliedJudges.filter(appliedJudge => !assignedJudges.find(assJudge => assJudge.id === appliedJudge.id))
                    this.availableTable.dataSource = this.availableJudges.slice();
                },
                error: (err: HttpErrorResponse) =>{
                  alert(err.error);
                }
              })
          },
          error: (err: HttpErrorResponse) =>{
            alert(err.error);
          }
        })
}



assignDJudge = () => {
  this.jpService.assignJudgeToPanel(this.availableTable.selectedRow!, this.dPanelId).subscribe({
    next :(response: string) => {
        this.assignedDJudges.push(this.availableTable.selectedRow!);
        this.assignedDTable.dataSource = this.assignedDJudges.slice();
        this.loadJudgesData()
    },
    error: (err: HttpErrorResponse) =>{
      alert(err.error);
    }
  })
}

assignEJudge = () => {
  this.jpService.assignJudgeToPanel(this.availableTable.selectedRow!, this.ePanelId).subscribe({
    next :(response: string) => {
        this.assignedEJudges.push(this.availableTable.selectedRow!);
        this.assignedETable.dataSource = this.assignedEJudges.slice();
        this.loadJudgesData()
    },
    error: (err: HttpErrorResponse) =>{
      alert(err.error);
    }
  })
}


selectRow = (row: JudgeBasicInfo) => {
  this.availableTable.selectedRow = row;
}


assignDScoreCalculationMethod = () =>{
    let method: ScoreCalcMethod ={
      scoreDeductionNum : 0
    } 

  this.jpService.assignScoreCalcMethodToPanel(method, this.dPanelId).subscribe({
    next :(response: string) => {
      
    },
    error: (err: HttpErrorResponse) =>{
      alert(err.error);
    }
  })
}


assignEScoreCalculationMethod = () =>{
    let method: ScoreCalcMethod ={
      scoreDeductionNum : this.DeductionNumber!
    } 

  this.jpService.assignScoreCalcMethodToPanel(method, this.ePanelId).subscribe({
    next :(response: string) => {
      this.router.navigate(['sports-org/competition/' + this.competitionId + '/judging-panel/unassigned']); 
    },
    error: (err: HttpErrorResponse) =>{
      alert(err.error);
    }
  })
}

}
