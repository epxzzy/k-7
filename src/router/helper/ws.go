package helper

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)
var SendQueue = Queue{Size: 10};
var RecieveQueue = Queue{Size: 10};

var websocketer = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true;
	},
}

func initWsServer(){
	flag.Parse();
	log.SetFlags(0);
	http.HandleFunc("/k7_client", httpshandler);
	log.Fatal(http.ListenAndServe(":8080", nil))
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

	defer conn.Close();

	WsServerLoop(conn)
}


func WsServerLoop(conn *websocket.Conn){
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("reader died lmao: ", err)
			break
		}

		RecieveQueue.Enqueue(string(message));
		//log.Printf("recv: %s", message)
		//send shit if there is shit
		if SendQueue.IsEmpty() == false {
			err = conn.WriteMessage(mt, []byte(SendQueue.Dequeue()))

			if err != nil {
				log.Println("cant send no mo: ", err)
				break
			}

		}
	}

}
