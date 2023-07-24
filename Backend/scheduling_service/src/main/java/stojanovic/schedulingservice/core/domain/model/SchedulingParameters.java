package stojanovic.schedulingservice.core.domain.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;
import java.util.UUID;

@Data
@NoArgsConstructor
@AllArgsConstructor
public class SchedulingParameters {
    // All durations are measured in minutes!
    private UUID scheduleId; //Used for updating existing schedule
    private UUID competitionId;
    private java.time.LocalTime startTime;
    private java.time.LocalTime endTime;
    private boolean warmupRoomAvailable;
    private int generalWarmupTime;
    private int warmupTime;
    private int warmupsPerApparatus;
    private int contestantNumPerApparatus;
    private int executionTime;
    private int apparatusRotationTime;
    private int medalCeremonyAfterOneSessionTime;
    private int finalMedalCeremonyTime;
    private boolean halfApparatusPerSessionMode;
    private List<Apparatus> apparatusOrder;
}
