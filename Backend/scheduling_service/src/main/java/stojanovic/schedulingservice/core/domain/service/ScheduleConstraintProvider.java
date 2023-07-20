package stojanovic.schedulingservice.core.domain.service;

import org.optaplanner.core.api.score.buildin.hardsoft.HardSoftScore;
import org.optaplanner.core.api.score.stream.Constraint;
import org.optaplanner.core.api.score.stream.ConstraintFactory;
import org.optaplanner.core.api.score.stream.ConstraintProvider;
import org.optaplanner.core.api.score.stream.Joiners;
import stojanovic.schedulingservice.core.domain.model.ScheduleSlot;

public class ScheduleConstraintProvider implements ConstraintProvider {
    @Override
    public Constraint[] defineConstraints(ConstraintFactory constraintFactory) {
        return new Constraint[]{

        };
    }

    private Constraint sessionConflict(ConstraintFactory constraintFactory) {
        // Slot with contestant of lesser age category cant be in session after greater age category
        //So it should be I cat, II cat, III cat

        // Select a slot ...
        return constraintFactory
                .forEach(ScheduleSlot.class)
                // ... and pair it with another slot ...
                .join(ScheduleSlot.class,
                        // ... greater session ...
                        Joiners.greaterThan(ScheduleSlot::getSession, ScheduleSlot::getSession),
                        // ... in the same room ...
                        Joiners.equal(ScheduleSlot::getContestant),
                        // ... and the pair is unique (different id, no reverse pairs) ...
                        Joiners.lessThan(ScheduleSlot::getId))
                // ... then penalize each pair with a hard weight.
                .penalize(HardSoftScore.ONE_HARD)
                .asConstraint("Session conflict");
    }


}
