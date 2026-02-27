package dao

type Inscription struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	UserID     int    `gorm:"not null;index"`
	ActivityID string `gorm:"type:VARCHAR(255);not null;index"`
	DayOfWeek  string `gorm:"type:VARCHAR(50);not null"`
}

