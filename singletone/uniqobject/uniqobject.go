package uniqobject

import (
	"math/rand"
	"sync"
)

var (
	uniqObject      *UniqObject
	uniqObjectDirty *UniqObject

	uniqObjectsCount      = 0
	uniqObjectsDirtyCount = 0

	m = &sync.Mutex{}
)

type UniqObject struct {
	RandomInt int
}

// GetUniqObject has mutex, so it always creates only 1 object event when called from multiple gorutines
func GetUniqObject() *UniqObject {
	m.Lock()
	if uniqObject == nil {
		uniqObjectsCount += 1
		uniqObject = &UniqObject{
			RandomInt: rand.Intn(100),
		}
	}
	m.Unlock()

	return uniqObject
}

// GetUniqObjectDirty hasn't mutex, so it can create more than 1 object when called from multiple gorutines
func GetUniqObjectDirty() *UniqObject {
	if uniqObjectDirty == nil {
		uniqObjectsDirtyCount += 1
		uniqObjectDirty = &UniqObject{
			RandomInt: rand.Intn(100),
		}
	}

	return uniqObjectDirty
}

func GetUniqObjectsCount() int {
	return uniqObjectsCount
}

func GetUniqObjectsDirtyCount() int {
	return uniqObjectsDirtyCount
}
