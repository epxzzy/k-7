package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)
var SendQueueC = Queue{Size: 10};
var RecieveQueueC = Queue{Size: 10};

func initWsClient(){
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:8080", Path: "/tunnel"}
	log.Printf("connecitng to ws:localhost:8080/k7endpoint")	

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dialer died dialing: ", err)
	}

	go WsClientRLoop(conn)
	go WsClientSLoop(conn)


}

func SendToServer(msg string){
	SendQueueC.Enqueue(msg)
}

func MessageInbox() Queue {
	return RecieveQueueC
}


func WsClientRLoop(conn *websocket.Conn){
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("client reader died lmao: ", err)
			defer conn.Close()
			break

		}
		RecieveQueueC.Enqueue(string(message));

	}
}

func WsClientSLoop(conn *websocket.Conn){
	for {
		//send shit if there is shit
		if SendQueueC.IsEmpty() == false {
			err := conn.WriteMessage(1, []byte(SendQueueC.Dequeue()))

			if err != nil {
				log.Println("client cant send no mo: ", err)
				defer conn.Close()
				break

			}
		}
	}
}

