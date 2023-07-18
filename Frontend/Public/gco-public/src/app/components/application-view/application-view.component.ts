import { Component, OnInit } from '@angular/core';
import { ContestantApplication, ContestantApplicationTable } from 'src/app/model/core/contestant-application';
import { formatAgeCategory }from  '../../model/core/age-category'
import { formatApparatusAnnouncements }from  '../../model/core/apparatus-announcement'
import { CompetitionService } from 'src/app/services/competition.service';
import { ActivatedRoute } from '@angular/router';
import { HttpErrorResponse } from '@angular/common/http';
import { JudgeApplication, JudgeApplicationTable } from '../../model/core/judge-application';
import { getGenderName } from 'src/app/model/core/gender';
import { getLicenceTypeName } from 'src/app/model/core/licence-type';


@Component({
  selector: 'app-application-view',
  templateUrl: './application-view.component.html',
  styleUrls: ['./application-view.component.css']
})
export class ApplicationViewComponent implements OnInit {

tableApplicationsContestant : ContestantApplicationTable = {
    displayedColumns : ["fullName", "organisation", "teamNumber", "ageCategory", "apparatuses"],
    dataSource : [], 
    selectedRow : null
}

formatAgeCategory = formatAgeCategory 
formatApparatusAnnouncements = formatApparatusAnnouncements

tableApplicationsJudge : JudgeApplicationTable = {
    displayedColumns : ["fullName", "email", "licenceType", "licenceName", "gender"],
    dataSource : [], 
    selectedRow : null
}

getLicenceTypeName = getLicenceTypeName
getGenderName = getGenderName

  constructor(
    private readonly compService : CompetitionService,
    private readonly route : ActivatedRoute,
  ) { }

  ngOnInit(): void {
    let competitionId = "";
    this.route.paramMap.subscribe((params) => {
      competitionId = params.get('id') || "";
      this.loadTables(competitionId);
    });

  }

  loadTables = (competitionId : string) => {
    this.compService.getContestantApplications(competitionId).subscribe({
      next: (response: ContestantApplication[]) => {
        this.tableApplicationsContestant.dataSource = response;

        this.compService.getJudgeApplications(competitionId).subscribe({
          next: (response: JudgeApplication[]) => {
            this.tableApplicationsJudge.dataSource = response;
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

}
