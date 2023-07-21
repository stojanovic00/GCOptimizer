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

import java.util.List;

@PlanningSolution

@Data
@NoArgsConstructor
@AllArgsConstructor
public class Schedule {
    @ValueRangeProvider
    // States that contestant is readonly.
    // Only thing that changes in this optimization problem is contestant field inside ScheduleSlot
    // ProblemFactCollectionProperties are available to constraint streams
    @ProblemFactCollectionProperty
    private List<Contestant> contestants;

    @PlanningEntityCollectionProperty
    private List<ScheduleSlot> slots;

    @PlanningScore(bendableHardLevelsSize = 2, bendableSoftLevelsSize = 3)
    // Solution is feasible if all hard score levels are at least 0
    private BendableScore score;



   public Schedule(List<Contestant> contestants, List<ScheduleSlot> slots){
       this.contestants = contestants;
       this.slots = slots;
   }

}
