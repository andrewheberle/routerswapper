// Package routerswapper implements a simple process to swap a [http.Handler]
// during runtime. This is aimed at allowing [http.Handler] changes based on
// a configuration change or update.
//
// By creating a new [Swapper] via [New] the [Swapper] can be used in place
// of a [http.Handler] with the ability to [Swapper.Swap] the underlying
// [http.Handler] at runtime.
package routerswapper

import (
	"net/http"
	"sync"
)

// A Swapper can be used in place of any [http.Handler].
//
// The Swapper must be initialised with a call to [New].
//
// After creation, changes to the underlying http.Handler must be completed
// by a call to [Swapper.Swap] to ensure the change is done in a safe manner with a
// lock.
type Swapper struct {
	mu      sync.RWMutex
	handler http.Handler
}

// Deprecated: Use [Swapper] instead
type RouterSwapper = Swapper

// Swap replaces the current [http.Handler] with a new version ensuring a lock is
// taken so the swap is safe for concurrent use.
func (rs *Swapper) Swap(handler http.Handler) {
	rs.mu.Lock()
	rs.handler = handler
	rs.mu.Unlock()
}

// New creates a new [Swapper] based on the provided [http.Handler] which can
// then be used where any [http.Handler] would be used
func New(handler http.Handler) *Swapper {
	rs := new(Swapper)
	rs.handler = handler
	return rs
}

// ServeHTTP method for the [Swapper]
func (rs *Swapper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rs.mu.RLock()
	handler := rs.handler
	rs.mu.RUnlock()

	handler.ServeHTTP(w, r)
}
