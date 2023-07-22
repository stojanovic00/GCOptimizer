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

import java.time.*;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;
import java.util.concurrent.ExecutionException;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class ScheduleService {

    private final ApplicationClientService applicationClientService;
    private final SolverManager<Schedule, UUID> solverManager;
    public Scheduling.ScheduleDto generateSchedule(SchedulingParameters parameters) throws StatusRuntimeException {
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

        //TODO Finish processing
        // - Delete fully unused sessions
        // - Sort contestants on apparatuses
        //      - Group by same organization
        //      - Then city
        //      - Then country
        //      - Put non competing on end


        List<Long> startingTimes = generateStartingTimes(parameters);


        List<Scheduling.ScheduleSlot> slotsPb = ProtoMapper.scheduleSlotListPb(solution.getSlots());

        Scheduling.Schedule schedulePb = Scheduling.Schedule.newBuilder()
                .addAllSlots(slotsPb)
                .build();

        Scheduling.ScheduleDto dto = Scheduling.ScheduleDto.newBuilder()
               .setSchedule(schedulePb)
               .addAllStartingTimes(startingTimes)
               .build();

        return dto;
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
}
