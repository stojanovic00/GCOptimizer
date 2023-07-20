package stojanovic.schedulingservice.core.domain.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

import javax.persistence.Embeddable;
import javax.persistence.Entity;
import javax.persistence.Id;
import java.util.UUID;

@Entity
@Embeddable
@Data
@NoArgsConstructor
@AllArgsConstructor
@Builder
public class AgeCategory {
    @Id
    UUID id;
    String name;
    int minAge;
    int maxAge;
}
