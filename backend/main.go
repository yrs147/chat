package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize: 1024,
    WriteBufferSize: 1024,

    CheckOrigin: func(r *http.Request) bool {return true},
}

func reader(conn *websocket.Conn){
    for{
        messageType , msg , err := conn.ReadMessage()
        if err !=nil{
            log.Println(err)
            return 
        }

        fmt.Println(string(msg))

        if err := conn.WriteMessage(messageType, msg); err!=nil{
            log.Println(err)
            return
        }
    }
}

// define websocket endpoint
func serverWs( w http.ResponseWriter , r *http.Request) {
    fmt.Println(r.Host)

    // upgrade the connection to a websocket
    ws , err := upgrader.Upgrade(w,r,nil)
    if err !=nil{
        log.Println(err)
    }

    // listen indefinitley for new msg coming 
    reader(ws)


}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, "Server")
    })

    //map `/ws` endpoint to the `serverWs` function
    http.HandleFunc("/ws",serverWs)
}

func main() {
    setupRoutes()
    http.ListenAndServe(":8080",nil)
}
