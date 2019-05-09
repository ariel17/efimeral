package sessions

import "time"

// SessionRequest responds to the required data needed to create a new session.
type SessionRequest struct {
	Distribution string `json:"distribution"`
}

// Session represents a SO container instance on the Docker service.
type Session struct {
	ID string `json:"id"`
	SessionRequest
	Tag       string     `json:"tag"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	Running   bool       `json:"running"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
