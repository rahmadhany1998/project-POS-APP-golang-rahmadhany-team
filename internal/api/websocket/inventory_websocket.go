package websocket

import (
    "encoding/json"
    "net/http"
    "sync"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "go.uber.org/zap"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all connections
    },
}

type Client struct {
    conn *websocket.Conn
    send chan []byte
}

type InventoryHub struct {
    clients    map[*Client]bool
    register   chan *Client
    unregister chan *Client
    broadcast  chan []byte
    logger     *zap.Logger
    mutex      sync.Mutex
}

func NewInventoryHub(logger *zap.Logger) *InventoryHub {
    return &InventoryHub{
        clients:    make(map[*Client]bool),
        register:   make(chan *Client),
        unregister: make(chan *Client),
        broadcast:  make(chan []byte),
        logger:     logger,
    }
}

func (h *InventoryHub) Run() {
    for {
        select {
        case client := <-h.register:
            h.mutex.Lock()
            h.clients[client] = true
            h.mutex.Unlock()
            h.logger.Info("Client connected to WebSocket")
            
        case client := <-h.unregister:
            h.mutex.Lock()
            if _, ok := h.clients[client]; ok {
                delete(h.clients, client)
                close(client.send)
            }
            h.mutex.Unlock()
            h.logger.Info("Client disconnected from WebSocket")
            
        case message := <-h.broadcast:
            h.mutex.Lock()
            for client := range h.clients {
                select {
                case client.send <- message:
                default:
                    close(client.send)
                    delete(h.clients, client)
                }
            }
            h.mutex.Unlock()
        }
    }
}

func (h *InventoryHub) HandleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        h.logger.Error("Failed to set WebSocket connection", zap.Error(err))
        return
    }
    
    client := &Client{
        conn: conn,
        send: make(chan []byte, 256),
    }
    h.register <- client
    
    // Start routines to handle WebSocket communication
    go h.writePump(client)
    go h.readPump(client)
}

func (h *InventoryHub) BroadcastInventoryUpdate(data interface{}) {
    jsonData, err := json.Marshal(data)
    if err != nil {
        h.logger.Error("Failed to marshal inventory data", zap.Error(err))
        return
    }
    
    h.broadcast <- jsonData
}

func (h *InventoryHub) writePump(client *Client) {
    ticker := time.NewTicker(60 * time.Second) // Ping every 60 seconds
    defer func() {
        ticker.Stop()
        client.conn.Close()
    }()
    
    for {
        select {
        case message, ok := <-client.send:
            client.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if !ok {
                client.conn.WriteMessage(websocket.CloseMessage, []byte{})
                return
            }
            
            w, err := client.conn.NextWriter(websocket.TextMessage)
            if err != nil {
                return
            }
            w.Write(message)
            
            if err := w.Close(); err != nil {
                return
            }
            
        case <-ticker.C:
            client.conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
            if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
                return
            }
        }
    }
}

func (h *InventoryHub) readPump(client *Client) {
    defer func() {
        h.unregister <- client
        client.conn.Close()
    }()
    
    client.conn.SetReadLimit(512)
    client.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
    client.conn.SetPongHandler(func(string) error {
        client.conn.SetReadDeadline(time.Now().Add(60 * time.Second))
        return nil
    })
    
    for {
        _, _, err := client.conn.ReadMessage()
        if err != nil {
            if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
                h.logger.Error("Unexpected WebSocket close error", zap.Error(err))
            }
            break
        }
    }
}