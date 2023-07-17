import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Judge, JudgeTable } from 'src/app/model/core/judge';
import { SprotsOrgService } from 'src/app/services/sprots-org.service';
import { getLicenceTypeName  } from '../../model/core/licence-type';

@Component({
  selector: 'app-judges-view-register',
  templateUrl: './judges-view-register.component.html',
  styleUrls: ['./judges-view-register.component.css']
})

export class JudgesViewRegisterComponent implements OnInit {

  table : JudgeTable = {
    displayedColumns : ["fullName", "email", "licenceType", "licenceName"],
    dataSource : [], 
    selectedRow : null
  }
  getLicenceTypeName = getLicenceTypeName

  constructor(
    private readonly soService : SprotsOrgService
  ) { }

  ngOnInit(): void {
    this.soService.getJudges().subscribe({
      next: (response: Judge[]) => {
        this.table.dataSource = response;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }

  selectRow = (row: Judge) => {
    this.table.selectedRow = row;
  }
}
