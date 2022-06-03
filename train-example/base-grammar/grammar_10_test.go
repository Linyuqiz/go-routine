package base_grammar

import (
	"sync"
	"testing"
)

func TestGrammar10_00(t *testing.T) {
	mutex := sync.Mutex{}
	mutex.Lock()
	mutex.Unlock()
}
