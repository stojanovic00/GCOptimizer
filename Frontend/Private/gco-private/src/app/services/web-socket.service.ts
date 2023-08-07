import { EventEmitter} from '@angular/core';
import { environment } from '../../environments/environment';
import { Apparatus } from '../model/core/apparatus';
import { WebSocketEventMessage } from '../model/web-socket/web-socket-event-message';

export class WebSocketService {

    private socket: WebSocket;
    private listener: EventEmitter<any> = new EventEmitter();
    webSocketPath: string = environment.webSocketPath


    public constructor(apparatus: Apparatus, compId: string) {


        // There is no authentication for web sockets right now, because it doesn't support auth headers
        // There should be made some challenge response authentication like first sending pre ws request that returns some code for response
        // and then sending that code when establishing ws
        const requestPath: string = this.webSocketPath+"?apparatus=" + apparatus + "&competitionId=" + compId;
        this.socket = new WebSocket(requestPath);
        
        this.socket.onopen = event => {
            this.listener.emit({"type": "open", "data": event});
        }
        this.socket.onclose = event => {
            this.listener.emit({"type": "close", "data": event});
        }
        this.socket.onmessage = event => {
            this.listener.emit({"type": "message", "data": JSON.parse(event.data)});
        }
    }

    public send(data: WebSocketEventMessage) {
        const jsonString :string = JSON.stringify(data)
        this.socket.send(jsonString);
    }

    public close() {
        this.socket.close();
    }

    public getEventListener() {
        return this.listener;
    }


}
