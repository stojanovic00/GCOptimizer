## Pitanja u vezi organizacije

- da li postoji ogranicenje koliko takmicara jedan klub moze da prijavi
    - da li je to ogranicenje na nivou kategorije?
    - da li postoji neki prioritet? 
        - da ne moze jedan klub da prijavi 20 takmicara, pa da ostali klubovi uopste ne mogu da ucestvuju? 

- administrator moze da odjavi nasilno nekog
    - on ce videti live ko se prijavljuje i onda moze da se doradi


- minimalno jedan trener mora da bude prijavljen?
    - moze cak i da se ne prijavi trener

- ako se na klupskom takmicenju prijavi nacionalni tim, da li tada taj tim tretirati kao da je klub
    - samo u nazivu da pise nacionalni tim bugarske na primer

- 5 5 3 format ekipa
    - 3 najbolja visebojca iz 3 razlicite kategorije(obavezno)
        - misli se na score u viseboju

- koje informacije bi se unosile ako bi u pitanju bila drzava
    - iste, samo bi umesto naziva kluba pisalo npr "Nacionalni tim Slovenije"

- zbog cega se unosi godiste ako vec postoji polje za kategorije
    - da li onda treba omoguciti korigovanje kategorije ako takmicar omasi kategoriju
        - da
    - nekad trener moze da zatrazi da se takmicar takmici u visim kategorijama
        - moze samo na visu kategoriju

- naznaciti ako klub nije obezbedio sudiju

- sta ako broj prijavljenih sudija nije dovoljan da se odrzi takmicenje
    - kako se dobavljaju sudije?
        - automatski posao oko sudija se zavrsava kada se formira spisak prijavljeinh sudija, ostalo je sve rucno
        - onda se rucno dodaju sudije i rasporedjuju po spravama


- formiranje rasporeda treba da izgleda tako da se unesu parametri, zatim administrator vidi raspored i onda moze opet da unosi parametre, dok ne dobije raspored koji mu najvise odgovara
    - format 3 3 ili 2 2 se stavlja kao parametar kada se formira raspored

- da li je proglasenje uvek na kraju turnusa ili moze za sve da se odrzi na samom kraju
    - moze i na samom kraju
    - da li treba obezbediti jos prostora za neko finalno proglasenje
- ako se broj takmicara jedne kategorije prelije u dva turnusa, da li se onda njihovo proglasenje vrsi kada svi takmicari odtakmice
    - da



## Pitanja u vezi sudjenja

!!!!!!!!!!!!!!!!!!!!!!!!!
- U KOM TRENUTKU SE DEFINISE BROJ D I E SUDIJA NA SPRAVAMA
    - odrzava se na opstem zagrevanju(moze i ranije, ali kada se zavrse svi tehnicki problemi)
    - kako se racuna E ocena zadaje se neposredno pre takmicenja
        - odbacivanje i racunanje srednje vrednosti
    - D ocena se jedna upisuje
        - moze se definisati hoce li biti 1, 2... D sudije i da li ce biti supervizor (on ima zaseban uredjaj)
            - zavrsi se izvedba, supervizor posalje svoju D, to vide D ziri i odna se dodje do konsenzusa uo finalnoj D oceni
        - supervizor vidi sve (i D i E ocene) i ima mogucnost da pregleda i video materijal (nije bitno za sad)


- ako dodje do greske ili zalbe trenera, samo supervizor ili administrator da promeni ocenu uz dozvolu!!!
!!!!!!!!!!!!!!!!!!!!!!!!!


- kada su jednaki bodovi sta raditi?
    - imas fajl koji to definise 
    - za preskok se gleda malo drugacije


- Jedini slučaj kada se konačna ocena koja je prosleđena glavnom računaru može promeniti je ako trener iznese žalbu na suđenje  i traži ponovno većanje oko ocene. Ako sudije uvaže žalbu ocena se može promeniti i nju vrši administrator takmičenja. 
    - trebalo bi da to pravo ima samo jedna osoba koja je apsolutno neutralna

-Postoji slučaj (na manjim takmičenjima) kada se samo jedan sudija nalazi na spravi. Tada on ima ulogu i E i D sudije i sekretara. Ukoliko nema sekretara, D sudije preuzimaju njegova zaduženja.

- Na koju decimalu se zaokružuje krajnja ocena?


## Pitanja u vezi kvalifikacija za dalja takmicenja

- Da li i u pojedinačno finale višeboja mogu da se kvalifikuju maksimalno dva takmičara iz svakog kluba/države uključujući rezerve?

- Da li se broj rezervi određuje prilikom definisanja formata takmičenja ili to administrator odlučuje na samom kraju kvalifikacionog takmičenja?

- Da li postoje rezerve i za finala višeboja i ekipna finala? Ako da, koliko rezervi se određuje?

- Da li se pojedinačna, ekipna i finala po spravama odvijaju u posebnim danima ili se mogu nabiti u manje dana?
    - da li ako se probije vremensko ogranicenje organizatora treba onemoguciti pravljenje takmicenja sa zeljenim formatom

- Budući da u ekipnom finalu učestvuje 8 timova, da li to znači da se finale mora organizovati iz dva turnusa po 4 tima? Kako bi se onda određivalo u kom turnusu se koji tim takmiči? Da li se možda redosled određuje na osnovu plasmana ekipa u kvalifikacijama?

