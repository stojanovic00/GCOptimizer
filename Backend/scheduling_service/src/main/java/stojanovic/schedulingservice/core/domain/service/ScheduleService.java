package stojanovic.schedulingservice.core.domain.service;

import application_pb.Application;
import io.grpc.StatusRuntimeException;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import stojanovic.schedulingservice.api.client.ApplicationClientService;
import stojanovic.schedulingservice.api.utils.ProtoMapper;
import stojanovic.schedulingservice.core.domain.model.*;

import java.time.Duration;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

@Service
@RequiredArgsConstructor
public class ScheduleService {

    private final ApplicationClientService applicationClientService;
    public void generateSchedule(SchedulingParameters parameters) throws StatusRuntimeException {
        //Get applications from application service
        Application.ContestantApplicationList applications = applicationClientService.getCompetitionApplications(parameters.getCompetitionId());

        //Prepare data for optaplanner
        List<Contestant> contestants = generateContestants(applications);
        List<ScheduleSlot> slots = generateScheduleSlots(parameters);

        Schedule schedule = new Schedule(contestants, slots);

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

    private List<ScheduleSlot> generateScheduleSlots(SchedulingParameters parameters){
       double maxSessionNum = calculateMaxSessionNum(parameters);

       List<ScheduleSlot> slots = new ArrayList<ScheduleSlot>();
       long slotCounter = 0;

       //Each session has table for each apparatus with row number equal to contestants per apparatus
       for(int sessionNum = 1; sessionNum <= maxSessionNum;sessionNum++){
          for(Apparatus apparatus : parameters.getApparatusOrder()){
              for(int i = 0; i < parameters.getContestantNumPerApparatus(); i++){
                  slotCounter ++;
                  ScheduleSlot slot = new ScheduleSlot(slotCounter, sessionNum, apparatus.getType());
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
                .organization(application.getContestant().getDelegationMember().getSportsOrganisation().getName())
                .ageCategory(ProtoMapper.ageCategoryDom(application.getAgeCategory()))
                .Country(application.getContestant().getDelegationMember().getSportsOrganisation().getAddress().getCountry())
                .City(application.getContestant().getDelegationMember().getSportsOrganisation().getAddress().getCity())
                .competingApparatuses(ProtoMapper.apparatusApplicationListDom(application.getApparatusAnnouncementsList()))
                .build();
    }

    private double calculateMaxSessionNum(SchedulingParameters params) {
        long totalTime = Duration.between(params.getStartTime(), params.getEndTime()).toMinutes();

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


        int sessionTime = generalWarmupTime + apparatusWarmupTime + executionTime + rotationTime + params.getMedalCeremonyAfterOneSessionTime();
        long availableTime = totalTime - params.getFinalMedalCeremonyTime();

        return Math.floor((double) availableTime / sessionTime);
    }
}
