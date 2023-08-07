package stojanovic.schedulingservice.core.domain.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.optaplanner.core.api.domain.entity.PlanningEntity;
import org.optaplanner.core.api.domain.lookup.PlanningId;
import org.optaplanner.core.api.domain.variable.PlanningVariable;

import java.util.List;
import java.util.Objects;

@Data
@NoArgsConstructor
@AllArgsConstructor

@PlanningEntity
public class ScheduleSlot {
    @PlanningId
    long planningId;
    Contestant contestant;

    //Helping fields
    int contestantsPerApparatus;
    List<ApparatusType> apparatusOrder;

    //Using Integer instead of int, because it is nullable
    @PlanningVariable(nullable = false)
    Integer session;

    @PlanningVariable(nullable = false)
    ApparatusType startingApparatus;



    public ScheduleSlot(long planningId, int contestantsPerApparatus, List<ApparatusType> apparatusOrder, Contestant contestant){
        this.planningId = planningId;
        this.contestantsPerApparatus = contestantsPerApparatus;
        this.apparatusOrder = apparatusOrder;
        this.contestant = contestant;
    }


    @Override
    public int hashCode() {
        return Objects.hash(planningId);
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (!(o instanceof ScheduleSlot)) return false;
        ScheduleSlot other = (ScheduleSlot) o;
        return planningId == other.planningId;
    }

}
