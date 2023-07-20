package stojanovic.schedulingservice.api.client;

import stojanovic.schedulingservice.api.middleware.UserInfoInterceptor;
import application_pb.Application;
import application_pb.ApplicationServiceGrpc;
import auth_pb.Auth;
import io.grpc.Channel;
import io.grpc.ClientInterceptors;
import io.grpc.Metadata;
import io.grpc.StatusRuntimeException;
import io.grpc.stub.MetadataUtils;
import net.devh.boot.grpc.client.inject.GrpcClient;
import org.springframework.stereotype.Service;

import java.util.UUID;


@Service
public class ApplicationClientService {
    @GrpcClient("grpc-application-service")
    ApplicationServiceGrpc.ApplicationServiceBlockingStub client; //synchronous client

    public Application.ContestantApplicationList getCompetitionApplications(UUID competitionId) throws StatusRuntimeException{
        Application.IdMessage idMessage =  Application.IdMessage.newBuilder()
                .setId(competitionId.toString())
                .build();

        //TODO: used for checking if this sports organization can get applications for this competition
        ApplicationServiceGrpc.ApplicationServiceBlockingStub modifiedClient = getClientWithAttachedUserInfoMetadata();

        return modifiedClient.getAllContestantApplications(idMessage);
    }

    private ApplicationServiceGrpc.ApplicationServiceBlockingStub getClientWithAttachedUserInfoMetadata() {
        // TODO null exception handling
        Auth.UserInfo userInfo = UserInfoInterceptor.USER_INFO.get();

        // Create metadata with the user information
        Metadata metadata = new Metadata();
        String userInfoStr = userInfo.toString().replaceAll("\\n", "");
        metadata.put(Metadata.Key.of("user-info", Metadata.ASCII_STRING_MARSHALLER), userInfoStr);

        // Attach the metadata to the gRPC call
        ClientInterceptors.intercept(client.getChannel(), MetadataUtils.newAttachHeadersInterceptor(metadata));

        // Create a new Channel with the modified metadata
        Channel modifiedChannel = ClientInterceptors.intercept(client.getChannel(), MetadataUtils.newAttachHeadersInterceptor(metadata));

        // Create a new stub with the modified Channel
        return ApplicationServiceGrpc.newBlockingStub(modifiedChannel);
    }
}
