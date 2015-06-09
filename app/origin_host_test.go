package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestAPIOriginHost(t *testing.T) {
	ts := httptest.NewServer(Router(StaticReport{}))
	defer ts.Close()

	is404(t, ts, "/api/origin/foobar")
	is404(t, ts, "/api/origin/host/foobar")

	{
		// Origin
		body := getRawJSON(t, ts, "/api/origin/host/hostA;<host>") // TODO MakeHostNodeID
		var o OriginHost
		if err := json.Unmarshal(body, &o); err != nil {
			t.Fatalf("JSON parse error: %s", err)
		}
		if want, have := "Linux", o.OS; want != have {
			t.Errorf("Origin error. Want %v, have %v", want, have)
		}
		if want, have := "3.14 2.71 1.61", o.Load; want != have {
			t.Errorf("Origin error. Want %v, have %v", want, have)
		}
	}
}
