package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
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

type player struct {
	Name string `json:"name"`
	Wins int    `json:"wins"`
}

func (p *playerServer) getLeaderBoard(w http.ResponseWriter) {
	players := make([]string, len(p.store.database))
	id := 0
	for name := range p.store.database {
		players[id] = name
		id++
	}
	fmt.Println("players slice", players)
	sort.SliceStable(players, func(i, j int) bool {
		return p.store.database[players[i]] > p.store.database[players[j]]
	})

	fmt.Println("players slice after sort", players)

	jsPlayers := make([]player, len(players))
	for i, name := range players {
		fmt.Println("woah? ", name, p.store.database[name])
		jsPlayers[i] = player{name, p.store.database[name]}
	}
	fmt.Println("js players", jsPlayers)
	ret, err := json.Marshal(jsPlayers)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println("ret:", string(ret))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	bytes, err := w.Write(ret)
	fmt.Println("bytes:", bytes)
	if err != nil {
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}

}
