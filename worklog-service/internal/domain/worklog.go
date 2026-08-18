package domain

import "time"

// Worklog represents time recorded by a team member for one project activity.
type Worklog struct {
	ID        string    `json:"id"`
	MemberID  string    `json:"member_id"`
	Project   string    `json:"project"`
	Label     string    `json:"label"`
	WorkDate  time.Time `json:"work_date"`
	Hours     float64   `json:"hours"`
	Approved  bool      `json:"approved"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
}
