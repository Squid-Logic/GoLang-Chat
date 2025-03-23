package main

import (
	"fmt"
	"log"
	"github.com/gorilla/websocket"
)

func main() {
	// Prompt the user to input the server's URL or IP
	var serverURL string
	fmt.Println("Enter the server URL (e.g., ws://localhost:8080 or ws://abcd1234.ngrok.io):")
	fmt.Scanln(&serverURL) // Capture user input for the server URL/port

	// Attempt to connect to the server using the input URL
	conn, _, err := websocket.DefaultDialer.Dial(serverURL, nil)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	fmt.Println("Successfully connected to", serverURL)

	// Here, you can add your client logic to send and receive messages
	for {
		var msg string
		fmt.Println("Enter message to send:")
		fmt.Scanln(&msg) // Capture user input for the message

		// Send the message to the WebSocket server
		err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
		if err != nil {
			log.Println("Error sending message:", err)
			break
		}

		// Wait for a response from the server
		_, response, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}
		fmt.Println("Received message from server:", string(response))
	}
}

