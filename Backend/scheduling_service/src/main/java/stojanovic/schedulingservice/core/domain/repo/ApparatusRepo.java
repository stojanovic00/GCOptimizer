package stojanovic.schedulingservice.core.domain.repo;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import stojanovic.schedulingservice.core.domain.model.Apparatus;
import stojanovic.schedulingservice.core.domain.model.ApparatusType;

import java.util.List;
import java.util.UUID;

public interface ApparatusRepo extends JpaRepository<Apparatus, UUID> {
    @Query("select  app " +
            "from Apparatus app " +
            " where app.type in :typeList")
    List<Apparatus> getApparatusesWithTypes(List<ApparatusType> typeList);
}
