
package dao

type Activity struct {
  ID        string     `gorm:"primaryKey"`
  Name      string     `gorm:"not null"`
  Duration  int        `gorm:"not null"`
  Intensity string     `gorm:"not null"`
  TrainerID string     `gorm:"not null"`
  Schedules []Schedule `gorm:"foreignKey:ActivityID;constraint:OnDelete:CASCADE"`
}
