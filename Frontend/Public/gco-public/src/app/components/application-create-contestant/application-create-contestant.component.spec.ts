import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ApplicationCreateContestantComponent } from './application-create-contestant.component';

describe('ApplicationCreateContestantComponent', () => {
  let component: ApplicationCreateContestantComponent;
  let fixture: ComponentFixture<ApplicationCreateContestantComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ApplicationCreateContestantComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ApplicationCreateContestantComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
