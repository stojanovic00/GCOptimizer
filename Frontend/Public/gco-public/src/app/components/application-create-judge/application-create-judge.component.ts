import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { getGenderName } from 'src/app/model/core/gender';
import { JudgeTable, Judge } from 'src/app/model/core/judge';
import { getLicenceTypeName } from 'src/app/model/core/licence-type';
import { CreateJudgeApplicationDto } from 'src/app/model/dto/create-judge-application-dto';
import { CompetitionService } from 'src/app/services/competition.service';
import { SprotsOrgService } from 'src/app/services/sprots-org.service';

@Component({
  selector: 'app-application-create-judge',
  templateUrl: './application-create-judge.component.html',
  styleUrls: ['./application-create-judge.component.css']
})
export class ApplicationCreateJudgeComponent implements OnInit {

  //View
  table : JudgeTable = {
    displayedColumns : ["fullName", "email", "licenceType", "licenceName", "gender"],
    dataSource : [], 
    selectedRow : null
  }
  getLicenceTypeName = getLicenceTypeName
  getGenderName = getGenderName

  competitionId : string = "";

  addedJudges : Judge[] = []
  addedTable : JudgeTable = {
    displayedColumns : ["fullName", "email", "licenceType", "licenceName", "gender"],
    dataSource : [], 
    selectedRow : null
  }

  constructor(
    private readonly route : ActivatedRoute,
    private readonly router : Router,
    private readonly soService : SprotsOrgService,
    private readonly compService : CompetitionService
  ) { }

  ngOnInit(): void {
    this.route.paramMap.subscribe((params) => {
      this.competitionId = params.get('id') || "";
    });

    this.loadTable();
  }

  loadTable = () => {
    this.soService.getJudges().subscribe({
      next: (response: Judge[]) => {
        this.table.dataSource = response;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }

  //View
  selectRow = (row: Judge) => {
    this.table.selectedRow = row;
  }
addApplication = () =>{
  let request : CreateJudgeApplicationDto = {
      judgeId : this.table.selectedRow?.delegationMember.id!
  }

    this.compService.addJudgeApplication(request, this.competitionId).subscribe({
      next: (response: string) => {
        this.addedJudges.push(this.table.selectedRow!)
        this.addedTable.dataSource = this.addedJudges.slice();
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
}

finishCreation = () =>
{
    this.router.navigate(['sports-org/competition/view']);
}


}
