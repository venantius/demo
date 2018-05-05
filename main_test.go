package demo_test

import (
	"demo"
	"time"
	"testing"
	"google.golang.org/appengine/aetest"
	"net/http/httptest"
	"net/http"
)

func TestTestHandler(t *testing.T) {
	inst, err := aetest.NewInstance(nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	defer inst.Close()

	req, err := inst.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	resp := httptest.NewRecorder()

	http.HandlerFunc(demo.Handler).ServeHTTP(resp, req)

	time.Sleep(10 * time.Second)
}
