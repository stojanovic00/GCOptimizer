package stojanovic.schedulingservice.core.domain.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.optaplanner.core.api.domain.entity.PlanningEntity;
import org.optaplanner.core.api.domain.lookup.PlanningId;
import org.optaplanner.core.api.domain.valuerange.CountableValueRange;
import org.optaplanner.core.api.domain.valuerange.ValueRangeFactory;
import org.optaplanner.core.api.domain.variable.PlanningVariable;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.OneToOne;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor

@PlanningEntity
public class ScheduleSlot {
    @Id
    @PlanningId
    long id;
    int session;
    ApparatusType startingApparatus;
    @OneToOne
    @PlanningVariable(nullable = true) // If not enough contestants some slots will be empty
    Contestant contestant;


    public ScheduleSlot(long id, int session, ApparatusType apparatus){
        this.id = id;
        this.session = session;
        this.startingApparatus = apparatus;
    }
}
