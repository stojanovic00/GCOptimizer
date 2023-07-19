package api.handler;


import api.middleware.UserInfoInterceptor;
import auth_pb.Auth;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;
import scheduling_pb.Scheduling;
import scheduling_pb.SchedulingServiceGrpc;

@GrpcService
public class SchedulingServerService extends SchedulingServiceGrpc.SchedulingServiceImplBase {
    @Override
    public void test(Scheduling.TestMessage request, StreamObserver<Scheduling.TestResponse> responseObserver) {
        Auth.UserInfo userInfo = UserInfoInterceptor.USER_INFO.get();

        String message = "email: " + userInfo.getEmail() + ", role: " + userInfo.getRole() + ", message: " + request.getMessage();

        Scheduling.TestResponse resp = Scheduling.TestResponse.newBuilder()
                .setResponse(message)
                .build();



        responseObserver.onNext(resp);
        responseObserver.onCompleted();
    }
}
