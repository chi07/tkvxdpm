package httpapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/nhom2/worklog-service/internal/store"
)

func TestHealth(t *testing.T) {
	r := httptest.NewRecorder(); NewHandler(store.NewSeededWorklogStore()).ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if r.Code != http.StatusOK { t.Fatalf("status = %d, want 200", r.Code) }
}
func TestListWorklogsFiltersByMember(t *testing.T) {
	r := httptest.NewRecorder(); NewHandler(store.NewSeededWorklogStore()).ServeHTTP(r, httptest.NewRequest(http.MethodGet, "/api/v1/worklogs?member_id=hai", nil))
	if r.Code != http.StatusOK { t.Fatalf("status = %d, want 200", r.Code) }
	if got := r.Body.String(); !bytes.Contains([]byte(got), []byte("wl-001")) || bytes.Contains([]byte(got), []byte("wl-002")) { t.Fatalf("unexpected body: %s", got) }
}
func TestCreateWorklogRejectsInvalidDate(t *testing.T) {
	body := bytes.NewBufferString(`{"member_id":"hai","project":"billing","label":"backend","work_date":"bad","hours":2,"note":"test"}`)
	r := httptest.NewRecorder(); NewHandler(store.NewSeededWorklogStore()).ServeHTTP(r, httptest.NewRequest(http.MethodPost, "/api/v1/worklogs", body))
	if r.Code != http.StatusBadRequest { t.Fatalf("status = %d, want 400", r.Code) }
}
