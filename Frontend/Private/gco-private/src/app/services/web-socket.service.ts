import { EventEmitter, Injectable } from '@angular/core';
import { environment } from '../../environments/environment';
import { Apparatus } from '../model/core/apparatus';
import { WebSocketMessage } from '../model/web-socket/web-socket-message';

// @Injectable({
//   providedIn: 'root'
// })
export class WebSocketService {

    private socket: WebSocket;
    private listener: EventEmitter<any> = new EventEmitter();
    webSocketPath: string = environment.webSocketPath


    public constructor(apparatus: Apparatus, compId: string) {


        // Second parameter puts value in header "Sec-Websocket-Protocol"
        // I put there my bearer and I modified auth middleware to read from that header if it doesn't find "Authorization" header
        const requestPath: string = this.webSocketPath+"?apparatus=" + apparatus + "&competitionId=" + compId;
        const jwt: string = localStorage.getItem("jwt")!; 
        this.socket = new WebSocket(requestPath, [jwt]);
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

    public send(data: WebSocketMessage) {
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
