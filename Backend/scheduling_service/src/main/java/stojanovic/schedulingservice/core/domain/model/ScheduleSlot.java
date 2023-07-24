package stojanovic.schedulingservice.core.domain.model;

import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;
import org.optaplanner.core.api.domain.entity.PlanningEntity;
import org.optaplanner.core.api.domain.lookup.PlanningId;
import org.optaplanner.core.api.domain.variable.PlanningVariable;

import java.util.List;

@Data
@NoArgsConstructor
@AllArgsConstructor

@PlanningEntity
public class ScheduleSlot {
    @PlanningId
    long planningId;
    //SAVE
    int session;
    //SAVE
    ApparatusType startingApparatus;
    @PlanningVariable(nullable = true) // If not enough contestants some slots will be empty
    //SAVE
    Contestant contestant;

    //Helping fields
    int allContestantsNum;
    List<ApparatusType> apparatusOrder;

    public ScheduleSlot(long planningId, int session, ApparatusType apparatus, int allContestantsNum, List<ApparatusType> apparatusOrder){
        this.planningId = planningId;
        this.session = session;
        this.startingApparatus = apparatus;
        this.allContestantsNum = allContestantsNum;
        this.apparatusOrder = apparatusOrder;
    }

}
