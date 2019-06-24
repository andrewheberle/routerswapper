// Package routerswapper implements a simple process to swap a Golang HTTP
// router (gorilla/mux etc) during runtime. This is aimed at allowing
// route/handler changes based on a configuration change or update.
package routerswapper

import (
	"net/http"
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

	// defaults
	testScheme := "http"
	testHost := "127.0.0.1:8080"

	// new router
	rt := http.NewServeMux()
	rt.HandleFunc("/200", test200Handler)
	rt.HandleFunc("/404", test404Handler)
	rs := New(rt)
	go func() {
		err := http.ListenAndServe(testHost, rs)
		assert.Nil(err)
	}()

	// do GET
	resp, err := http.Get(testScheme+"://"+testHost+"/200")

	if assert.Nil(err) {
		assert.NotEqual(resp.StatusCode, http.StatusNotFound, "they should not be equal")
		assert.Equal(resp.StatusCode, http.StatusOK, "they should be equal")
	}

	// do GET
	resp, err = http.Get(testScheme+"://"+testHost+"/404")

	if assert.Nil(err) {
		assert.Equal(resp.StatusCode, http.StatusNotFound, "they should be equal")
		assert.NotEqual(resp.StatusCode, http.StatusOK, "they should not be equal")
	}

	// swap router
	rt = http.NewServeMux()
	rt.HandleFunc("/200", test404Handler)
	rt.HandleFunc("/404", test200Handler)
	rs.Swap(rt)

	// do GET
	resp, err = http.Get(testScheme+"://"+testHost+"/200")

	if assert.Nil(err) {
		assert.Equal(resp.StatusCode, http.StatusNotFound, "they should be equal")
		assert.NotEqual(resp.StatusCode, http.StatusOK, "they should not be equal")
	}

	// do GET
	resp, err = http.Get(testScheme+"://"+testHost+"/404")

	if assert.Nil(err) {
		assert.NotEqual(resp.StatusCode, http.StatusNotFound, "they should not be equal")
		assert.Equal(resp.StatusCode, http.StatusOK, "they should be equal")
	}
}
