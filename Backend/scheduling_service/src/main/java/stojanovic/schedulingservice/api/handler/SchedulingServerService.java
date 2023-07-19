package stojanovic.schedulingservice.api.handler;


import stojanovic.schedulingservice.api.client.ApplicationClientService;
import stojanovic.schedulingservice.api.middleware.UserInfoInterceptor;
import application_pb.Application;
import auth_pb.Auth;
import io.grpc.Status;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.StreamObserver;
import lombok.RequiredArgsConstructor;
import net.devh.boot.grpc.server.service.GrpcService;
import scheduling_pb.Scheduling;
import scheduling_pb.SchedulingServiceGrpc;

@GrpcService
@RequiredArgsConstructor
public class SchedulingServerService extends SchedulingServiceGrpc.SchedulingServiceImplBase {

    private final ApplicationClientService applicationClientService;

    @Override
    public void test(Scheduling.TestMessage request, StreamObserver<Scheduling.TestResponse> responseObserver) {
        try{
            Application.SportsOrganisation sportsOrganisation = applicationClientService.getCurrentSportsOrg();
            Scheduling.TestResponse response = Scheduling.TestResponse.newBuilder()
                    .setResponse(sportsOrganisation.getName())
                    .build();

            responseObserver.onNext(response);
            responseObserver.onCompleted();
        }
        catch (StatusRuntimeException e){
            String errorMessage = e.getStatus().getDescription();
            Status status = Status.NOT_FOUND.withDescription(errorMessage);
            responseObserver.onError(status.asRuntimeException());
        }
    }
}
