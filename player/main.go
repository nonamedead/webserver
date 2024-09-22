package main

import (
	"net/http"
)

func (p *playerServer) serverHandler(w http.ResponseWriter, req *http.Request) {

	name := p.getPlayerName(req.URL.Path)

	if req.Method == http.MethodGet {
		p.getPlayerWins(w, name)
	} else if req.Method == http.MethodPost {
		p.storePlayerWin(w, name)
	}
}
func main() {
	store := NewPlayerStore()
	server := newPlayerServer(store)
	http.HandleFunc("/players/", server.serverHandler)
	http.ListenAndServe("localhost:9099", nil)
}
