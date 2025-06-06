// backend/dao/schedule_dao.go
package dao

type Schedule struct {
  ID         uint   `gorm:"primaryKey;autoIncrement"`
  ActivityID string `gorm:"type:VARCHAR(255);not null;index"`
  DayOfWeek  string `gorm:"type:VARCHAR(50);not null"`
  StartTime  string `gorm:"type:VARCHAR(5);not null"`
  EndTime    string `gorm:"type:VARCHAR(5);not null"`
}
