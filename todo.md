- Dodeli sudije na svakoj spravi:
    - ides redom za svaku spravu
        - dodelis d sudije
        - dodelis e sudije
        - nacin ocenjivanja (koliko se odbija e ocena)
        - ujedno se prave nalozi za sudije i salju im se sifre na mail
        - ako imades vremena napravi menjanje sifre

- startuj takmicenje
    - dobavi schedule preko comp id 
        - isparsiras ga za potrebe scoringa
        - snimis
    - kliknes start dugme i tad se prva sesija startuje
        - obaveste se sve sudije na frontu da citaju takmicare
            - ucita sve takmicare na spravi i onog ko je na redu
                - onaj ko je na redu je onaj koji prvi naidje u listi, da takmici tu spravu i da nema tu spravu u listi scored apparatuses
            - ovo je isto i za D i za E sudije tj sudiji ce da se uradi jedan ngIf da mu da ili E ili D ili oba interfejsa 

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



    - Kako se cita sesija?
        - Ucita sesiju i vidi koja je trenutna rotacija
        - Tada uradi huga buga matematiku i skonta na osnovu redosleda  

        KONJA SUDI ROTACIJA 0 - dobavi sve koji imaju startin app konja 1 + (4 - 0) % 4  = 1 (vaistinu)
        KONJA SUDI ROTACIJA 1 - dobavi sve koji imaju startin app parter (1 + (4 -1)) % 4 = 0 (vaistinu) 
        KONJA SUDI ROTACIJA 2 - dobavi sve koji imaju startin app preskok (1 + (4 - 2)) % 4 = 3 (vaistinu)
        KONJA SUDI ROTACIJA 3 - dobavi sve koji imaju startin app preskok (1 + (4 - 3)) % 4 = 2 (vaistinu)

koga citam> : (starting app index + (broj sprava -  rotacija)) % broj sprava

parter konj karike preskok






Endpoints:
    - Start + 

    - get judges - from app
    - get assigned judges - from 
    - Assign judge 
    - Assign score calculation 
    - Change password?

    - Get contestants for apparatus
    - Get current contestant for apparatus

    - Submit temp score
    - Save final score
    - Check rotation finish
    - Check session finish
    - Check competition finish

    - Generate scoreboard
    - Generate team scoreboard