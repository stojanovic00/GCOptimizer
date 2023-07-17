import { HttpErrorResponse } from '@angular/common/http';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { SportsOrg } from 'src/app/model/core/sports-org';
import { SprotsOrgService } from 'src/app/services/sprots-org.service';

@Component({
  selector: 'app-sports-org-view',
  templateUrl: './sports-org-view.component.html',
  styleUrls: ['./sports-org-view.component.css']
})
export class SportsOrgViewComponent implements OnInit {


  sportsOrg! : SportsOrg;

  get address () :string {
    return this.sportsOrg!.address.country! + " " + this.sportsOrg!.address.city! + ", " + this.sportsOrg!.address.street! + " " + this.sportsOrg!.address.streetNumber!;
  }

  constructor(
    private readonly soService: SprotsOrgService,
    private readonly router: Router,
  ) { }

  ngOnInit(): void {
    this.soService.getLoggedIn().subscribe({
      next: (response: SportsOrg) => {
        this.sportsOrg = response;
      },
      error: (err: HttpErrorResponse) => {
        alert(err.error);
      }
    })
  }

  goToCreateCompetition = () => {
    this.router.navigate(['sports-org/competition']);
  }
}
