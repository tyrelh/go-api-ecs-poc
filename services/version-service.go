package services

// Version service to manage the version identifier of this application

import (
	"fmt"
	"math/rand"
	"sync"
)

var adjectives = []string{"red", "blue", "green", "yellow", "gold", "silver", "crystal", "sapphire", "ruby", "emerald", "firered", "leafgreen"}
var nouns = []string{"horse", "sea", "sky", "forest", "mountain", "river", "tree", "flower", "star", "moon", "sun", "earth"}

var currentVersion = ""
var versionLock sync.Mutex

func GetVersion() string {
	versionLock.Lock()
	defer versionLock.Unlock()
	if currentVersion == "" {
		createVersion()
	}
	return currentVersion
}

func createVersion() {
	currentVersion = generateRandomName()
}

func generateRandomName() string {
	adj := adjectives[rand.Intn(len(adjectives))]
	noun := nouns[rand.Intn(len(nouns))]
	return fmt.Sprintf("%s-%s", adj, noun)
}
