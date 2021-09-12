package album

type Album struct {
	ID     int    `gorm:"AUTO_INCREMENT"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	// CreatedAt time.Time `json:"createdAt"`
	// UpdatedAt time.Time `json:"updatedAt"`
}
