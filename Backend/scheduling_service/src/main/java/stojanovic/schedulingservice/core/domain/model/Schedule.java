package stojanovic.schedulingservice.core.domain.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.optaplanner.core.api.domain.solution.PlanningEntityCollectionProperty;
import org.optaplanner.core.api.domain.solution.PlanningScore;
import org.optaplanner.core.api.domain.solution.PlanningSolution;
import org.optaplanner.core.api.domain.solution.ProblemFactCollectionProperty;
import org.optaplanner.core.api.domain.valuerange.ValueRangeProvider;
import org.optaplanner.core.api.score.buildin.bendable.BendableScore;
import org.springframework.data.annotation.Id;
import org.springframework.data.annotation.Transient;
import org.springframework.data.mongodb.core.index.Indexed;
import org.springframework.data.mongodb.core.mapping.Document;

import java.util.List;
import java.util.UUID;

@PlanningSolution

@Data
@NoArgsConstructor
@AllArgsConstructor
@Document
public class Schedule {
    @Id
    UUID id;
    //Things that change inside schedule: slot (session, apparatus)
    // States that contestant is readonly.
    // ProblemFactCollectionProperties are available to constraint streams
    @ValueRangeProvider
    @ProblemFactCollectionProperty
    @Transient
    private List<Integer> sessions;
    @ValueRangeProvider

    @ProblemFactCollectionProperty
    @Transient
    private List<ApparatusType> startingApparatuses;
@PlanningEntityCollectionProperty private List<ScheduleSlot> slots;

    @PlanningScore(bendableHardLevelsSize = 2, bendableSoftLevelsSize = 3)
    @Transient
    // Solution is feasible if all hard score levels are at least 0
    private BendableScore score;

    //For persisting order when reading from in database
    private List<Long> startingTimes;
    private List<ApparatusType> apparatusOrder;
    @Indexed
    private UUID competitionId;

    public Schedule(List<Integer> sessions, List<ApparatusType> startingApparatuses, List<ScheduleSlot> slots) {
        this.sessions = sessions;
        this.startingApparatuses = startingApparatuses;
        this.slots = slots;
    }
}
