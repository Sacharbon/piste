package main

import (
    "encoding/json"
    "fmt"
    "log"
    "os"
    "os/signal"

    "github.com/gorilla/websocket"
)

type Payload struct {
    Method string `json:"method"`
    Keys []string `json:"keys,omitempty"`
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("./piste [ADDRESS]")
        return
    }

    conn, _, err := websocket.DefaultDialer.Dial("wss://pumpportal.fun/api/data", nil)
    if err != nil {
        log.Fatal("websocket connection failed:", err)
    }
    defer conn.Close()

    subscribeAccountTrade := Payload{
        Method: "subscribeAccountTrade",
        Keys: os.Args[1:],
    }
    err = conn.WriteJSON(subscribeAccountTrade)
    if err != nil {
        log.Fatal("subscribeAccountTrade failed:", err);
    }

    interrupt := make(chan os.Signal, 1)
    signal.Notify(interrupt, os.Interrupt)

    done := make(chan struct{})

    go func() {
        defer close(done)
        for {
            _, message, err := conn.ReadMessage()
            if err != nil {
                log.Println("Failed reading message:", err)
                return
            }
            var data map[string]interface{}
            if err := json.Unmarshal(message, &data); err != nil {
                log.Println("Failed parsing:", err)
            } else {
                if data["message"] != nil {
                    fmt.Println(data["message"])
                } else {
                    fmt.Println(data["traderPublicKey"], ":",data["txType"], data["tokenAmount"], data["bondingCurveKey"])
                }
            }
        }
    }()

    for {
        select {
            case <- done:
                return
            case <- interrupt:
                err := conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
                if err != nil {
                  log.Println("closing websocket failed:", err)
                }
                select {
                case <- done:
                default:

                }
                return
        }
    }
}
