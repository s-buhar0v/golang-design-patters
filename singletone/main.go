package main

import (
	"fmt"
	"sync"

	"github.com/s-buhar0v/golang-design-patters/singletone/uniqobject"
)

const (
	uniqObjectGetTries = 10
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(uniqObjectGetTries)
	for i := 0; i < uniqObjectGetTries; i++ {
		go func() {
			o := uniqobject.GetUniqObject()
			fmt.Printf("Value from object: %d\n", o.RandomInt) // Always prints the same value
			wg.Done()
		}()
	}

	wg.Add(uniqObjectGetTries)
	for i := 0; i < uniqObjectGetTries; i++ {
		go func() {
			o := uniqobject.GetUniqObjectDirty()
			fmt.Printf("Value from dirty object: %d\n", o.RandomInt) // Cant print different values
			wg.Done()
		}()
	}

	wg.Wait()

	uniqObjectsCount := uniqobject.GetUniqObjectsCount()
	uniqObjectsDirtyCount := uniqobject.GetUniqObjectsDirtyCount()

	fmt.Printf(
		"Uniq objects were created: %d\n", uniqObjectsCount,
	) // Always prints "Uniq objects were created: 1"
	fmt.Printf(
		"Uniq dirty objects were created: %d\n", uniqObjectsDirtyCount,
	) // Prints "Uniq dirty objects were created: X",
	// where X >= 1, because GetUniqObjectDirty hasn't mutex,

}
