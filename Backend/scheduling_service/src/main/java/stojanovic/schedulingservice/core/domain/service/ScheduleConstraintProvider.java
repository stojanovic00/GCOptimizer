package stojanovic.schedulingservice.core.domain.service;

import org.optaplanner.core.api.score.stream.Constraint;
import org.optaplanner.core.api.score.stream.ConstraintFactory;
import org.optaplanner.core.api.score.stream.ConstraintProvider;

public class ScheduleConstraintProvider implements ConstraintProvider {
    @Override
    public Constraint[] defineConstraints(ConstraintFactory constraintFactory) {
        return new Constraint[]{

        };
    }

//    Constraint

//    HARD 1:
//    Contestant can be assigned to only one slot
//    There can't be contestant that isn't assigned to any slot

//    HARD 2
//    Contestants from same organisation must have same starting apparatus in one session

//    SOFT 1 (Maybe even HARD 2)
//    Contestants with lower age category should be in lower session

//    SOFT 2
//    All contestants with same age category should be in same session
//    There should be minimal number of sessions

//    SOFT 3
//    Contestants from different age categories that happen to be in same session
//        Should have same starting apparatus if:
//          1: They are from same sports organization! (Defined as HARD 2)
//          2: They are from same country
//          3: They are from same city
//    SOFT 4
//    For contestants that don't compete on all apparatuses:
//        Give them starting apparatus, so they finish as fast as possible:
//            Minimal waiting between apparatuses (1)
//            Finish with all competing apparatuses as soon as possible
//
//            If competition apparatus order is 0 1 2 3 4 5
//            Best solution for contestant that competes only apparatuses 0 3 5 is to start on apparatus 3
}
