package stojanovic.schedulingservice.api.exception;

public class NotFoundException extends Exception {
    public NotFoundException() {
        super("Entity not found");
    }
}
