// Package routerswapper implements a simple process to swap a Golang HTTP
// router (gorilla/mux etc) during runtime. This is aimed at allowing
// route/handler changes based on a configuration change or update.
package routerswapper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func test200Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func test404Handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func TestSwap(t *testing.T) {

	assert := assert.New(t)

	// new router
	rt := http.NewServeMux()
	rt.HandleFunc("/", test200Handler)
	rs := New(rt)

	ts := httptest.NewServer(rs)
	defer ts.Close()

	// do GET
	resp, err := http.Get(ts.URL)

	if assert.Nil(err) {
		assert.Equal(resp.StatusCode, http.StatusOK, "they should be equal")
	}

	// swap router
	rt = http.NewServeMux()
	rt.HandleFunc("/", test404Handler)
	rs.Swap(rt)

	// do GET
	resp, err = http.Get(ts.URL)

	if assert.Nil(err) {
		assert.Equal(resp.StatusCode, http.StatusNotFound, "they should be equal")
	}
}
