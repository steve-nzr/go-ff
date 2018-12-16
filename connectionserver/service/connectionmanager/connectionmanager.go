package connectionmanager

import (
	"go-ff/common/service/external"
	"sync"
)

var clients = make(map[uint32]*external.Client)
var clientsMut sync.RWMutex

// Add a new Client
func Add(c *external.Client) bool {
	clientsMut.Lock()
	defer clientsMut.Unlock()

	// Exists ?
	_, ok := clients[c.ID]
	if ok == true {
		return false
	}

	clients[c.ID] = c
	return true
}

// Remove a Client
func Remove(c *external.Client) {
	clientsMut.Lock()
	defer clientsMut.Unlock()

	delete(clients, c.ID)
}

// Get Client by ID
func Get(id uint32) *external.Client {
	client, ok := clients[id]
	if ok == false {
		return nil
	}

	return client
}
