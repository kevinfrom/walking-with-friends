package server

import (
	"errors"
	"net"
	"sync"
)

type Player struct {
	ID   int
	Conn net.Conn
	X, Y int
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
	player := &Player{
		ID:   g.nextID,
		Conn: conn,
		X:    0,
		Y:    0,
	}
	g.nextID++
	g.players[player.ID] = player
	return player, nil
}

func (g *GameState) RemovePlayer(p *Player) {
	g.mu.Lock()
	defer g.mu.Unlock()
	delete(g.players, p.ID)
}
