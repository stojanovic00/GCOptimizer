package stojanovic.schedulingservice.api.utils;

import application_pb.Application;
import scheduling_pb.Scheduling;
import stojanovic.schedulingservice.core.domain.model.*;

import java.time.LocalDateTime;
import java.time.ZoneOffset;
import java.util.List;
import java.util.UUID;
import java.util.stream.Collectors;

public class ProtoMapper {

    public static LocalDateTime unixTimestampToLocalDateTime(long timestamp){
        return LocalDateTime.ofEpochSecond(timestamp, 0, ZoneOffset.UTC);
    }


    public static Apparatus apparatusDom(Scheduling.Apparatus apparatusPb){
        Apparatus apparatus = new Apparatus();
        apparatus.setType(ApparatusType.values()[apparatusPb.getType().getNumber()]);
        return apparatus;
    }

    public static List<Apparatus> apparatusListDom(List<Scheduling.Apparatus> apparatusesPb){
        return apparatusesPb.stream()
                .map(ProtoMapper::apparatusDom)
                .collect(Collectors.toList());
    }

    public static SchedulingParameters schedulingParametersDom(Scheduling.SchedulingParameters params){
        SchedulingParameters domParams = new SchedulingParameters();
        if(!params.getId().isEmpty()){
            domParams.setId(UUID.fromString(params.getId()));
        }
        domParams.setCompetitionId(UUID.fromString(params.getCompetitionId()));
        domParams.setStartTime(unixTimestampToLocalDateTime(params.getStartTime()).toLocalTime());
        domParams.setEndTime(unixTimestampToLocalDateTime(params.getEndTime()).toLocalTime());
        domParams.setWarmupRoomAvailable(params.getWarmupRoomAvailable());
        domParams.setGeneralWarmupTime(params.getGeneralWarmupTime());
        domParams.setWarmupTime(params.getWarmupTime());
        domParams.setWarmupsPerApparatus(params.getWarmupsPerApparatus());
        domParams.setContestantNumPerApparatus(params.getContestantNumPerApparatus());
        domParams.setExecutionTime(params.getExecutionTime());
        domParams.setApparatusRotationTime(params.getApparatusRotationTime());
        domParams.setMedalCeremonyAfterOneSessionTime(params.getMedalCeremonyAfterOneSessionTime());
        domParams.setFinalMedalCeremonyTime(params.getFinalMedalCeremonyTime());
        domParams.setHalfApparatusPerSessionMode(params.getHalfApparatusPerSessionMode());
        domParams.setApparatusOrder(apparatusListDom(params.getApparatusOrderList()));
        return domParams;
    }

    //From application service
    public static AgeCategory ageCategoryDom(Application.AgeCategory category){
            return AgeCategory.builder()
                    .id(UUID.fromString(category.getId()))
                    .name(category.getName())
                    .minAge(category.getMinAge())
                    .maxAge(category.getMaxAge())
                    .build();
    }
    public static Apparatus apparatusApplicationDom(Application.ApparatusAnnouncement apparatusAnnouncement){
        return Apparatus.builder()
                .type(ApparatusType.values()[apparatusAnnouncement.getApparatus().getNumber()])
                .build();
    }

    public static List<Apparatus> apparatusApplicationListDom(List<Application.ApparatusAnnouncement> apparatusAnnouncements){
        return apparatusAnnouncements.stream()
                .map(ProtoMapper::apparatusApplicationDom)
                .collect(Collectors.toList());
    }

   public static Scheduling.ScheduleSlot scheduleSlotPb(ScheduleSlot slot){
        Scheduling.ScheduleSlot.Builder  builder =  Scheduling.ScheduleSlot.newBuilder()
                .setSession(slot.getSession())
                .setStartingApparatus(Scheduling.ApparatusType.values()[slot.getStartingApparatus().ordinal()]);

        //Some slots will remain unassigned
        if(slot.getContestant() != null){
            builder.setContestantInfo(contestantToContestantInfo(slot.getContestant()));
        }

        return builder.build();
    }

    public static List<Scheduling.ScheduleSlot> scheduleSlotListPb(List<ScheduleSlot> slots){
       return slots.stream()
               .map(ProtoMapper::scheduleSlotPb)
               .collect(Collectors.toList());
    }

    public static Scheduling.ContestantInfo contestantToContestantInfo(Contestant contestant){
        if(contestant == null)
            return null;
       return Scheduling.ContestantInfo.newBuilder()
               .setContestantCompId(contestant.getContestantCompId())
               .setName(contestant.getName())
               .setTeamNumber(contestant.getTeamNumber())
               .setOrganization(contestant.getOrganization())
               .setAgeCategory(contestant.getAgeCategory().getName())
               .setLocation(contestant.getCountry() + ", " + contestant.getCity())
               .build();
    }


}