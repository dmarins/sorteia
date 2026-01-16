package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type (
	ShuffleRequest struct {
		Name           string   `json:"name"`
		Players        []string `json:"players"`
		PlayersPerTeam int      `json:"players_per_team"` // Ex: 5 para Futsal
	}

	Team struct {
		Name    string   `json:"name"`
		Players []string `json:"players"`
	}

	ShuffleResponse struct {
		ID    string   `json:"id"`
		Teams []Team   `json:"teams"`
		Bench []string `json:"bench"` // Reservas
	}

	Match struct {
		ID        string    `gorm:"primaryKey" json:"id"`
		Name      string    `json:"name"`
		Result    string    `json:"result"` // Guardaremos o JSON dos times como string/text
		CreatedAt time.Time `json:"created_at"`
	}
)

func initDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar no banco de dados")
	}
	db.AutoMigrate(&Match{})
	return db
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://sorteia-ui.vercel.app/") // Em produção, mude "*" para o domínio da Vercel
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	db := initDB()
	r := gin.Default()
	r.Use(CORSMiddleware())

	// POST /shuffle - Realiza o sorteio e salva
	r.POST("/shuffle", func(c *gin.Context) {
		var req ShuffleRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
			return
		}

		// Chama a nova lógica de distribuição
		teams, bench := DistributeTeams(req.Players, req.PlayersPerTeam)

		// Serializa o resultado completo (times + banco) para o SQLite
		resultData := gin.H{
			"teams": teams,
			"bench": bench,
		}

		resultJSON, _ := json.Marshal(resultData)

		match := Match{
			ID:     uuid.New().String()[:8],
			Name:   req.Name,
			Result: string(resultJSON),
		}

		db.Create(&match)

		c.JSON(http.StatusOK, gin.H{
			"id":    match.ID,
			"teams": teams,
			"bench": bench,
		})
	})

	// GET /match/:id - Recupera um sorteio já feito
	r.GET("/match/:id", func(c *gin.Context) {
		id := c.Param("id")
		var match Match
		if err := db.First(&match, "id = ?", id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Sorteio não encontrado"})
			return
		}

		var resultData map[string]any
		json.Unmarshal([]byte(match.Result), &resultData)

		c.JSON(http.StatusOK, gin.H{
			"id":    match.ID,
			"name":  match.Name,
			"date":  match.CreatedAt,
			"teams": resultData["teams"],
			"bench": resultData["bench"],
		})
	})

	r.Run(":8080")
}

func ShufflePlayers(players []string) []string {
	n := len(players)
	res := make([]string, n)
	copy(res, players)

	for i := n - 1; i > 0; i-- {
		// Gera um número aleatório seguro entre 0 e i
		randomInt, _ := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		j := randomInt.Int64()
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func DistributeTeams(allPlayers []string, perTeam int) ([]Team, []string) {
	// 1. Embaralha tudo usando a nossa função ShufflePlayers (com crypto/rand)
	shuffled := ShufflePlayers(allPlayers)

	totalPlayers := len(shuffled)
	if perTeam <= 0 {
		perTeam = totalPlayers // Evita divisão por zero
	}

	numTeams := totalPlayers / perTeam
	var teams []Team

	// 2. Preenche os times
	for i := 0; i < numTeams; i++ {
		start := i * perTeam
		end := start + perTeam

		teams = append(teams, Team{
			Name:    fmt.Sprintf("Time %d", i+1),
			Players: shuffled[start:end],
		})
	}

	// 3. O que sobrar vira reserva
	var bench []string
	if totalPlayers%perTeam != 0 {
		bench = shuffled[numTeams*perTeam:]
	}

	return teams, bench
}
