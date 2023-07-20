package stojanovic.schedulingservice.core.domain.repo;

import org.springframework.data.jpa.repository.JpaRepository;
import stojanovic.schedulingservice.core.domain.model.SchedulingParameters;

import java.util.Optional;
import java.util.UUID;

public interface SchedulingParamsRepo extends JpaRepository<SchedulingParameters, UUID> {
    public Optional<SchedulingParameters> getByCompetitionId(UUID id);
}
