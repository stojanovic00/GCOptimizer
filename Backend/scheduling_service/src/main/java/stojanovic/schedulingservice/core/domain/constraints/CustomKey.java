package stojanovic.schedulingservice.core.domain.constraints;

import lombok.AllArgsConstructor;
import lombok.Data;
import org.kie.api.definition.rule.All;
import scheduling_pb.Scheduling;
import stojanovic.schedulingservice.core.domain.model.ApparatusType;

@AllArgsConstructor
@Data
public class CustomKey {
    private Integer session;
    private ApparatusType apparatus;
    private int contestantsNum;
}