## Predlozi

- na kraju se mogu izabrati dva ili vise takmicenja i na osnovu tih rezultata da se sracuna apsoultni pobednik
- za takmičare koji nastupaju na manje sprava, softver bi mogao da proba da ih smesti u početnu rotaciju u kojoj bi najmanje morali da čekaju na svoj nastup na svim spravama, da se ne bi “ohladili” između sprava 

- Može se omogućiti naknadno pomeranje rasporeda u slučaju nezgode ili nekog neočekivanog tehničkog problema, kao i ukoliko administrator takmičenja uoči da je ipak potrebno više vremena za jednu rotaciju. Pomeranje bi se vršilo navođenjem vremenskog intervala za koji da se pomeri čitav
raspored ili definisanjem novog vremena potrebnog za rotaciju, gde bi se na osnovu njega iznova iskalkulisao raspored.

- Dozvoliti administratoru takmičenja da ručno doradi raspored nakon što je sistem automatski izgenerisao raspored. Tačne promene koje je administrator odradio bile bi zabeležene u sistemu, radi sprečavanja zloupotrebe ove opcije.
   


## Sudjenje

$ proveriti jos jednom citav postupak ocenjivanja takmicara!!!

$ koji sve tipovi penalizacija postoje, da li ih treba ubaciti u neki enum ili ih ostaviti da se rucno navedu
    $ proveri jos sta sve treba da se nalazi na interfejsima, koliko toga treba da se poklapa sa slikama koje je Deja poslao


$ penalizacije
    $ generalna za elemente
        $ dva polja, jedno vezano samo za broj elemenata, a drugo za opstu penalizaciju
    $ vreme (parter)
    $ linija 1 (parter, preskok)
    $ linija 2 (parter, preskok)

$ kada su jednaki bodovi sta raditi?
    $ Ko ima manji odbitak na E oceni ima prednost (treba da bude modularno, po promeni pravilnika)


## Kvalifikacije

$ Da li se pojedinačna, ekipna i finala po spravama odvijaju u posebnim danima ili se mogu nabiti u manje dana?
    $ tri finala 3 dana

$ Budući da u ekipnom finalu učestvuje 8 timova, da li to znači da se finale mora organizovati iz dva turnusa po 4 tima? Kako bi se onda određivalo u kom turnusu se koji tim takmiči? Da li se možda redosled određuje na osnovu plasmana ekipa u kvalifikacijama?
    $ u ekipi je najcesce 3 4 takmicara i onda sve moze da stane u jedan turnus (dve ekipe po spravi)

$ za finale po spravama moze da postoji tiebrake i da updane deveti ako dele mesto
$ svaka sprava ide pojedinacno, ne moze da se paralelizuje
    $ ako jedan takmici dva finala udaljiti ga sto vise



## Raspored

$ Ako se rotacija pravi da se deli na dva zagrevanja, da li ce onda tako biti na svakoj spravi ili se za svaku spravu zasebno definise hoce li biti iz delova ili odjednom
    $ na svim spravama mora da se ili ima jedna ili dve

$ kako za finala po spravama, svih osam ide odjednom?
    $ ne moze u paraleli?
    $ ako se potrefi da nijedan od takmicara ne takmici obe sprave u istom trenutku, da li da paralelizuje ta dva finala
        $ ne (gore je naveden odgovor)
    $ da li moze istovremeno vise kategorija odjednom da takmici?
    $ da li sve vazi isto kao i za kvalifikaciono takmicenenje sto se tice proglasenja, zagrevanja, opsteg zagrevanja
        $ da, sve isto
        $ moze da se navede da je jedno proglasenje posle svakog finala

$ U kom trenutku vise ne mogu da se menjaju informacije o takmicarima? Da li ostaviti da mogu uvek da se menjaju?
    $ Da li nakon tehnickog sastanka moze da se menja starosna kategorija ili sprave na kojima ce se takmicar takmiciti?
        $ nakon tehnickog se sve zakljucava
    $ do tada se vec napravi raspored, a na njemu su minimalne izmene
        $ poenta je da se napravi rapsored neskoliko dana ranije, a onda ce se ljudi zaliti pa ti izmenis

$ kako se menja raspored ako dodje do neke nezgode za vreme takmičenja?
    $ Može se omogućiti naknadno pomeranje rasporeda u slučaju nezgode ili nekog neočekivanog tehničkog problema, kao i ukoliko administrator takmičenja uoči da je ipak potrebno više vremena za jednu rotaciju. Pomeranje bi se vršilo navođenjem vremenskog intervala za koji da se pomeri čitav
        raspored ili definisanjem novog vremena potrebnog za rotaciju, gde bi se na osnovu njega iznova iskalkulisao raspored.
    $ nema vecih izmena, samo se navede da nije nastupio

$ ako se na priemr srusi neka sprava, pauzira se celo takmicenje da ne bi rotacije bile poremecene
$ ubaciti opciju nije nastupio



- koje sve role postoje?
    - glavni sudija
    - administrator(tehnicka pomoc)
    - D sudija
    - E sudija
    - supervizor
    - neautorizovani korisnik
    - da li da postoji korisnik poput kluba, da moze da vidi broj svojih prijavljenih takmicara
      - verovatno ce morati da se uloguje kao neki klub sa mejlom kluba 

