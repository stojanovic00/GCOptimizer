+ Dodeli sudije na svakoj spravi:
    + ides redom za svaku spravu
        + napravis panel 
            + dodelis d sudije
            + dodelis e sudije
            + nacin ocenjivanja (koliko se odbija e ocena)
            + ujedno se prave nalozi za sudije i salju im se sifre na mail
            + ako imades vremena napravi menjanje sifre

+ startuj takmicenje
    + dobavi schedule preko comp id 
        + isparsiras ga za potrebe scoringa
        + snimis
    + kliknes start dugme i tad se prva sesija startuje
        + obaveste se sve sudije na frontu da citaju takmicare
            + ucita sve takmicare na spravi i onog ko je na redu
                + onaj ko je na redu je onaj koji prvi naidje u listi, da takmici tu spravu i da nema tu spravu u listi scored apparatuses
            + ovo je isto i za D i za E sudije tj sudiji ce da se uradi jedan ngIf da mu da ili E ili D ili oba interfejsa 

        - Ocenjivanje:
            - D sudija da svoju ocenu i submituje
            - E sudija da svoju ocenu i submituje
            - kad god neko submituje D sudija bude pingovan i on salje request da refreshuje ocene
            - kad sve ocene stignu pokrene se algoritam za racunanje prosecne E i D ocene i dobija se finalna ocena
            - to se onda submituje i nastupa sledeci takmicar

        - svaki put kad se submituje final score adminov front proveri jesu li svi takmicari zavrsili trenutnu rotaciju
            - ako jesu moze da izvrsi promeu rotacije gde su samo azurira rotacija sesije
            - ako su zavrsene sve rotacije zapocinje nova sesija, ako su zavrsene sve sesije pokrece se pravljenje rang liste
                - rang liste su samo final scoreovi grupisani po age category i sortirani po visini ocene
                    - prednost ima onaj sa visom E ocenom
                    - tiebrake = true nema deljenja mesta, false moze se deli mesto 


Eventovi:
  - Judge side:
    - Submited temp score
      - Refresh temp scores(judge)
    - Submited score
      - Get current contestant(judge) - if null = no more contestants to score
      - check rotation finish(admin)
      - check session finish(admin)
  - Admin side
    - finish rotation: pristupi sesiji i poveca trenutnu rotaciju
      - get all contestants(judge)
      - get current contestant(judge)
    - finish session: trenutnu sesiju stavi na finished
      - check competition finish - ako su sve sesije finished
      - get all contestants(judge)
      - get current contestant(judge)



Endpoints:
    - Preparation DONE

        + Start + 

        + get apparatuses that need panel +
        + create panels for apparatus +
        + get judges(applications) - from app +
        + get assigned judges - from scoring +
        + Assign judge +
            + register judge account +
        + Assign score calculation +

        ? Change password


    - Live DONE
        + get logged judge (judge info + panel info)

        + Get contestants for apparatus <-- new rot
        + Get current contestant for apparatus <-- final score submitted <-- new rot

        + Submit temp score --> temp score submitted
        + get temp scores for contestant <-- temp score submitted
        + get final score for contestant <-- temp score submitted
        + Submit final score --> final score submitted


        + Finish rotation
        + Finish session

        + Check rotation finish 
        + Check session finish
        + Check competition finish

    - Scoreboards
        - Finish competition:
          - Generate scoreboard
          - Generate team scoreboard

        - Get scoreboard
        - Get team scoreboard
