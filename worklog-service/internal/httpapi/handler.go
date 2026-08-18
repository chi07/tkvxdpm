package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync/atomic"
	"time"

	"github.com/nhom2/worklog-service/internal/domain"
	"github.com/nhom2/worklog-service/internal/store"
)

type Handler struct { store *store.WorklogStore; sequence atomic.Uint64 }

func NewHandler(s *store.WorklogStore) http.Handler {
	h := &Handler{store: s}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", h.health)
	mux.HandleFunc("GET /api/v1/worklogs", h.listWorklogs)
	mux.HandleFunc("POST /api/v1/worklogs", h.createWorklog)
	mux.HandleFunc("GET /api/v1/worklogs/{id}", h.getWorklog)
	return logging(mux)
}

func (h *Handler) health(w http.ResponseWriter, _ *http.Request) { respond(w, http.StatusOK, map[string]string{"status":"ok"}) }

func (h *Handler) listWorklogs(w http.ResponseWriter, r *http.Request) {
	respond(w, http.StatusOK, h.store.List(strings.TrimSpace(r.URL.Query().Get("member_id"))))
}

func (h *Handler) getWorklog(w http.ResponseWriter, r *http.Request) {
	item, err := h.store.Get(r.PathValue("id"))
	if err == store.ErrNotFound { respondError(w, http.StatusNotFound, err.Error()); return }
	respond(w, http.StatusOK, item)
}

type createWorklogRequest struct { MemberID, Project, Label, WorkDate, Note string; Hours float64 `json:"hours"`; Approved bool `json:"approved"` }

func (h *Handler) createWorklog(w http.ResponseWriter, r *http.Request) {
	var req createWorklogRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil { respondError(w, http.StatusBadRequest, "invalid JSON body"); return }
	if strings.TrimSpace(req.MemberID)=="" || strings.TrimSpace(req.Project)=="" || strings.TrimSpace(req.Label)=="" || strings.TrimSpace(req.Note)=="" || req.Hours <= 0 || req.Hours > 24 { respondError(w, http.StatusBadRequest, "member_id, project, label, note and hours (0,24] are required"); return }
	workDate, err := time.Parse("2006-01-02", req.WorkDate)
	if err != nil { respondError(w, http.StatusBadRequest, "work_date must be YYYY-MM-DD"); return }
	id := fmt.Sprintf("wl-%03d", h.sequence.Add(1)+100)
	item := domain.Worklog{ID:id, MemberID:strings.TrimSpace(req.MemberID), Project:strings.TrimSpace(req.Project), Label:strings.TrimSpace(req.Label), WorkDate:workDate, Hours:req.Hours, Approved:req.Approved, Note:strings.TrimSpace(req.Note), CreatedAt:time.Now().UTC()}
	respond(w, http.StatusCreated, h.store.Create(item))
}

func respond(w http.ResponseWriter, status int, value any) { w.Header().Set("Content-Type","application/json"); w.WriteHeader(status); _=json.NewEncoder(w).Encode(value) }
func respondError(w http.ResponseWriter, status int, message string) { respond(w,status,map[string]string{"error":message}) }
func logging(next http.Handler) http.Handler { return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){ next.ServeHTTP(w,r) }) }
