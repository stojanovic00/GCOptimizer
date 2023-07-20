package stojanovic.schedulingservice.core.domain.service;


import lombok.RequiredArgsConstructor;
import lombok.var;
import org.springframework.stereotype.Service;
import stojanovic.schedulingservice.api.exception.NotFoundException;
import stojanovic.schedulingservice.core.domain.repo.ApparatusRepo;
import stojanovic.schedulingservice.core.domain.repo.SchedulingParamsRepo;
import stojanovic.schedulingservice.core.domain.model.Apparatus;
import stojanovic.schedulingservice.core.domain.model.SchedulingParameters;

import java.util.Optional;
import java.util.UUID;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class SchedulingParamsService {

    private final SchedulingParamsRepo schParamRepo;
    private final ApparatusRepo apparatusRepo;

    public SchedulingParameters getByCompetitionId(UUID id) throws NotFoundException {
       Optional<SchedulingParameters> res =  schParamRepo.getByCompetitionId(id);
       if(!res.isPresent()){
          throw new NotFoundException();
       }
       return res.get();
    }

    //TODO repair
   public SchedulingParameters save(SchedulingParameters params){
       //Save new order
       var apparatusTypes = params.getApparatusOrder().stream()
               .map(Apparatus::getType).collect(Collectors.toList());

        if(params.getId() == null){
            params.setId(UUID.randomUUID());
        }
        else
        {
            //Delete old
            schParamRepo.deleteById(params.getId());
        }

       //Set new order
       params.setApparatusOrder(apparatusRepo.getApparatusesWithTypes(apparatusTypes));
       return schParamRepo.save(params);
   }

}
