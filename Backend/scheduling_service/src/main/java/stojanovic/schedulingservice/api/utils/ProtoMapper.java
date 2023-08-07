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

    public static List<Scheduling.ApparatusType> apparatusTypeListPb(List<Apparatus> types){
       return types.stream().map(apparatus -> Scheduling.ApparatusType.forNumber(apparatus.getType().ordinal())).collect(Collectors.toList());
    }

    public static Scheduling.ContestantInfo contestantToContestantInfo(Contestant contestant){
        if(contestant == null)
            return null;
       return Scheduling.ContestantInfo.newBuilder()
               .setId(contestant.getId().toString())
               .setContestantCompId(contestant.getContestantCompId())
               .setName(contestant.getName())
               .setTeamNumber(contestant.getTeamNumber())
               .setOrganization(contestant.getOrganization())
               .setAgeCategory(contestant.getAgeCategory().getName())
               .setLocation(contestant.getCountry() + ", " + contestant.getCity())
               .addAllCompetingApparatuses(apparatusTypeListPb(contestant.getCompetingApparatuses()))
               .build();
    }

    public  static Scheduling.ApparatusType apparatusTypeDomToPb(ApparatusType apparatusType){
       return Scheduling.ApparatusType.forNumber(apparatusType.ordinal());
    }
    public  static List<Scheduling.ApparatusType> apparatusTypeListDomToPb(List<ApparatusType> apparatusTypes){
        return apparatusTypes.stream().map(ProtoMapper::apparatusTypeDomToPb).collect(Collectors.toList());
    }
    public static Scheduling.Schedule schedulePb(Schedule schedule) {
        List<Scheduling.ScheduleSlot> slotsPb = ProtoMapper.scheduleSlotListPb(schedule.getSlots());
        return Scheduling.Schedule.newBuilder()
                .setId(schedule.getId().toString())
                .addAllSlots(slotsPb)
                .addAllStartingTimes(schedule.getStartingTimes())
                .addAllApparatusOrder(apparatusTypeListDomToPb(schedule.getApparatusOrder()))
                .build();
    }

}
