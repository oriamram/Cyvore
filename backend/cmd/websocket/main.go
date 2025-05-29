package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	_ "modernc.org/sqlite"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins for now
		},
	}

	// Store all active connections
	clients = make(map[*websocket.Conn]bool)
	clientsMutex sync.RWMutex
)

type Asset struct {
	ID        string `json:"id"`
	CreatedAt string `json:"created_at"`
	Type      string `json:"type"`
	Content   string `json:"content"`
	LastSeen  string `json:"last_seen"`
}

type Relation struct {
	ID          string `json:"id"`
	CreatedAt   string `json:"created_at"`
	Type        string `json:"type"`
	FromAssetID string `json:"from_asset_id"`
	ToAssetID   string `json:"to_asset_id"`
	LastSeen    string `json:"last_seen"`
}

type AmassData struct {
	Assets    []Asset    `json:"assets"`
	Relations []Relation `json:"relations"`
	Total     struct {
		Assets    int `json:"assets"`
		Relations int `json:"relations"`
	} `json:"total"`
}

type ClientState struct {
	AssetPage     int    `json:"assetPage"`
	AssetPageSize int    `json:"assetPageSize"`
	AssetType     string `json:"assetType"`
	AssetFilter   string `json:"assetFilter"`
	
	RelationPage     int    `json:"relationPage"`
	RelationPageSize int    `json:"relationPageSize"`
	RelationType     string `json:"relationType"`
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}

	// Add client to the map with initial state
	clientsMutex.Lock()
	clients[conn] = true
	clientsMutex.Unlock()

	// Remove client when they disconnect
	defer func() {
		clientsMutex.Lock()
		delete(clients, conn)
		clientsMutex.Unlock()
		conn.Close()
	}()

	// Send initial data
	if err := sendCurrentData(conn, ClientState{
		AssetPage:     1,
		AssetPageSize: 50,
		RelationPage:  1,
		RelationPageSize: 50,
	}); err != nil {
		log.Printf("Failed to send initial data: %v", err)
		return
	}

	// Handle client messages
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}

		var state ClientState
		if err := json.Unmarshal(message, &state); err != nil {
			log.Printf("Failed to parse client state: %v", err)
			continue
		}

		if err := sendCurrentData(conn, state); err != nil {
			log.Printf("Failed to send data: %v", err)
			break
		}
	}
}

func sendCurrentData(conn *websocket.Conn, state ClientState) error {
	db, err := sql.Open("sqlite", "data/amass/amass.sqlite")
	if err != nil {
		return err
	}
	defer db.Close()

	// Get total counts
	var totalAssets, totalRelations int
	err = db.QueryRow("SELECT COUNT(*) FROM assets").Scan(&totalAssets)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT COUNT(*) FROM relations").Scan(&totalRelations)
	if err != nil {
		return err
	}

	// Build asset query
	assetQuery := "SELECT id, created_at, type, content, last_seen FROM assets"
	assetArgs := []interface{}{}
	if state.AssetType != "" {
		assetQuery += " WHERE type = ?"
		assetArgs = append(assetArgs, state.AssetType)
	}
	if state.AssetFilter != "" {
		if len(assetArgs) > 0 {
			assetQuery += " AND"
		} else {
			assetQuery += " WHERE"
		}
		assetQuery += " content LIKE ?"
		assetArgs = append(assetArgs, "%"+state.AssetFilter+"%")
	}
	assetQuery += " LIMIT ? OFFSET ?"
	assetArgs = append(assetArgs, state.AssetPageSize, (state.AssetPage-1)*state.AssetPageSize)

	// Query assets
	rows, err := db.Query(assetQuery, assetArgs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	var assets []Asset
	for rows.Next() {
		var asset Asset
		if err := rows.Scan(&asset.ID, &asset.CreatedAt, &asset.Type, &asset.Content, &asset.LastSeen); err != nil {
			return err
		}
		assets = append(assets, asset)
	}

	// Build relation query
	relationQuery := "SELECT id, created_at, type, from_asset_id, to_asset_id, last_seen FROM relations"
	relationArgs := []interface{}{}
	if state.RelationType != "" {
		relationQuery += " WHERE type = ?"
		relationArgs = append(relationArgs, state.RelationType)
	}
	relationQuery += " LIMIT ? OFFSET ?"
	relationArgs = append(relationArgs, state.RelationPageSize, (state.RelationPage-1)*state.RelationPageSize)

	// Query relations
	rows, err = db.Query(relationQuery, relationArgs...)
	if err != nil {
		return err
	}
	defer rows.Close()

	var relations []Relation
	for rows.Next() {
		var relation Relation
		if err := rows.Scan(&relation.ID, &relation.CreatedAt, &relation.Type, &relation.FromAssetID, &relation.ToAssetID, &relation.LastSeen); err != nil {
			return err
		}
		relations = append(relations, relation)
	}

	data := AmassData{
		Assets:    assets,
		Relations: relations,
	}
	data.Total.Assets = totalAssets
	data.Total.Relations = totalRelations

	return conn.WriteJSON(data)
}

func main() {
	// Setup websocket endpoint
	http.HandleFunc("/ws", handleWebSocket)

	// Start server
	log.Println("WebSocket server starting on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 