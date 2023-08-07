package stojanovic.schedulingservice.core.domain.constraints;

import org.optaplanner.core.api.score.buildin.bendable.BendableScore;
import org.optaplanner.core.api.score.stream.Constraint;
import org.optaplanner.core.api.score.stream.ConstraintFactory;
import org.optaplanner.core.api.score.stream.ConstraintProvider;
import stojanovic.schedulingservice.core.domain.model.Apparatus;
import stojanovic.schedulingservice.core.domain.model.ApparatusType;
import stojanovic.schedulingservice.core.domain.model.ScheduleSlot;


import java.util.List;
import java.util.stream.Collectors;

import static org.optaplanner.core.api.score.stream.ConstraintCollectors.*;
import static org.optaplanner.core.api.score.stream.Joiners.equal;
import static org.optaplanner.core.api.score.stream.Joiners.lessThan;

public class ScheduleConstraintProvider implements ConstraintProvider {
    private static final int BENDABLE_SCORE_HARD_LEVELS_SIZE = 2;
    private static final int BENDABLE_SCORE_SOFT_LEVELS_SIZE = 3;
    @Override
    public Constraint[] defineConstraints(ConstraintFactory constraintFactory) {
        return new Constraint[ ]{
                apparatusNumGreaterThanGiven(constraintFactory),
                sameSportOrgContestantsWithSameSessionOnDifferentApparatuses(constraintFactory),
                contestantWithLesserAgeCategoryInGreaterSession(constraintFactory),
                greaterSessionHasMoreFilledSlots(constraintFactory),
                contestantsWithSameAgeCategoryInDiffSession(constraintFactory),
                contestantsWithDiffAgeCategoryInSameSession(constraintFactory),
                contestantsWithSameCountryInSameSessionOnSameApparatus(constraintFactory),
                contestantsWithSameCityInSameSessionOnSameApparatus(constraintFactory),
                contestantPenalizeWait(constraintFactory)
        };
    }

//******************************
//    HARD 1
//******************************

// Contestants per apparatus must be lower or equal to give number of contestants per apparatus
private Constraint apparatusNumGreaterThanGiven(ConstraintFactory factory) {
    // Group by the number of contestants per apparatus
    return factory.forEach(ScheduleSlot.class)
            .groupBy( slot -> new CustomKey(slot.getSession(), slot.getStartingApparatus(), slot.getContestantsPerApparatus()), count())
            .filter((key, count) ->{
                return count > key.getContestantsNum();
            })
            .penalize(BendableScore.ofHard(
                    BENDABLE_SCORE_HARD_LEVELS_SIZE,
                    BENDABLE_SCORE_SOFT_LEVELS_SIZE,
                    0,
                    1
            ))
            .asConstraint("Apparatus number greater than given");
}

//******************************
//    HARD 2
//******************************

//    Contestants from same organisation must have same starting apparatus in one session

    private Constraint sameSportOrgContestantsWithSameSessionOnDifferentApparatuses(ConstraintFactory factory) {
        return factory.forEach(ScheduleSlot.class)
                .join(ScheduleSlot.class,
                        equal(ScheduleSlot::getSession, ScheduleSlot::getSession),
                        equal(slot -> {
                            return slot.getContestant().getOrganization();
                        }, slot ->{
                            return slot.getContestant().getOrganization();
                        })
                )
                .filter(
                        (slot1, slot2) -> slot1.getStartingApparatus() != slot2.getStartingApparatus()
                )
                .penalize(BendableScore.ofHard(
                        BENDABLE_SCORE_HARD_LEVELS_SIZE,
                        BENDABLE_SCORE_SOFT_LEVELS_SIZE,
                        1,
                        1
                ))
                .asConstraint("Same sports organization contestants in same session on different apparatuses");
    }


//******************************
//    SOFT 1
//******************************
//    Contestants with same age category should be in same session
    private Constraint contestantsWithSameAgeCategoryInDiffSession(ConstraintFactory factory) {
        return factory.forEach(ScheduleSlot.class)
                .join(ScheduleSlot.class,
                       equal(
                               slot -> slot.getContestant().getAgeCategory(),
                               slot -> slot.getContestant().getAgeCategory()
                       ),
                        lessThan(ScheduleSlot::getPlanningId, ScheduleSlot::getPlanningId)
                    )
                .filter( (slot1, slot2) -> {
                    return slot1.getSession() != slot2.getSession();
                })
                .penalize(BendableScore.ofSoft(
                        BENDABLE_SCORE_HARD_LEVELS_SIZE,
                        BENDABLE_SCORE_SOFT_LEVELS_SIZE,
                        0,
                        1
                ))
                .asConstraint("Contestants with same age category in different session");
    }

    private Constraint contestantsWithDiffAgeCategoryInSameSession(ConstraintFactory factory) {
        return factory.forEach(ScheduleSlot.class)
                .join(ScheduleSlot.class,
                       equal(ScheduleSlot::getSession, ScheduleSlot::getSession),
                        lessThan(ScheduleSlot::getPlanningId, ScheduleSlot::getPlanningId)
                )
                .filter((slot1, slot2) -> {
                    return  !slot1.getContestant().getAgeCategory().getName().equals(slot2.getContestant().getAgeCategory().getName());
                })
                .penalize(BendableScore.ofSoft(
                        BENDABLE_SCORE_HARD_LEVELS_SIZE,
                        BENDABLE_SCORE_SOFT_LEVELS_SIZE,
                        0,
                        1
                ))
                .asConstraint("Contestants different  age category in same session");
    }
//    Contestants with lesser age category shouldn't be in greater session
//    In layman's terms : schedule should be sorted by age category ascending

