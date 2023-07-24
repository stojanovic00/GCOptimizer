package stojanovic.schedulingservice.core.domain.repo;

import org.springframework.data.mongodb.repository.MongoRepository;
import org.springframework.data.mongodb.repository.Query;
import stojanovic.schedulingservice.core.domain.model.Schedule;

import java.util.UUID;

public interface ScheduleRepo extends MongoRepository<Schedule, UUID> {

    Schedule findFirstByCompetitionId(UUID competitionId);

    void deleteByCompetitionId(UUID competitionId);
}
