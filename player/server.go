package main

import (
	"net/http"
	"strconv"
	"strings"
)

type playerServer struct {
	store *playerStore
}

func newPlayerServer(store *playerStore) *playerServer {
	return &playerServer{store: store}
}

func (p *playerServer) getPlayerWins(w http.ResponseWriter, name string) {
	win := strconv.Itoa(p.store.database[name])
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(win))
}

func (p *playerServer) storePlayerWin(w http.ResponseWriter, name string) {
	p.store.database[name]++
}
func (p *playerServer) getPlayerName(path string) string {
	pathParts := strings.Split(path, "/")
	if len(pathParts) != 3 || pathParts[1] != "players" {
		return ""
	}
	return pathParts[2]
}
