package api.handler;


import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import scheduling_pb.Scheduling;
import scheduling_pb.SchedulingServiceGrpc;

@GrpcService
public class SchedulingServerService extends SchedulingServiceGrpc.SchedulingServiceImplBase {
    @Override
    public void test(Scheduling.TestMessage request, StreamObserver<Scheduling.TestResponse> responseObserver) {
        String uppercase = request.getMessage().toUpperCase();
        Scheduling.TestResponse resp = Scheduling.TestResponse.newBuilder()
                .setResponse(uppercase)
                .build();

        responseObserver.onNext(resp);
        responseObserver.onCompleted();
    }
}
