package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
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
}

func AmassDataHandler(c *gin.Context) {
	// Open SQLite database
	db, err := sql.Open("sqlite", "data/amass/amass.sqlite")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to open database: %v", err)})
		return
	}
	defer db.Close()

	// Query all assets
	rows, err := db.Query("SELECT id, created_at, type, content, last_seen FROM assets")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to query assets: %v", err)})
		return
	}
	defer rows.Close()

	var assets []Asset
	for rows.Next() {
		var asset Asset
		if err := rows.Scan(&asset.ID, &asset.CreatedAt, &asset.Type, &asset.Content, &asset.LastSeen); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to scan asset: %v", err)})
			return
		}
		assets = append(assets, asset)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error iterating assets: %v", err)})
		return
	}

	// Query all relations
	rows, err = db.Query("SELECT id, created_at, type, from_asset_id, to_asset_id, last_seen FROM relations")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to query relations: %v", err)})
		return
	}
	defer rows.Close()

	var relations []Relation
	for rows.Next() {
		var relation Relation
		if err := rows.Scan(&relation.ID, &relation.CreatedAt, &relation.Type, &relation.FromAssetID, &relation.ToAssetID, &relation.LastSeen); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to scan relation: %v", err)})
			return
		}
		relations = append(relations, relation)
	}

	if err = rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error iterating relations: %v", err)})
		return
	}

	c.JSON(http.StatusOK, AmassData{
		Assets:    assets,
		Relations: relations,
	})
} 