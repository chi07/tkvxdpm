package store

import (
	"errors"
	"sort"
	"sync"
	"time"

	"github.com/nhom2/worklog-service/internal/domain"
)

var ErrNotFound = errors.New("worklog not found")

// WorklogStore is deliberately small but has realistic concerns: synchronization,
// filtering, validation at the transport boundary, and stable list ordering.
type WorklogStore struct {
	mu    sync.RWMutex
	items map[string]domain.Worklog
}

func NewSeededWorklogStore() *WorklogStore {
	base := time.Date(2026, 8, 10, 0, 0, 0, 0, time.UTC)
	items := []domain.Worklog{
		{ID: "wl-001", MemberID: "hai", Project: "billing", Label: "backend", WorkDate: base, Hours: 3, Approved: true, Note: "Fix invoice validation", CreatedAt: base.Add(9 * time.Hour)},
		{ID: "wl-002", MemberID: "chi", Project: "billing", Label: "testing", WorkDate: base.AddDate(0, 0, 1), Hours: 2.5, Approved: false, Note: "Add API regression tests", CreatedAt: base.Add(24*time.Hour + 9*time.Hour)},
		{ID: "wl-003", MemberID: "trang", Project: "portal", Label: "frontend", WorkDate: base.AddDate(0, 0, 2), Hours: 4, Approved: true, Note: "Improve loading state", CreatedAt: base.Add(48*time.Hour + 9*time.Hour)},
		{ID: "wl-004", MemberID: "hung", Project: "portal", Label: "backend", WorkDate: base.AddDate(0, 0, 3), Hours: 1.5, Approved: true, Note: "Update activity endpoint", CreatedAt: base.Add(72*time.Hour + 9*time.Hour)},
		{ID: "wl-005", MemberID: "hai", Project: "portal", Label: "testing", WorkDate: base.AddDate(0, 0, 4), Hours: 2, Approved: false, Note: "Review acceptance cases", CreatedAt: base.Add(96*time.Hour + 9*time.Hour)},
	}
	s := &WorklogStore{items: make(map[string]domain.Worklog, len(items))}
	for _, item := range items { s.items[item.ID] = item }
	return s
}

func (s *WorklogStore) List(memberID string) []domain.Worklog {
	s.mu.RLock(); defer s.mu.RUnlock()
	result := make([]domain.Worklog, 0, len(s.items))
	for _, item := range s.items {
		if memberID == "" || item.MemberID == memberID { result = append(result, item) }
	}
	sort.Slice(result, func(i, j int) bool { return result[i].ID < result[j].ID })
	return result
}

func (s *WorklogStore) Get(id string) (domain.Worklog, error) {
	s.mu.RLock(); defer s.mu.RUnlock()
	item, ok := s.items[id]
	if !ok { return domain.Worklog{}, ErrNotFound }
	return item, nil
}

func (s *WorklogStore) Create(item domain.Worklog) domain.Worklog {
	s.mu.Lock(); defer s.mu.Unlock()
	s.items[item.ID] = item
	return item
}
