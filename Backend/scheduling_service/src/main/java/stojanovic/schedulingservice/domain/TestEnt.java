package stojanovic.schedulingservice.domain;

import lombok.*;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import java.util.UUID;

@Data
@AllArgsConstructor
@NoArgsConstructor
@Entity
public class TestEnt {
    @Id
    private UUID id;
    @Column(nullable = false)
    private String name;
}
