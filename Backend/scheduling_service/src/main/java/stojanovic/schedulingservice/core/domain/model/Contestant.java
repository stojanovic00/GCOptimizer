package stojanovic.schedulingservice.core.domain.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.*;
import java.util.List;
import java.util.UUID;

@Entity
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class Contestant {
    @Id
    UUID id;
    //Assigned after retrieving info from application service
    int contestantCompId;
    int teamNumber;
    String organization;
    @Embedded
    @AttributeOverrides({
        @AttributeOverride( name = "name", column = @Column(name = "age_cat_name")),
        @AttributeOverride( name = "minAge", column = @Column(name = "age_cat_min_age")),
        @AttributeOverride( name = "maxAge", column = @Column(name = "age_cat_max_age"))
    })
    AgeCategory ageCategory;
    String Country;
    String City;
    @OneToMany
    List<Apparatus> competingApparatuses;
}
