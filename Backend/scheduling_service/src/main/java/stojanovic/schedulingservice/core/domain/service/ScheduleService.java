package stojanovic.schedulingservice.core.domain.service;

import application_pb.Application;
import io.grpc.StatusRuntimeException;
import lombok.RequiredArgsConstructor;
import org.optaplanner.core.api.solver.SolverJob;
import org.optaplanner.core.api.solver.SolverManager;
import org.springframework.stereotype.Service;
import scheduling_pb.Scheduling;
import stojanovic.schedulingservice.api.client.ApplicationClientService;
import stojanovic.schedulingservice.api.utils.ProtoMapper;
import stojanovic.schedulingservice.core.domain.model.*;
import stojanovic.schedulingservice.core.domain.repo.ScheduleRepo;

import java.time.*;
import java.util.*;
import java.util.concurrent.ExecutionException;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class ScheduleService {

    private final ApplicationClientService applicationClientService;
    private final SolverManager<Schedule, UUID> solverManager;
    private final ScheduleRepo scheduleRepo;
    public Schedule generateSchedule(SchedulingParameters parameters) throws StatusRuntimeException {
        // Delete existing schedule for this competition
       scheduleRepo.deleteByCompetitionId(parameters.getCompetitionId());

        //Get competition and applications from application service
        Application.ContestantApplicationList applications = applicationClientService.getCompetitionApplications(parameters.getCompetitionId());
        List<ApparatusType> apparatusOrder = parameters.getApparatusOrder().stream()
                .map(Apparatus::getType).collect(Collectors.toList());

        //Prepare data for optaplanner
        List<Contestant> contestants = generateContestants(applications);
        List<ScheduleSlot> slots = generateScheduleSlots(parameters, contestants.size(), apparatusOrder);

        //Initialize planning solution
        Schedule schedule = new Schedule(contestants, slots);


        UUID problemId = UUID.randomUUID();
        // Submit the problem to start solving
        SolverJob<Schedule, UUID> solverJob = solverManager.solve(problemId, schedule);
        Schedule solution;
        try {
            // Wait until the solving ends
            solution = solverJob.getFinalBestSolution();
        } catch (InterruptedException | ExecutionException e) {
            throw new IllegalStateException("Solving failed.", e);
        }
        // Post-processing
        // - Delete fully unused sessions
        // - Sort contestants on apparatuses
        //      - Group by same organization
        //      - Then city
        //      - Then country
        List<ApparatusType> appOrder = parameters.getApparatusOrder().stream().map(Apparatus::getType).collect(Collectors.toList());
        List<ScheduleSlot> processedSlots = processSlots(solution.getSlots(), appOrder);

        List<Long> startingTimes = generateStartingTimes(parameters);
        //Save
        solution.setSlots(processedSlots);
        solution.setStartingTimes(startingTimes);
        solution.setId(UUID.randomUUID());
        solution.setCompetitionId(parameters.getCompetitionId());

        scheduleRepo.save(solution);
        return solution;
    }

    private List<Long> generateStartingTimes(SchedulingParameters parameters) {
        List<Long> startingTimes = new ArrayList<Long>();
        int sessionDuration = calculateSessionDuration(parameters);
        for(int i = 0; i < calculateMaxSessionNum(parameters); i++){
            LocalTime startTime = parameters.getStartTime().plusMinutes(sessionDuration * i);
            LocalDateTime localDateTime = LocalDateTime.of(2001, 1, 1, startTime.getHour(), startTime.getMinute(), startTime.getSecond());
            ZonedDateTime zonedDateTime = localDateTime.atZone(ZoneOffset.UTC);

            startingTimes.add(zonedDateTime.toEpochSecond());
        }
        return startingTimes;
    }

    private List<Contestant> generateContestants(Application.ContestantApplicationList applications){
        List<Contestant> contestants = new ArrayList<Contestant>();
        int contestantCounter = 0;
        for(Application.ContestantApplication application : applications.getContestantApplicationsList()){
            contestantCounter++;
            contestants.add(parseContestant(application, contestantCounter));
        }
        return contestants;
    }

    private List<ScheduleSlot> generateScheduleSlots(SchedulingParameters parameters, int allContestantsNum, List<ApparatusType> apparatusOrder){
       double maxSessionNum = calculateMaxSessionNum(parameters);

       List<ScheduleSlot> slots = new ArrayList<ScheduleSlot>();
       long slotCounter = 0;

       //Each session has table for each apparatus with row number equal to contestants per apparatus
       for(int sessionNum = 1; sessionNum <= maxSessionNum;sessionNum++){
          for(Apparatus apparatus : parameters.getApparatusOrder()){
              for(int i = 0; i < parameters.getContestantNumPerApparatus(); i++){
                  slotCounter ++;
                  ScheduleSlot slot = new ScheduleSlot(slotCounter, sessionNum, apparatus.getType(), allContestantsNum, apparatusOrder);
                  slots.add(slot);
              }
          }
       }

        return slots;
    }

    private Contestant parseContestant(Application.ContestantApplication application,int compId){
        return   Contestant.builder()
                .id(UUID.fromString(application.getContestant().getDelegationMember().getId()))
                .contestantCompId(compId)
                .teamNumber(application.getTeamNumber())
                .name(application.getContestant().getDelegationMember().getFullName())
                .organization(application.getContestant().getDelegationMember().getSportsOrganisation().getName())
                .ageCategory(ProtoMapper.ageCategoryDom(application.getAgeCategory()))
                .Country(application.getContestant().getDelegationMember().getSportsOrganisation().getAddress().getCountry())
                .City(application.getContestant().getDelegationMember().getSportsOrganisation().getAddress().getCity())
                .competingApparatuses(ProtoMapper.apparatusApplicationListDom(application.getApparatusAnnouncementsList()))
                .build();
    }
    private int calculateSessionDuration(SchedulingParameters params){
        int generalWarmupTime;
        if(params.isWarmupRoomAvailable()){
            generalWarmupTime = 0;
        }
        else {
            generalWarmupTime = params.getGeneralWarmupTime();
        }

        //TODO 3-3 and 2-2 regime
        int numOfApparatusesInSession = params.getApparatusOrder().size();

        int apparatusWarmupTime = params.getWarmupTime()* params.getWarmupsPerApparatus()*numOfApparatusesInSession;
        int executionTime = params.getExecutionTime() * params.getContestantNumPerApparatus() * numOfApparatusesInSession;
        int rotationTime = params.getApparatusRotationTime() * (numOfApparatusesInSession - 1);


        return generalWarmupTime + apparatusWarmupTime + executionTime + rotationTime + params.getMedalCeremonyAfterOneSessionTime();
    }
    private int calculateMaxSessionNum(SchedulingParameters params) {
        long totalTime = Duration.between(params.getStartTime(), params.getEndTime()).toMinutes();

        long availableTime = totalTime - params.getFinalMedalCeremonyTime();
        int sessionTime = calculateSessionDuration(params);

        return (int) Math.floor((double) availableTime / sessionTime);
    }
     private List<ScheduleSlot> processSlots(List<ScheduleSlot> slots, List<ApparatusType> apparatusOrder){
         List<ScheduleSlot> processed = new ArrayList<ScheduleSlot>();
         List<List<ScheduleSlot>> groupedBySession = groupBySession(slots);

         List<List<ScheduleSlot>> notEmpty = new ArrayList<List<ScheduleSlot>>();

         //Remove completely empty sessions
         for(List<ScheduleSlot> session : groupedBySession){
             if(session.stream().anyMatch(slot -> slot.getContestant() != null)){
                notEmpty.add(session);
             }
         }

         for(List<ScheduleSlot> oneSessionList : notEmpty){
              List<List<ScheduleSlot>>  groupedByApparatus = groupByStartingApparatus(oneSessionList);
              //While grouping order messes up
              List<List<ScheduleSlot>>  groupedByApparatusOrdered = new ArrayList<List<ScheduleSlot>>();

              //Establish given apparatus order
              for(ApparatusType apparatus : apparatusOrder){
                    for(List<ScheduleSlot> unorderedMember : groupedByApparatus){
                        if(unorderedMember.get(0).getStartingApparatus() == apparatus){
                            groupedByApparatusOrdered.add(unorderedMember);
                            break;
                        }
                    }
              }

              //Now sorting inside every apparatus table
              for(List<ScheduleSlot> oneApparatusList : groupedByApparatusOrdered){

                  List<ScheduleSlot> sameOrganizationList = new ArrayList<ScheduleSlot>();
                  List<ScheduleSlot> sameCityList = new ArrayList<ScheduleSlot>();
                  List<ScheduleSlot> sameCountryList = new ArrayList<ScheduleSlot>();

                  // Organization groups go first
                  List<List<ScheduleSlot>> groupedByOrganization = groupByOrganization(oneApparatusList);
                  for(List<ScheduleSlot> oneOrgList : groupedByOrganization){
                      //Add only if there is more than one from same org
                      if(oneOrgList.size() > 1){
                          sameOrganizationList.addAll(oneOrgList);
                      }
                  }
                  //Remove them from further queries
                  oneApparatusList.removeAll(sameOrganizationList);

                  // Cities
                  List<List<ScheduleSlot>> groupedByCity = groupByCity(oneApparatusList);
                  for(List<ScheduleSlot> oneCityList : groupedByCity){
                      //Add only if there is more than one from same org
                      if(oneCityList.size() > 1){
                          sameCityList.addAll(oneCityList);
                      }
                  }
                  //Remove them from further queries
                  oneApparatusList.removeAll(sameCityList);

                  // Countries
                  List<List<ScheduleSlot>> groupedByCountry = groupByCountry(oneApparatusList);
                  for(List<ScheduleSlot> oneCountryList : groupedByCountry){
                      //Add only if there is more than one from same org
                      if(oneCountryList.size() > 1){
                          sameCountryList.addAll(oneCountryList);
                      }
                  }
                  //Remove them from further queries
                  oneApparatusList.removeAll(sameCountryList);

                  processed.addAll(sameOrganizationList);
                  processed.addAll(sameCityList);
                  processed.addAll(sameCountryList);
                  //Leftovers
                  processed.addAll(oneApparatusList);
              }
         }

         return processed;
     }


     // Grouping functions
    private List<List<ScheduleSlot>> groupBySession(List<ScheduleSlot> slots) {
        Map<Integer, List<ScheduleSlot>> groupedSlots = new HashMap<>();

        for (ScheduleSlot slot : slots) {
            int session = slot.getSession();
            groupedSlots.putIfAbsent(session, new ArrayList<>());
            groupedSlots.get(session).add(slot);
        }

        return new ArrayList<>(groupedSlots.values());
    }

    private  List<List<ScheduleSlot>> groupByStartingApparatus(List<ScheduleSlot> sessionSlots) {
        Map<String, List<ScheduleSlot>> groupedSlots = new HashMap<>();

        for (ScheduleSlot slot : sessionSlots) {
            String apparatusKey = ApparatusType.values()[slot.getStartingApparatus().ordinal()].toString();
            groupedSlots.putIfAbsent(apparatusKey, new ArrayList<>());
            groupedSlots.get(apparatusKey).add(slot);
        }

        return new ArrayList<>(groupedSlots.values());
    }
    private  List<List<ScheduleSlot>> groupByOrganization(List<ScheduleSlot> sessionSlots) {
        Map<String, List<ScheduleSlot>> groupedSlots = new HashMap<>();

        for (ScheduleSlot slot : sessionSlots) {
            if (slot.getContestant() != null && slot.getContestant().getOrganization() != null) {
                String organizationId = slot.getContestant().getOrganization();
                groupedSlots.putIfAbsent(organizationId, new ArrayList<>());
                groupedSlots.get(organizationId).add(slot);
            }
        }

        return new ArrayList<>(groupedSlots.values());
    }
    public List<List<ScheduleSlot>> groupByCity(List<ScheduleSlot> sessionSlots) {
        Map<String, List<ScheduleSlot>> groupedSlots = new HashMap<>();

        for (ScheduleSlot slot : sessionSlots) {
            if (slot.getContestant() != null && slot.getContestant().getCity() != null) {
                String city = slot.getContestant().getCity();
                groupedSlots.putIfAbsent(city, new ArrayList<>());
                groupedSlots.get(city).add(slot);
            }
        }

        return new ArrayList<>(groupedSlots.values());
    }

    public List<List<ScheduleSlot>> groupByCountry(List<ScheduleSlot> sessionSlots) {
        Map<String, List<ScheduleSlot>> groupedByCountry = new HashMap<>();

        for (ScheduleSlot slot : sessionSlots) {
            if (slot.getContestant() != null && slot.getContestant().getCountry() != null) {
                    String country = slot.getContestant().getCountry();
                    groupedByCountry.putIfAbsent(country, new ArrayList<>());
                    groupedByCountry.get(country).add(slot);
            }
        }

        return new ArrayList<>(groupedByCountry.values());
    }

    public Schedule getByCompetitionId(UUID id){
       return scheduleRepo.findFirstByCompetitionId(id);
    }
}
