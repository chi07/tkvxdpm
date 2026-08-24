package store

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/nhom2/worklog-service/internal/domain"
)

var ErrNotFound = errors.New("worklog not found")

// WorklogStore persists worklogs in SQLite. The default constructor uses an
// in-memory SQLite database so tests and demos remain self-contained; callers
// that need a file-backed database can use NewSQLiteWorklogStore.
type WorklogStore struct{ db *sql.DB }

func NewSeededWorklogStore() *WorklogStore {
	s, err := NewSQLiteWorklogStore(":memory:")
	if err != nil {
		panic(fmt.Sprintf("initialize sqlite worklog store: %v", err))
	}
	if err := s.seed(); err != nil {
		panic(fmt.Sprintf("seed sqlite worklog store: %v", err))
	}
	return s
}

func NewSQLiteWorklogStore(dsn string) (*WorklogStore, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(1)
	s := &WorklogStore{db: db}
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS worklogs (
		id TEXT PRIMARY KEY, member_id TEXT NOT NULL, project TEXT NOT NULL,
		label TEXT NOT NULL, work_date TEXT NOT NULL, hours REAL NOT NULL,
		approved INTEGER NOT NULL DEFAULT 0, note TEXT NOT NULL, created_at TEXT NOT NULL
	)`); err != nil {
		_ = db.Close()
		return nil, err
	}
	return s, nil
}

func (s *WorklogStore) seed() error {
	base := time.Date(2026, 8, 10, 0, 0, 0, 0, time.UTC)
	items := []domain.Worklog{
		{ID: "wl-001", MemberID: "hai", Project: "billing", Label: "backend", WorkDate: base, Hours: 3, Approved: true, Note: "Fix invoice validation", CreatedAt: base.Add(9 * time.Hour)},
		{ID: "wl-002", MemberID: "chi", Project: "billing", Label: "testing", WorkDate: base.AddDate(0, 0, 1), Hours: 2.5, Note: "Add API regression tests", CreatedAt: base.Add(33 * time.Hour)},
		{ID: "wl-003", MemberID: "trang", Project: "portal", Label: "frontend", WorkDate: base.AddDate(0, 0, 2), Hours: 4, Approved: true, Note: "Improve loading state", CreatedAt: base.Add(57 * time.Hour)},
		{ID: "wl-004", MemberID: "hung", Project: "portal", Label: "backend", WorkDate: base.AddDate(0, 0, 3), Hours: 1.5, Approved: true, Note: "Update activity endpoint", CreatedAt: base.Add(81 * time.Hour)},
		{ID: "wl-005", MemberID: "hai", Project: "portal", Label: "testing", WorkDate: base.AddDate(0, 0, 4), Hours: 2, Note: "Review acceptance cases", CreatedAt: base.Add(105 * time.Hour)},
	}
	for _, item := range items {
		if _, err := s.db.Exec(`INSERT INTO worklogs (id,member_id,project,label,work_date,hours,approved,note,created_at) VALUES (?,?,?,?,?,?,?,?,?)`, item.ID, item.MemberID, item.Project, item.Label, item.WorkDate.Format("2006-01-02"), item.Hours, item.Approved, item.Note, item.CreatedAt.Format(time.RFC3339)); err != nil {
			return err
		}
	}
	return nil
}

// SeedIfEmpty inserts the demo fixtures only when the database has no rows.
func (s *WorklogStore) SeedIfEmpty() error {
	var count int
	if err := s.db.QueryRow(`SELECT COUNT(*) FROM worklogs`).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	return s.seed()
}

func (s *WorklogStore) List(memberID string) []domain.Worklog {
	query := `SELECT id,member_id,project,label,work_date,hours,approved,note,created_at FROM worklogs`
	args := []any{}
	if memberID != "" {
		query += ` WHERE member_id = ?`
		args = append(args, memberID)
	}
	query += ` ORDER BY id`
	rows, err := s.db.Query(query, args...)
	if err != nil {
		return []domain.Worklog{}
	}
	defer rows.Close()
	result := []domain.Worklog{}
	for rows.Next() {
		if item, err := scanWorklog(rows); err == nil {
			result = append(result, item)
		}
	}
	return result
}

func (s *WorklogStore) Get(id string) (domain.Worklog, error) {
	row := s.db.QueryRow(`SELECT id,member_id,project,label,work_date,hours,approved,note,created_at FROM worklogs WHERE id = ?`, id)
	item, err := scanWorklog(row)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.Worklog{}, ErrNotFound
	}
	return item, err
}

func (s *WorklogStore) Create(item domain.Worklog) domain.Worklog {
	_, _ = s.db.Exec(`INSERT INTO worklogs (id,member_id,project,label,work_date,hours,approved,note,created_at) VALUES (?,?,?,?,?,?,?,?,?)`, item.ID, item.MemberID, item.Project, item.Label, item.WorkDate.Format("2006-01-02"), item.Hours, item.Approved, item.Note, item.CreatedAt.Format(time.RFC3339))
	return item
}
func (s *WorklogStore) Close() error { return s.db.Close() }

type rowScanner interface{ Scan(dest ...any) error }

func scanWorklog(row rowScanner) (domain.Worklog, error) {
	var item domain.Worklog
	var workDate, createdAt string
	var approved bool
	err := row.Scan(&item.ID, &item.MemberID, &item.Project, &item.Label, &workDate, &item.Hours, &approved, &item.Note, &createdAt)
	if err != nil {
		return domain.Worklog{}, err
	}
	item.Approved = approved
	item.WorkDate, err = time.Parse("2006-01-02", workDate)
	if err != nil {
		return domain.Worklog{}, err
	}
	item.CreatedAt, err = time.Parse(time.RFC3339, createdAt)
	return item, err
}
