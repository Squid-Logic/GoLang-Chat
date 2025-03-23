package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "github.com/gorilla/websocket"
    "sync"
)

// Define a structure to hold the client's connection and ID
type Client struct {
    conn *websocket.Conn
    id   string
}

var clients = make(map[string]*Client) // Store all clients
var clientsMutex sync.Mutex            // Mutex to handle concurrent access to clients map

// Broadcast a message to all clients except the sender
func broadcastMessage(message string, senderID string) {
    clientsMutex.Lock()
    defer clientsMutex.Unlock()

    for id, client := range clients {
        // Don't send the message to the sender
        if id != senderID {
            err := client.conn.WriteMessage(websocket.TextMessage, []byte(message))
            if err != nil {
                log.Println("Error sending message to client", id, ":", err)
            }
        }
    }
}

// Define the WebSocket upgrader to handle the HTTP to WebSocket upgrade
var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all connections
    },
}

func handleClientConnection(w http.ResponseWriter, r *http.Request) {
    // Upgrade the HTTP connection to a WebSocket connection
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println("Error upgrading connection:", err)
        return
    }
    defer conn.Close()

    // Generate a unique ID for the client (this could be more sophisticated)
    clientID := fmt.Sprintf("%s", conn.RemoteAddr())

    // Add the new client to the list
    clientsMutex.Lock()
    clients[clientID] = &Client{conn: conn, id: clientID}
    clientsMutex.Unlock()

    defer func() {
        // Remove the client when they disconnect
        clientsMutex.Lock()
        delete(clients, clientID)
        clientsMutex.Unlock()
    }()

    // Listen for incoming messages from the client
    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            log.Println("Error reading message:", err)
            break
        }

        // Convert the message from byte array to string
        message := fmt.Sprintf("Client %s says: %s", clientID, string(msg))

        // Broadcast the message to all other clients, including the sender's ID
        broadcastMessage(message, clientID)
    }
}

func main() {
    // Get the port from the environment variable, if set
    port := os.Getenv("PORT")
    if port == "" {
        // Default to 8080 if no environment variable is set
        port = ":8080"
    }

    // Handle incoming HTTP requests and upgrade them to WebSocket connections
    http.HandleFunc("/", handleClientConnection)

    // Listen on the port Render assigns or 8080 if not set
    log.Printf("Server started on %s\n", port)
    log.Fatal(http.ListenAndServe(port, nil))
}

