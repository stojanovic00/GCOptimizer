package stojanovic.schedulingservice;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;

@SpringBootApplication
@ComponentScan(basePackages = {"api.handler", "api.middleware"})
public class SchedulingServiceApplication {
	private static String grpcServerPort;


	public SchedulingServiceApplication(@Value("${grpc.server.port}") String grpcServerPort) {
		this.grpcServerPort = grpcServerPort;
	}

	public static void main(String[] args) {
		SpringApplication.run(SchedulingServiceApplication.class, args);
		System.out.println("Scheduling microservice started");
		System.out.println("Listening on port: " + grpcServerPort);
	}
}
