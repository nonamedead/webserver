package main

type playerStore struct {
	database map[string]int
}

func NewPlayerStore() *playerStore {
	return &playerStore{
		database: make(map[string]int),
	}
}

func (store *playerStore) updatePlayerWin(name string) {
	store.database[name]++
}

func (store *playerStore) getPlayerWin(name string) int {
	return store.database[name]
}

func (store *playerStore) isPlayerPresent(name string) bool {
	if _, ok := store.database[name]; ok {
		return true
	}
	return false
}
