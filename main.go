package main

import (
	flatbuffer "clubsocket/eventHandler"
	"context"
	"fmt"
	"log"
	"time"

	"nhooyr.io/websocket"
)

type WebSocketConfig struct {
	URL     string
	Token   string
	ClubId  string
	UserId  string
	Headers map[string]string
}

func main() {
	config := WebSocketConfig{
		URL:    "wss://stage-api.getstan.app/ws/club",
		Token:  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZ21pUHJvZmlsZUlkIjozNTcwMCwiZXhwIjoxNzMzOTA2Mzg4LCJmcmVlZmlyZVByb2ZpbGVJZCI6MzU3MDEsImlhdCI6MTczMzgxOTk4OCwiaWQiOiIyMDU5MCJ9.CWrnGf9g8gwG-wx49gL4i600U36BtWt4BVsamFCUYpY",
		ClubId: "7XWWZ4SS",
		UserId: "20590",
		Headers: map[string]string{
			"Authorization": "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJiZ21pUHJvZmlsZUlkIjozNTcwMCwiZXhwIjoxNzMzOTA2Mzg4LCJmcmVlZmlyZVByb2ZpbGVJZCI6MzU3MDEsImlhdCI6MTczMzgxOTk4OCwiaWQiOiIyMDU5MCJ9.CWrnGf9g8gwG-wx49gL4i600U36BtWt4BVsamFCUYpY",
		},
	}

	err := connectWebSocket(config)
	if err != nil {
		log.Fatalf("Error in connection === Could not connect to the club : Error : %v", err)
	}
}

func connectWebSocket(config WebSocketConfig) error {
	socketUrl := fmt.Sprintf("%s?clubId=%s&userId=%s&token=%s", config.URL, config.ClubId, config.UserId, config.Token)
	log.Printf("The socker uril will be %v", socketUrl)
	log.Printf("The club id will be %v", config.ClubId)

	ctx := context.Background()

	options := &websocket.DialOptions{
		HTTPHeader: make(map[string][]string),
	}

	for key, value := range config.Headers {
		options.HTTPHeader[key] = []string{value}
	}

	conn, _, err := websocket.Dial(ctx, socketUrl, options)
	if err != nil {
		return fmt.Errorf("failed to connect: %v", err)
	}
	log.Println("Connected successfully!")
	defer conn.Close(websocket.StatusInternalError, "Connection closed")

	// join a club after web socket connection
	clubJoinData := flatbuffer.JoinClubEvent(config.ClubId)
	err = conn.Write(ctx, websocket.MessageBinary, clubJoinData)
	if err != nil {
		log.Fatalf("Error sending join club event: %v", err)
	}
	log.Printf("The user joined the club successfully")

	time.Sleep(5 * time.Second)

	// send a message to the club
	messageEventData := flatbuffer.SendMessageEvent(config.ClubId, "Check for script!", "message")
	err = conn.Write(ctx, websocket.MessageBinary, messageEventData)
	if err != nil {
		log.Fatalf("Error sending message event: %v", err)
	}
	log.Printf("The message is sent successfully")

	go func() {
		for {
			messageType, response, err := conn.Read(ctx)
			if err != nil {
				if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
					log.Println("Connection closed normally.")
					return
				}
				log.Printf("Error reading response: %v", err)
				return
			}

			log.Printf("Received message (type %v): %s", messageType, string(response))
		}
	}()

	time.Sleep(30 * time.Second)

	// exit the club and close the connection
	exitClubData := flatbuffer.ExitClubEvent(config.ClubId, config.UserId)
	err = conn.Write(ctx, websocket.MessageBinary, exitClubData)
	if err != nil {
		log.Fatalf("Error sending exit club event: %v", err)
	}
	log.Printf("The user exited successfully")

	time.Sleep(800 * time.Second)

	return nil
}
