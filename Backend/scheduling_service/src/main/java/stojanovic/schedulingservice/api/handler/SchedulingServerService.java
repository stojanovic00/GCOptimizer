package stojanovic.schedulingservice.api.handler;


import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import net.devh.boot.grpc.server.service.GrpcService;
import scheduling_pb.Scheduling;
import scheduling_pb.SchedulingServiceGrpc;
import stojanovic.schedulingservice.api.utils.ProtoMapper;
import stojanovic.schedulingservice.core.domain.model.Schedule;
import stojanovic.schedulingservice.core.domain.model.SchedulingParameters;
import stojanovic.schedulingservice.core.domain.service.ScheduleService;

import java.util.List;

@GrpcService
@RequiredArgsConstructor
public class SchedulingServerService extends SchedulingServiceGrpc.SchedulingServiceImplBase {

    private final ScheduleService scheduleService;

    @Override
    public void generateSchedule(Scheduling.SchedulingParameters request, StreamObserver<Scheduling.Schedule> responseObserver) {
        SchedulingParameters parameters = ProtoMapper.schedulingParametersDom(request);

        Schedule schedule;
        try{
            schedule = scheduleService.generateSchedule(parameters);

            List<Scheduling.ScheduleSlot> slots = ProtoMapper.scheduleSlotListPb(schedule.getSlots());

            Scheduling.Schedule response = Scheduling.Schedule.newBuilder()
                    .addAllSlots(slots)
                    .build();

            responseObserver.onNext(response);
            responseObserver.onCompleted();
    }
        catch (StatusRuntimeException e){
        String errorMessage=e.getStatus().getDescription();
        Status status= Status.NOT_FOUND.withDescription(errorMessage);
        responseObserver.onError(status.asRuntimeException());
    }

    }
}











