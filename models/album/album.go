package album

import "time"

type Album struct {
	ID        uint
	Title     string
	Artist    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
