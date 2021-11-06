package vault

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"sync"
)

type Vault struct {
	sessions map[string]int
	users    map[int]string
	mu       sync.Mutex
}

func NewVault() *Vault {
	v := &Vault{
		sessions: make(map[string]int),
		users:    make(map[int]string),
		mu:       sync.Mutex{},
	}
	v.users[1] = "w00p"
	return v
}

func (v *Vault) CheckLogin(user string, password string) (int, error) {
	for k, v := range v.users {
		if v == user {
			return k, nil
		}
	}
	return 0, nil
}

func (v *Vault) StartSessionForUser(userId int) (string, error) {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "", err
	}
	sessId := base64.URLEncoding.EncodeToString(b)
	v.mu.Lock()
	defer v.mu.Unlock()
	v.sessions[sessId] = userId
	return sessId, nil
}

func (v *Vault) GetSessionUserId(sessionId string) (int, error) {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.sessions[sessionId], nil
}

func (v *Vault) GetUserById(userId int) string {
	v.mu.Lock()
	defer v.mu.Unlock()
	return v.users[userId]
}
