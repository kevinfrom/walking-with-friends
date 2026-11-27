package server

import (
	"errors"
	"net"
	"sync"
)

const (
	MAX_PLAYERS int = 4
)

type Player struct {
	ID   int
	Conn net.Conn
	X, Y int
}

type GameState struct {
	mu      sync.Mutex
	players map[int]*Player
	nextID  int
}

func NewGameState() *GameState {
	return &GameState{
		players: make(map[int]*Player),
	}
}

func (g *GameState) AddPlayer(conn net.Conn) (*Player, error) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if len(g.players) >= MAX_PLAYERS {
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
