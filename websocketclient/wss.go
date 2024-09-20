package websocketclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"time"

	"upspowershow/models"

	"github.com/gorilla/websocket"
)

var userId = "change this to your user id"
var upsId = "change this to your ups id"
var upsInfoChan = make(chan models.UPSInfo)
var lastInfo models.UPSInfo = models.UPSInfo{Percent: -1, ChargePow: 0, ChargeRemainTime: 0, DischargePow: 0, DischargeRemainTime: 0}
var message = fmt.Sprintf(`{"userId":"%s","content":"ugreenSocketConnection"}`, userId)
var conn *websocket.Conn

func readMessages() {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		// log.Printf("Received: %s", msg)

		var upsInfo models.UPSInfo
		var genericMessage map[string]interface{}
		if err := json.Unmarshal(msg, &genericMessage); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		if _, ok := genericMessage["deviceModelName"]; ok &&
			genericMessage["deviceModelName"] != nil &&
			genericMessage["battery_percentage"] != nil &&
			genericMessage["bluetooth_id"] == nil {
			continue
		}

		if _, ok := genericMessage["message"]; ok &&
			genericMessage["message"] != nil {
			continue
		}

		if err := json.Unmarshal(msg, &upsInfo); err != nil {
			log.Println("Error unmarshalling message:", err)
			continue
		}

		if lastInfo.Percent == -1 {
			lastInfo = upsInfo
			upsInfoChan <- upsInfo
			continue
		}

		if lastInfo.Percent-upsInfo.Percent >= 3 || upsInfo.Percent-lastInfo.Percent >= 3 {
			continue
		}
		lastInfo = upsInfo
		upsInfoChan <- upsInfo
	}
}

func handleHeartbeat() {
	heartbeatMessage := message
	for {
		err := conn.WriteMessage(websocket.TextMessage, []byte(heartbeatMessage))
		if err != nil {
			log.Println("Ping error:", err)
			return
		}
		time.Sleep(time.Second * 5)
	}
}

func GetUPSInfo() <-chan models.UPSInfo {
	return upsInfoChan
}

func Disconnect() {
	conn.Close()
}

func StartWebSocket() {
	u := url.URL{
		Scheme: "wss",
		Host:   "powerapi.ugreengroup.com:8089",
		Path:   fmt.Sprintf("/app/device/websocket/%s/%s", userId, upsId),
	}
	var err error
	conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("Dial error:", err)
	}

	initialMessage := message
	err = conn.WriteMessage(websocket.TextMessage, []byte(initialMessage))
	if err != nil {
		log.Println("Write error:", err)
		return
	}

	go handleHeartbeat()
	go readMessages()
}
