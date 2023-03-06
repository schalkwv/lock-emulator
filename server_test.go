package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetNukiLocks(t *testing.T) {
	t.Run("return 200 and ok on nuki locks", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "nuki/smartlock", nil)
		response := httptest.NewRecorder()

		LockServer(response, request)

		got := response.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		gotBody := response.Body.String()
		wantBody := "ok"
		if gotBody != wantBody {
			t.Errorf("got %q, want %q", gotBody, wantBody)
		}

	})
}
