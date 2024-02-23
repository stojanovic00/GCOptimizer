# Gymnastics Competitions Optimizer

## About

- Microservice based web platform that supports whole process of organizing, scheduling and executing gymnastics competitions.

- Application microservice handles initial specification and organization of gymnastics competition and also applications of contestants and judges for desired competitions.

- Scheduling microservice uses OptaPlanner to tackle NP hard problem of creating contestants per apparatuses schedule for each competition session taking into account various constraints and optimizations to enhance the quality of each contestants experience and also  minimize overall competitions duration.

- Scoring service is based on custom websocket server that serves purpose of providing real time user interface for judges and competition administrators to score, calculate final scores and at the end create rang lists for individual contestants and teams.

- In depth analysis and explanation of development of this project can be found in `aleksandar-stojanovic-diplomski.pdf`