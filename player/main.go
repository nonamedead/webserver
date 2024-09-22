package main

import (
	"net/http"
)

func (p *playerServer) playerHandler(w http.ResponseWriter, req *http.Request) {

	name := p.getPlayerName(req.URL.Path)

	if req.Method == http.MethodGet {
		p.getPlayerWins(w, name)
	} else if req.Method == http.MethodPost {
		p.storePlayerWin(w, name)
	}
}

func (p *playerServer) leaderBoardHandler(w http.ResponseWriter, req *http.Request) {
	p.getLeaderBoard(w)
}

// todo items: router, handle multiple paths
func main() {
	store := NewPlayerStore()
	server := newPlayerServer(store)
	http.HandleFunc("/players/", server.playerHandler)
	http.HandleFunc("/leaderboard", server.leaderBoardHandler)
	http.ListenAndServe("localhost:9099", nil)
}
