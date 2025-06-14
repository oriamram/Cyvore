package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	_ "modernc.org/sqlite"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	// Store all active connections and their states
	clients = make(map[*websocket.Conn]*ClientState)
	clientsMutex sync.RWMutex

	// Add these variables to track last update times
	lastAssetCount    int
	lastRelationCount int
	lastUpdateMutex   sync.RWMutex
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

	SortColumn    string `json:"sortColumn"`
	SortDirection string `json:"sortDirection"`
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade connection: %v", err)
		return
	}

	// Initialize client state
	initialState := &ClientState{
		AssetPage:     1,
		AssetPageSize: 50,
		RelationPage:  1,
		RelationPageSize: 50,
	}

	// Add client to the map with initial state
	clientsMutex.Lock()
	clients[conn] = initialState
	clientsMutex.Unlock()

	// Remove client when they disconnect
	defer func() {
		clientsMutex.Lock()
		delete(clients, conn)
		clientsMutex.Unlock()
		conn.Close()
	}()

	// Send initial data
	if err := sendCurrentData(conn, initialState); err != nil {
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

		// Update client's state
		clientsMutex.Lock()
		clients[conn] = &state
		clientsMutex.Unlock()

		if err := sendCurrentData(conn, &state); err != nil {
			log.Printf("Failed to send data: %v", err)
			break
		}
	}
}

func sendCurrentData(conn *websocket.Conn, state *ClientState) error {
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

	// Add sorting if specified
	if state.SortColumn != "" {
		assetQuery += " ORDER BY " + state.SortColumn
		if state.SortDirection == "desc" {
			assetQuery += " DESC"
		}
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
	relationQuery := `
		SELECT 
			r.id, 
			r.created_at, 
			r.type, 
			r.from_asset_id,
			fa.content as from_asset_content,
			r.to_asset_id,
			ta.content as to_asset_content,
			r.last_seen 
		FROM relations r
		LEFT JOIN assets fa ON r.from_asset_id = fa.id
		LEFT JOIN assets ta ON r.to_asset_id = ta.id`
	relationArgs := []interface{}{}
	if state.RelationType != "" {
		relationQuery += " WHERE r.type = ?"
		relationArgs = append(relationArgs, state.RelationType)
	}

	// Add sorting if specified
	if state.SortColumn != "" {
		relationQuery += " ORDER BY r." + state.SortColumn
		if state.SortDirection == "desc" {
			relationQuery += " DESC"
		}
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
		var fromContent, toContent string
		if err := rows.Scan(
			&relation.ID,
			&relation.CreatedAt,
			&relation.Type,
			&relation.FromAssetID,
			&fromContent,
			&relation.ToAssetID,
			&toContent,
			&relation.LastSeen,
		); err != nil {
			return err
		}
		// Format the content to show in the UI
		relation.FromAssetID = formatAssetContent(fromContent)
		relation.ToAssetID = formatAssetContent(toContent)
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

func formatAssetContent(content string) string {
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(content), &parsed); err != nil {
		return content
	}
	
	if name, ok := parsed["name"].(string); ok {
		return name
	}
	if address, ok := parsed["address"].(string); ok {
		if typeStr, ok := parsed["type"].(string); ok {
			return fmt.Sprintf("%s (%s)", address, typeStr)
		}
		return address
	}
	return content
}

func monitorDatabaseChanges() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		db, err := sql.Open("sqlite", "data/amass/amass.sqlite")
		if err != nil {
			log.Printf("Failed to open database: %v", err)
			continue
		}

		// Get current counts
		var currentAssetCount, currentRelationCount int
		err = db.QueryRow("SELECT COUNT(*) FROM assets").Scan(&currentAssetCount)
		if err != nil {
			log.Printf("Failed to get asset count: %v", err)
			db.Close()
			continue
		}
		err = db.QueryRow("SELECT COUNT(*) FROM relations").Scan(&currentRelationCount)
		if err != nil {
			log.Printf("Failed to get relation count: %v", err)
			db.Close()
			continue
		}

		// Check if counts have changed
		lastUpdateMutex.RLock()
		assetsChanged := currentAssetCount != lastAssetCount
		relationsChanged := currentRelationCount != lastRelationCount
		lastUpdateMutex.RUnlock()

		if assetsChanged || relationsChanged {
			// Update last counts
			lastUpdateMutex.Lock()
			lastAssetCount = currentAssetCount
			lastRelationCount = currentRelationCount
			lastUpdateMutex.Unlock()

			// Broadcast update to all clients
			clientsMutex.RLock()
			for client, state := range clients {
				// Send updated data using client's current state
				if err := sendCurrentData(client, state); err != nil {
					log.Printf("Failed to send update to client: %v", err)
					// Remove disconnected client
					clientsMutex.RUnlock()
					clientsMutex.Lock()
					delete(clients, client)
					clientsMutex.Unlock()
					clientsMutex.RLock()
				}
			}
			clientsMutex.RUnlock()
		}

		db.Close()
	}
}

func main() {
	// Initialize last counts
	db, err := sql.Open("sqlite", "data/amass/amass.sqlite")
	if err == nil {
		db.QueryRow("SELECT COUNT(*) FROM assets").Scan(&lastAssetCount)
		db.QueryRow("SELECT COUNT(*) FROM relations").Scan(&lastRelationCount)
		db.Close()
	}

	// Start database monitoring in a goroutine
	go monitorDatabaseChanges()

	// Setup websocket endpoint
	http.HandleFunc("/ws", handleWebSocket)

	// Start server
	log.Println("WebSocket server starting on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} 