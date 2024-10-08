package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type HyperionStreamClient struct {
	socket             *websocket.Conn
	socketURL          string
	lastReceivedBlock  int
	dataQueue          []IncomingData
	options            HyperionClientOptions
	libDataQueue       []IncomingData
	reversibleBuffer   []IncomingData
	onDataAsync        AsyncHandlerFunction
	onLibDataAsync     AsyncHandlerFunction
	online             bool
	savedRequests      []SavedRequest
	eventListeners     map[string][]EventListener
	tempEventListeners map[string][]EventListener
}

// Constructor for HyperionStreamClient
func NewHyperionStreamClient(options HyperionClientOptions) *HyperionStreamClient {
	client := &HyperionStreamClient{
		options:            options,
		dataQueue:          []IncomingData{},
		libDataQueue:       []IncomingData{},
		reversibleBuffer:   []IncomingData{},
		eventListeners:     make(map[string][]EventListener),
		tempEventListeners: make(map[string][]EventListener),
	}
	return client
}

// Disconnect from the WebSocket server
func (client *HyperionStreamClient) Disconnect() {
	if client.socket != nil {
		client.lastReceivedBlock = 0
		client.socket.Close()
		client.savedRequests = []SavedRequest{}
	} else {
		fmt.Println("Nothing to disconnect!")
	}
}

// Set the WebSocket endpoint
func (client *HyperionStreamClient) SetEndpoint(endpoint string) {
	client.socketURL = endpoint
}

// Connect to the WebSocket server
func (client *HyperionStreamClient) Connect() error {
	if client.socketURL == "" {
		return fmt.Errorf("endpoint was not defined")
	}
	var err error
	client.socket, _, err = websocket.DefaultDialer.Dial(client.socketURL, nil)
	if err != nil {
		return err
	}

	client.online = true
	go client.listenToMessages()
	go client.listenToEvents()

	return nil
}

// Listen to incoming messages
func (client *HyperionStreamClient) listenToMessages() {
	for {
		_, msg, err := client.socket.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			return
		}

		var incomingData IncomingData
		if err := json.Unmarshal(msg, &incomingData); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}

		client.dataQueue = append(client.dataQueue, incomingData)
		if client.onDataAsync != nil {
			if err := client.onDataAsync(incomingData); err != nil {
				log.Println("Error in onDataAsync handler:", err)
			}
		}
	}
}

// Listen to events (to be implemented)
func (client *HyperionStreamClient) listenToEvents() {
	// Implement event listening logic if needed
}

// Stream actions to the server
func (client *HyperionStreamClient) StreamActions(request StreamActionsRequest) error {
	if client.socket != nil && client.online {
		data, err := json.Marshal(request)
		if err != nil {
			return err
		}

		err = client.socket.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			return err
		}

		client.savedRequests = append(client.savedRequests, SavedRequest{Type: "action", Req: request})
		return nil
	}
	return fmt.Errorf("client is not connected")
}

// Stream deltas to the server
func (client *HyperionStreamClient) StreamDeltas(request StreamDeltasRequest) error {
	if client.socket != nil && client.online {
		data, err := json.Marshal(request)
		if err != nil {
			return err
		}

		err = client.socket.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			return err
		}

		client.savedRequests = append(client.savedRequests, SavedRequest{Type: "delta", Req: request})
		return nil
	}
	return fmt.Errorf("client is not connected")
}

// Public method to retrieve the data queue
func (client *HyperionStreamClient) GetDataQueue() []IncomingData {
	return client.dataQueue
}

// Public method to clear the data queue
func (client *HyperionStreamClient) ClearDataQueue() {
	client.dataQueue = []IncomingData{}
}
