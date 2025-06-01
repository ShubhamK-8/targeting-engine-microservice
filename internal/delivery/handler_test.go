package delivery

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestDelivery(t *testing.T) {
    req := httptest.NewRequest("GET", "/v1/delivery?app=com.gametion.ludokinggame&country=us&os=android", nil)
    w := httptest.NewRecorder()

    HandleDelivery(w, req)

    res := w.Result()
    if res.StatusCode != http.StatusOK {
        t.Errorf("expected 200 OK, got %d", res.StatusCode)
    }
}
