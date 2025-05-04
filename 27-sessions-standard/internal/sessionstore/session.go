package sessionstore

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

var sessions = make(map[string]string)
var mu sync.RWMutex

func CreateSession(username string) string {
	token := generateToken()
	mu.Lock()
	sessions[token] = username
	mu.Unlock()
	return token
}

func GetSession(token string) (string, bool) {
	mu.RLock()
	defer mu.RUnlock()
	username, ok := sessions[token]
	return username, ok
}

func DeleteSession(token string) {
	mu.Lock()
	delete(sessions, token)
	mu.Unlock()
}

func generateToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b)
}
