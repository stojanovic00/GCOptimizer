package stojanovic.schedulingservice.core.domain.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.util.List;
import java.util.UUID;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
public class SchedulingParameters {
    // All durations are measured in minutes!
    @Id
    private UUID id;
    @Column(nullable = false)
    private UUID competitionId;
    @Column(nullable = false)
    private java.time.LocalTime startTime;
    @Column(nullable = false)
    private java.time.LocalTime endTime;
    @Column(nullable = false)
    private boolean warmupRoomAvailable;
    @Column(nullable = false)
    private int generalWarmupTime;
    @Column(nullable = false)
    private int warmupTime;
    @Column(nullable = false)
    private int warmupsPerApparatus;
    @Column(nullable = false)
    private int contestantNumPerApparatus;
    @Column(nullable = false)
    private int executionTime;
    @Column(nullable = false)
    private int apparatusRotationTime;
    @Column(nullable = false)
    private int medalCeremonyAfterOneSessionTime;
    @Column(nullable = false)
    private int finalMedalCeremonyTime;
    @Column(nullable = false)
    private boolean halfApparatusPerSessionMode;
    @OneToMany(cascade = CascadeType.REMOVE)
    private List<Apparatus> apparatusOrder;
}
