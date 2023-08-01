import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import {JudgePanelService} from '../../services/judge-panel.service'
import { HttpErrorResponse } from '@angular/common/http';
import { Apparatus, ApparatusTable, getApparatusName } from 'src/app/model/core/apparatus';

@Component({
  selector: 'app-choose-panel',
  templateUrl: './choose-panel.component.html',
  styleUrls: ['./choose-panel.component.css']
})
export class ChoosePanelComponent implements OnInit {

  competitionId : string = ""

  unassignedApps : ApparatusTable = {
    displayedColumns : ["name"],
    dataSource : [],
    selectedRow : null 
  }
  getApparatusName = getApparatusName

  constructor(
    private readonly router: Router,
    private readonly route: ActivatedRoute,
    private readonly jpService: JudgePanelService,
  ){}

  ngOnInit(): void {
    let competitionId = "";
    this.route.paramMap.subscribe((params) => {
      this.competitionId = params.get('id') || "";
      this.loadTables(this.competitionId);
    });
  }


  loadTables = (competitionId : string) => {
    this.jpService.getAppsWithoutPanel(competitionId).subscribe({
      next: (response: Apparatus[]) => {
        if(!response){
          this.router.navigate(['sports-org/competition/'  + this.competitionId + '/monitoring']);
          return;
        }
        this.unassignedApps.dataSource = response
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    });
  }

  isApparatusSelected = (apparatus : Apparatus) : boolean => {
    return this.unassignedApps.selectedRow === apparatus;
  }

selectRow(row: Apparatus) {
  this.unassignedApps.selectedRow = row;
}

  createPanels = () => {
    let apparatusName = getApparatusName(this.unassignedApps.selectedRow!)
    this.router.navigate(['sports-org/competition/'  + this.competitionId + '/judging-panel/form/' + apparatusName]);
  }
}
