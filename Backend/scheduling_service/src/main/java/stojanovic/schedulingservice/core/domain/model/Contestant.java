package stojanovic.schedulingservice.core.domain.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import java.util.List;
import java.util.UUID;

@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Contestant {
    UUID id;
    //Assigned after retrieving info from application service
    int contestantCompId;
    int teamNumber;
    String name;
    String organization;
    AgeCategory ageCategory;
    String Country;
    String City;
    List<Apparatus> competingApparatuses;
}