    private Constraint contestantWithLesserAgeCategoryInGreaterSession(ConstraintFactory factory) {
        return factory.forEach(ScheduleSlot.class)
                .join(ScheduleSlot.class)
                .filter((slot1, slot2) -> {
                        return
                                slot1.getContestant().getAgeCategory().getMinAge() < slot2.getContestant().getAgeCategory().getMinAge()
                                    &&
                                slot1.getSession() > slot2.getSession();
                    }
                )
                .penalize(BendableScore.ofSoft(
                        BENDABLE_SCORE_HARD_LEVELS_SIZE,
                        BENDABLE_SCORE_SOFT_LEVELS_SIZE,
                        0,
                        1
                ))
                .asConstraint("Contestants with lesser age category in greater session");
    }



//******************************
//    SOFT 2
//******************************


    private Constraint greaterSessionHasMoreFilledSlots(ConstraintFactory factory) {
        return factory.forEach(ScheduleSlot.class)
                .groupBy(ScheduleSlot::getSession, count())
                .penalize(BendableScore.ofSoft(
                        BENDABLE_SCORE_HARD_LEVELS_SIZE,
                        BENDABLE_SCORE_SOFT_LEVELS_SIZE,
                        0,
                        2
                        //This is penalty function (it was sufficient to state it for previous constraints(default value is 1)).
                        //Final penalty is weight(in this case 1) * penalty function result
                        //This function penalizes more, greater sessions with higher count
                ), (session, count) -> session* count )
                .asConstraint("Greater session has more filled slots than lesser session");
    }

//    Contestants should have same starting apparatus if:
//          2: They are from same city

    private Constraint contestantsWithSameCityInSameSessionOnSameApparatus(ConstraintFactory factory) {
        return factory.forEach(ScheduleSlot.class)
                .join(ScheduleSlot.class,
                        equal(ScheduleSlot::getStartingApparatus, ScheduleSlot::getStartingApparatus)
                )
                .filter(
                        (slot1, slot2) -> slot1.getContestant().getCity().equals(slot2.getContestant().getCity())
                )
                .reward(BendableScore.ofSoft(
                        BENDABLE_SCORE_HARD_LEVELS_SIZE,
                        BENDABLE_SCORE_SOFT_LEVELS_SIZE,
                        1,
                        2
                ))
                .asConstraint("Contestants from same city in same session");
    }

//    Contestants should have same starting apparatus if:
//      3:they are from same country

    private Constraint contestantsWithSameCountryInSameSessionOnSameApparatus(ConstraintFactory factory) {
        return factory.forEach(ScheduleSlot.class)
                .join(ScheduleSlot.class,
                        equal(ScheduleSlot::getStartingApparatus, ScheduleSlot::getStartingApparatus)
                )
                .filter(
                        (slot1, slot2) -> slot1.getContestant().getCountry().equals(slot2.getContestant().getCountry())
                )
                .reward(BendableScore.ofSoft(
                        BENDABLE_SCORE_HARD_LEVELS_SIZE,
                        BENDABLE_SCORE_SOFT_LEVELS_SIZE,
                        1,
                        1
                ))
                .asConstraint("Contestants from same country in same session");
    }

//******************************
//    SOFT 3
//******************************

//    For contestants that don't compete on all apparatuses:
//        Give them starting apparatus, so they finish as fast as possible:
//            Minimal waiting between apparatuses
//
//            If competition apparatus order is 0 1 2 3 4 5
//            Best solution for contestant that competes only apparatuses 0 3 5 is to start on apparatus 3

    private Constraint contestantPenalizeWait(ConstraintFactory factory) {
        return factory.forEach(ScheduleSlot.class)
                .penalize(BendableScore.ofSoft(
                        BENDABLE_SCORE_HARD_LEVELS_SIZE,
                        BENDABLE_SCORE_SOFT_LEVELS_SIZE,
                        2,
                        1
                ),
                        slot ->{
                            List<ApparatusType> apparatusOrder = slot.getApparatusOrder();
                            int numOfApparatuses = apparatusOrder.size();
                            int startingAppIndex = apparatusOrder.indexOf(slot.getStartingApparatus());

                            int waiting = 0;

                            List<ApparatusType> contestantsApparatuses = slot.getContestant().getCompetingApparatuses().stream()
                                    .map(Apparatus::getType).
                                    collect(Collectors.toList());
                            // Goes on each apparatus once
                            //Handling overflow with %
                            for(int i = 0; i < numOfApparatuses; i++){
                                int index = (startingAppIndex + i) % numOfApparatuses;
                                if(!contestantsApparatuses.remove(apparatusOrder.get(index))){
                                    //If it doesn't find and remove apparatus from contestants apparatuses
                                    //That means he waits
                                    waiting++;
                                }

                                //Finished all his apparatuses
                                if(contestantsApparatuses.isEmpty()){
                                    break;
                                }
                            }
                            return waiting;
                        }
                        )
                .asConstraint("Contestants non competing apparatus wait");
    }
}
