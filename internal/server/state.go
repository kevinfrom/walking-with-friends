package server

import (
	"errors"
	"net"
	"sync"
)

const (
	GRID_SIZE int = 12
)

type Position struct {
	X, Y int
}

type Player struct {
	ID        int
	Conn      net.Conn
	IPAddress string
	Position
}

type GameState struct {
	mu      sync.Mutex
	nextID  int
	config  ServerConfig
	players map[int]*Player
}

func NewGameState(config ServerConfig) *GameState {
	return &GameState{
		config:  config,
		players: make(map[int]*Player),
	}
}

func (g *GameState) AddPlayer(conn net.Conn) (*Player, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if len(g.players) >= g.config.maxPlayers {
		return nil, errors.New("server full")
	}
	g.nextID++
	player := &Player{
		ID:        g.nextID,
		Conn:      conn,
		IPAddress: conn.RemoteAddr().String(),
		Position:  Position{X: 0, Y: 0},
	}
	g.players[player.ID] = player
	return player, nil
}

func (g *GameState) RemovePlayer(p *Player) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.players, p.ID)
}
