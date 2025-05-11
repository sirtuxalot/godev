/* internal/models/models.go */

package models

import "time"

type Tickets struct {
	TicketID    int
	Requestor   string
	Title       string
	Description string
	Closed      bool
	InProgress  bool
	CategoryID  int
	Hours       float32
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Comments struct {
	CommentID int
	TicketID  int
	Comment   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Categories struct {
	CategoryID int
	Category   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
