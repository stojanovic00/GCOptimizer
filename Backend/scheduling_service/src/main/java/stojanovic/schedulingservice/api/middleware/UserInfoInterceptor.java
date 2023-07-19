package stojanovic.schedulingservice.api.middleware;

import auth_pb.Auth;
import io.grpc.*;
import lombok.var;
import net.devh.boot.grpc.server.interceptor.GrpcGlobalServerInterceptor;

@GrpcGlobalServerInterceptor
public class UserInfoInterceptor implements ServerInterceptor {
    //Each request has its own context so there are no concurrency problems
    public static final Context.Key<Auth.UserInfo> USER_INFO = Context.key("user-info");

    public <ReqT, RespT> ServerCall.Listener<ReqT> interceptCall(ServerCall<ReqT, RespT> call,
                                                                 Metadata headers,
                                                                 ServerCallHandler<ReqT, RespT> next) {

        String userInfoStr = headers.get(Metadata.Key.of("user-info", Metadata.ASCII_STRING_MARSHALLER));
        // Reject the request if userInfo is empty
        if (userInfoStr == null || userInfoStr.isEmpty()) {
            call.close(Status.INVALID_ARGUMENT.withDescription("Missing user-info"), headers);
            return new ServerCall.Listener() {
            };
        }

        //Parsing
        var userInfoParts = userInfoStr.split("\"");
        Auth.UserInfo userInfo = Auth.UserInfo.newBuilder()
                .setEmail(userInfoParts[1])
                .setRole(userInfoParts[3])
                .build();

        Context context = Context.current().withValue(USER_INFO, userInfo);
        return Contexts.interceptCall(context, call, headers, next);
    }
}