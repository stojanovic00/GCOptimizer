package stojanovic.schedulingservice.core.domain.constraints;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class SessionCategoryKey {
    private Integer session;
    private String category;
}
