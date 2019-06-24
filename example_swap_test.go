package routerswapper_test

import (
	"fmt"
	"log"
	"net/http"

	"gitlab.com/andrewheberle/routerswapper"
)

func Example() {
	r := http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	rs := routerswapper.New(r)

	// Start HTTP server in a goroutine
	go func() {
		if err := http.ListenAndServe("127.0.0.1:8080", rs); err != nil {
			log.Fatalf("http error: %s", err)
		}
	}()

	// do GET
	resp, err := http.Get("http://127.0.0.1:8080/")
	if err != nil {
		log.Fatalf("get error: %s", err)
	}

	fmt.Printf("status: %d\n", resp.StatusCode)

	// Swap the router
	r = http.NewServeMux()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	rs.Swap(r)

	// second GET
	resp, err = http.Get("http://127.0.0.1:8080/")
	if err != nil {
		log.Fatalf("get error: %s", err)
	}

	fmt.Printf("status: %d\n", resp.StatusCode)

	// Output:
	// status: 200
	// status: 404

}
