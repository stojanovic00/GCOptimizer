- http i https (dva servera)
  - da li koristiti traefik ili samo api gateway kao reverse proxy

- dependency injection
    - da li za svaki handler da se inicijalizuje posebni servis
    - kako to funkcionise uopste,  citav put od mux-a pa do baze
        - da li ako guramo svuda istu instancu servisa svi hendleri koriste isti servis i tu moze da dodje do bottlenceka
- da li da se sudije autenitifikuju na sistemu uz pomoc one time passworda
