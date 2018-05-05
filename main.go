package demo

import (
	"net/http"
	"fmt"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
	"google.golang.org/appengine"
)

func Worker(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Infof(ctx, "Worker succeeded")
}

func Handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	t := taskqueue.NewPOSTTask("/worker", map[string][]string{"key": {"val"}})
	_, err := taskqueue.Add(ctx, t, "")
	if err != nil {
		log.Errorf(ctx, "Failed to add task");
	}
	fmt.Fprintf(w, "Success");
}

func init() {
	http.HandleFunc("/", Handler)
	http.HandleFunc("/worker", Worker)
}
