package helper

import (
	"log"
	"net/http"
	//"flag"
	"github.com/gorilla/websocket"
)
var SendQueue = Queue{Size: 10};
var RecieveQueue = Queue{Size: 10};

var websocketer = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true;
	},
}


//var addr = flag.String("addr", "localhost:8080", "http service address")

func InitWsServer() bool {
	//flag.Parse()
	http.HandleFunc("/tunnel", httpshandler);
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
	return true
}

func SendToClient(msg string){
	SendQueue.Enqueue(msg)
}

func MessageInbox() Queue {
	return RecieveQueue
}

func httpshandler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketer.Upgrade(w, r, nil);

	if err != nil {
		log.Println("websocketer died: ", err)
		return;
	}


	go WsServerSLoop(conn)
	go WsServerRLoop(conn)


}


func WsServerRLoop(conn *websocket.Conn){
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("reader died lmao: ", err)
			defer conn.Close();
			break

		}

		RecieveQueue.Enqueue(string(message));
		//log.Printf("recv: %s", message)
		//send shit if there is shit
	}


}
func WsServerSLoop(conn *websocket.Conn){
	for {
		//log.Printf("recv: %s", message)
		//send shit if there is shit
		if SendQueue.IsEmpty() == false {
			err := conn.WriteMessage(1, []byte(SendQueue.Dequeue()))

			if err != nil {
				log.Println("cant send no mo: ", err)
				defer conn.Close();
				break

			}
		}
	}
}

